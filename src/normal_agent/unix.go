package main

import "os/exec"

func getShellCmd() *exec.Cmd {
	return exec.Command("/bin/bash")
}
