//go:build windows

package platform

import (
	"os/exec"
	"syscall"
)

func ConfigurePlatformSpecific(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
