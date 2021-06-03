// +build linux_bpf

package http

import (
	"math"
	"os"

	ddebpf "github.com/DataDog/datadog-agent/pkg/ebpf"
	"github.com/DataDog/datadog-agent/pkg/ebpf/bytecode"
	"github.com/DataDog/datadog-agent/pkg/network/config"
	netebpf "github.com/DataDog/datadog-agent/pkg/network/ebpf"
	"github.com/DataDog/datadog-agent/pkg/network/ebpf/probes"
	"github.com/DataDog/ebpf"
	"github.com/DataDog/ebpf/manager"
	"golang.org/x/sys/unix"
)

const (
	// maxActive configures the maximum number of instances of the kretprobe-probed functions handled simultaneously.
	// This value should be enough for typical workloads (e.g. some amount of processes blocked on the accept syscall).
	maxActive                = 128
	defaultClosedChannelSize = 500

	// TODO: this is only hardcoded here for the PoC
	libSSL = "/lib/x86_64-linux-gnu/libssl.so.1.1"
)

type ebpfProgram struct {
	*manager.Manager
	cfg         *config.Config
	perfHandler *ddebpf.PerfHandler
	bytecode    bytecode.AssetReader
	offsets     []manager.ConstantEditor
}

func newEBPFProgram(c *config.Config, offsets []manager.ConstantEditor) (*ebpfProgram, error) {
	bytecode, err := netebpf.ReadHTTPModule(c.BPFDir, c.BPFDebug)
	if err != nil {
		return nil, err
	}

	closedChannelSize := defaultClosedChannelSize
	if c.ClosedChannelSize > 0 {
		closedChannelSize = c.ClosedChannelSize
	}
	perfHandler := ddebpf.NewPerfHandler(closedChannelSize)

	mgr := &manager.Manager{
		Maps: []*manager.Map{
			{Name: string(probes.HttpInFlightMap)},
			{Name: string(probes.HttpBatchesMap)},
			{Name: string(probes.HttpBatchStateMap)},
			{Name: "sockfd_by_pid_tgid"},
			{Name: "tup_by_pid_sockfd"},
			{Name: "tup_by_ssl_ctx"},
			{Name: "ssl_read_args"},
		},
		PerfMaps: []*manager.PerfMap{
			{
				Map: manager.Map{Name: string(probes.HttpNotificationsMap)},
				PerfMapOptions: manager.PerfMapOptions{
					PerfRingBufferSize: 8 * os.Getpagesize(),
					Watermark:          1,
					DataHandler:        perfHandler.DataHandler,
					LostHandler:        perfHandler.LostHandler,
				},
			},
		},
		Probes: []*manager.Probe{
			{Section: string(probes.TCPSendMsgReturn), KProbeMaxActive: maxActive},
			{Section: string(probes.SockMapFdReturn), KProbeMaxActive: maxActive},
			{Section: "kprobe/sockfd_lookup_light", KProbeMaxActive: maxActive},
			{Section: "uprobe/SSL_set_fd", BinaryPath: libSSL},
			{Section: "uprobe/SSL_read", BinaryPath: libSSL},
			{Section: "uretprobe/SSL_read", BinaryPath: libSSL},
			{Section: "uprobe/SSL_write", BinaryPath: libSSL},
			{Section: "uprobe/SSL_shutdown", BinaryPath: libSSL},
			{Section: string(probes.SocketHTTPFilter)},
		},
	}

	return &ebpfProgram{
		Manager:     mgr,
		perfHandler: perfHandler,
		bytecode:    bytecode,
		cfg:         c,
		offsets:     offsets,
	}, nil
}

func (e *ebpfProgram) Init() error {
	defer e.bytecode.Close()

	return e.InitWithOptions(e.bytecode, manager.Options{
		RLimit: &unix.Rlimit{
			Cur: math.MaxUint64,
			Max: math.MaxUint64,
		},
		MapSpecEditors: map[string]manager.MapSpecEditor{
			string(probes.HttpInFlightMap): {
				Type:       ebpf.Hash,
				MaxEntries: uint32(e.cfg.MaxTrackedConnections),
				EditorFlag: manager.EditMaxEntries,
			},
		},
		ActivatedProbes: []manager.ProbesSelector{
			&manager.ProbeSelector{
				ProbeIdentificationPair: manager.ProbeIdentificationPair{
					Section: string(probes.SocketHTTPFilter),
				},
			},
			&manager.ProbeSelector{
				ProbeIdentificationPair: manager.ProbeIdentificationPair{
					Section: string(probes.TCPSendMsgReturn),
				},
			},
			&manager.ProbeSelector{
				ProbeIdentificationPair: manager.ProbeIdentificationPair{
					Section: string(probes.SockMapFdReturn),
				},
			},
			&manager.ProbeSelector{
				ProbeIdentificationPair: manager.ProbeIdentificationPair{
					Section: "kprobe/sockfd_lookup_light",
				},
			},
			&manager.ProbeSelector{
				ProbeIdentificationPair: manager.ProbeIdentificationPair{
					Section: "uprobe/SSL_set_fd",
				},
			},
			&manager.ProbeSelector{
				ProbeIdentificationPair: manager.ProbeIdentificationPair{
					Section: "uprobe/SSL_read",
				},
			},
			&manager.ProbeSelector{
				ProbeIdentificationPair: manager.ProbeIdentificationPair{
					Section: "uretprobe/SSL_read",
				},
			},
			&manager.ProbeSelector{
				ProbeIdentificationPair: manager.ProbeIdentificationPair{
					Section: "uprobe/SSL_write",
				},
			},
			&manager.ProbeSelector{
				ProbeIdentificationPair: manager.ProbeIdentificationPair{
					Section: "uprobe/SSL_shutdown",
				},
			},
		},
		ConstantEditors: e.offsets,
	})
}
