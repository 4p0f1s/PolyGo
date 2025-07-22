//go:build windows
// +build windows

package commandloader

import (
	"os/exec"
	"syscall"
)

func ShellCmd(command string) *exec.Cmd {
	c := exec.Command("cmd.exe", "/C", command)
	c.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	return c
}
