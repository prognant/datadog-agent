// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// +build functionaltests

package tests

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
	"testing"
	"time"

	"github.com/DataDog/datadog-agent/pkg/security/rules"
	"gotest.tools/assert"
)

const TEST_BOOL_NAME = "selinuxuser_ping"

func TestSELinux(t *testing.T) {
	rules := []*rules.RuleDefinition{
		{
			ID:         "test_selinux_enforce",
			Expression: `selinux.file.name == "enforce"`,
		},
		{
			ID:         "test_selinux_write_bool_true",
			Expression: `selinux.file.name == "selinuxuser_ping" && selinux.write.bool_value == true`,
		},
		{
			ID:         "test_selinux_write_bool_false",
			Expression: `selinux.file.name == "selinuxuser_ping" && selinux.write.bool_value == false`,
		},
	}

	test, err := newTestModule(nil, rules, testOpts{})
	if err != nil {
		t.Fatal(err)
	}
	defer test.Close()

	if cmd := exec.Command("sudo", "-n", "sestatus"); cmd.Run() != nil {
		t.Skipf("SELinux is not available, skipping tests")
	}

	savedBoolValue, err := getBoolValue(TEST_BOOL_NAME)
	if err != nil {
		t.Errorf("failed to save bool state: %v", err)
	}
	defer setBoolValue(TEST_BOOL_NAME, savedBoolValue)

	t.Run("setenforce", func(t *testing.T) {
		if cmd := exec.Command("sudo", "-n", "setenforce", "0"); cmd.Run() != nil {
			t.Errorf("Failed to run setenforce")
		}

		event, rule, err := test.GetEvent()
		if err != nil {
			t.Error(err)
		} else {
			assertTriggeredRule(t, rule, "test_selinux_enforce")
			assert.Equal(t, event.GetType(), "selinux", "wrong event type")

			assertFieldEqual(t, event, "selinux.write.bool_value", false, "wrong enforce value")

			fileName := "/sys/fs/selinux/enforce"
			assertFieldEqual(t, event, "selinux.file.path", fileName, "wrong file path")
			assertFieldEqual(t, event, "selinux.file.name", "enforce", "wrong file name")
			assertFieldEqual(t, event, "selinux.file.inode", int(getInode(t, fileName)), "wrong inode")

			if testEnvironment == DockerEnvironment {
				testContainerPath(t, event, "rename.file.container_path")
				testContainerPath(t, event, "rename.file.destination.container_path")
			}
		}
	})

	t.Run("setsebool_true_value", func(t *testing.T) {
		if err := setBoolValue(TEST_BOOL_NAME, true); err != nil {
			t.Errorf("failed to run setsebool: %v", err)
		}

		event, rule, err := test.GetEvent()
		if err != nil {
			t.Error(err)
		} else {
			assertTriggeredRule(t, rule, "test_selinux_write_bool_true")
			assert.Equal(t, event.SELinux.File.BasenameStr, TEST_BOOL_NAME, "wrong bool name")

			assert.Equal(t, event.GetType(), "selinux", "wrong event type")

			assertFieldEqual(t, event, "selinux.write.bool_value", true, "wrong bool value")

			fileName := "/sys/fs/selinux/booleans/selinuxuser_ping"
			assertFieldEqual(t, event, "selinux.file.path", fileName, "wrong file path")
			assertFieldEqual(t, event, "selinux.file.name", TEST_BOOL_NAME, "wrong file name")
			assertFieldEqual(t, event, "selinux.file.inode", int(getInode(t, fileName)), "wrong inode")

			if testEnvironment == DockerEnvironment {
				testContainerPath(t, event, "rename.file.container_path")
				testContainerPath(t, event, "rename.file.destination.container_path")
			}
		}
	})

	t.Run("setsebool_false_value", func(t *testing.T) {
		if err := setBoolValue(TEST_BOOL_NAME, false); err != nil {
			t.Errorf("failed to run setsebool: %v", err)
		}

		event, rule, err := test.GetEvent()
		if err != nil {
			t.Error(err)
		} else {
			assertTriggeredRule(t, rule, "test_selinux_write_bool_false")
			assert.Equal(t, event.SELinux.File.BasenameStr, TEST_BOOL_NAME, "wrong bool name")

			assert.Equal(t, event.GetType(), "selinux", "wrong event type")

			assertFieldEqual(t, event, "selinux.write.bool_value", false, "wrong bool value")

			fileName := "/sys/fs/selinux/booleans/selinuxuser_ping"
			assertFieldEqual(t, event, "selinux.file.path", fileName, "wrong file path")
			assertFieldEqual(t, event, "selinux.file.name", TEST_BOOL_NAME, "wrong file name")
			assertFieldEqual(t, event, "selinux.file.inode", int(getInode(t, fileName)), "wrong inode")

			if testEnvironment == DockerEnvironment {
				testContainerPath(t, event, "rename.file.container_path")
				testContainerPath(t, event, "rename.file.destination.container_path")
			}
		}
	})

	t.Run("setsebool_error_value", func(t *testing.T) {
		cmd := exec.Command("sudo", "-n", "tee", "/sys/fs/selinux/booleans/httpd_enable_cgi")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			t.Errorf("failed to get stdin of tee cmd: %v", err)
		}

		go func() {
			defer stdin.Close()
			io.WriteString(stdin, "test_error")
		}()

		cmd.Run()

		_, _, err = test.GetEventWithTimeout(1 * time.Second)
		assert.Equal(t, err.Error(), "timeout", "wrong error type, expected timeout")
	})
}

func getBoolValue(boolName string) (bool, error) {
	cmd := exec.Command("sudo", "-n", "getsebool", boolName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}

	var name, value string
	fmt.Sscanf(string(output), "%s --> %s", &name, &value)

	if name != boolName {
		return false, errors.New("bool name mismatch")
	}

	switch value {
	case "on":
		return true, nil
	case "off":
		return false, nil
	default:
		return false, fmt.Errorf("unknown bool representation: %v", value)
	}
}

func setBoolValue(boolName string, value bool) error {
	var valueStr string
	if value {
		valueStr = "on"
	} else {
		valueStr = "off"
	}

	cmd := exec.Command("sudo", "-n", "setsebool", boolName, valueStr)
	return cmd.Run()
}
