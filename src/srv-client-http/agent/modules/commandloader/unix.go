//go:build !windows
// +build !windows

package commandloader

import "os/exec"

func ShellCmd(command string) *exec.Cmd {
	return exec.Command("bash", "-c", command)
}
