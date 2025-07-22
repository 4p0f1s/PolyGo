package main

import (
	"os/exec"
	"syscall"
)

func getShellCmd() *exec.Cmd {
	cmd := exec.Command("cmd.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	return cmd
}
