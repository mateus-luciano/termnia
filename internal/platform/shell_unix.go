//go:build !windows

package platform

import "os/exec"

func ConfigurePlatformSpecific(cmd *exec.Cmd) {
	// On Unix-like systems we don't need to hide windows
}
