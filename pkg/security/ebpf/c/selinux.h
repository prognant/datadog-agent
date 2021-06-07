#ifndef _SELINUX_H_
#define _SELINUX_H_

#include "defs.h"
#include "filters.h"
#include "syscalls.h"
#include "process.h"

struct selinux_event_t {
    struct kevent_t event;
    struct process_context_t process;
    struct container_context_t container;
    struct file_t file;
    u32 magic;
};

int __attribute__((always_inline)) handle_selinux_event(void *ctx, struct file *file) {
    struct syscall_cache_t syscall = {
        .type = EVENT_SELINUX,
        .policy = fetch_policy(EVENT_SELINUX),
        .selinux = {},
    };

    struct dentry *dentry = get_file_dentry(file);
    syscall.selinux.dentry = dentry;
    syscall.selinux.file.path_key.mount_id = get_file_mount_id(file);

    cache_syscall(&syscall);
    return 0;
}

int __attribute__((always_inline)) handle_selinux_ret_event(void *ctx, int retval) {
    if (IS_UNHANDLED_ERROR(retval))
        return 0;

    struct syscall_cache_t *syscall = peek_syscall(EVENT_SELINUX);
    if (!syscall)
        return 0;

    if (syscall->discarded)
        return 0;

    fill_file_metadata(syscall->selinux.dentry, &syscall->selinux.file.metadata);
    set_file_inode(syscall->selinux.dentry, &syscall->selinux.file, 0);

    syscall->resolver.key = syscall->selinux.file.path_key;
    syscall->resolver.dentry = syscall->selinux.dentry;
    syscall->resolver.discarder_type = syscall->policy.mode != NO_FILTER ? EVENT_SELINUX : 0;
    syscall->resolver.callback = DR_SELINUX_CALLBACK_KPROBE_KEY;
    syscall->resolver.iteration = 0;
    syscall->resolver.ret = 0;

    // tail call
    resolve_dentry(ctx, DR_KPROBE);

    // if the tail call fails, we need to pop the syscall cache entry
    pop_syscall(EVENT_SELINUX);
    return 0;
}

int __attribute__((always_inline)) dr_selinux_callback(void *ctx, int retval) {
    struct syscall_cache_t *syscall = pop_syscall(EVENT_SELINUX);
    if (!syscall)
        return 0;

    if (syscall->resolver.ret == DENTRY_DISCARDED || syscall->resolver.ret == DENTRY_INVALID)
        return 0;

    struct selinux_event_t event = {
        .magic = 42,
    };
    event.file = syscall->selinux.file;

    struct proc_cache_t *entry = fill_process_context(&event.process);
    fill_container_context(entry, &event.container);

    send_event(ctx, EVENT_SELINUX, event);
    return 0;
}

SEC("kprobe/dr_selinux_callback")
int __attribute__((always_inline)) kprobe_dr_selinux_callback(struct pt_regs *ctx) {
    int retval = PT_REGS_RC(ctx);
    return dr_selinux_callback(ctx, retval);
}

#define PROBE_SEL_WRITE_FUNC(func_name)                        \
    SEC("kprobe/" #func_name)                                  \
    int kprobe__##func_name(struct pt_regs *ctx) {             \
        struct file *file = (struct file *)PT_REGS_PARM1(ctx); \
        return handle_selinux_event(ctx, file);                \
    }                                                          \
                                                               \
    SEC("kretprobe/" #func_name)                               \
    int kretprobe__##func_name(struct pt_regs *ctx) {          \
        int retval = PT_REGS_RC(ctx);                          \
        return handle_selinux_ret_event(ctx, retval);          \
    }

PROBE_SEL_WRITE_FUNC(sel_write_enforce)
PROBE_SEL_WRITE_FUNC(sel_write_bool)

#endif
