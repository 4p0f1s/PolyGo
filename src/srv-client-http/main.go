// main.go
package main

import (
	"badserver/modules"
	"badserver/modules/srv"
)

func main() {
	modules.Banner()
	srv.ServerStart("8080")
}
