package main

import (
	"agent/modules/connector"
	"time"
)

func main() {

	const wait time.Duration = 3
	for {

		connector.Controll()
		time.Sleep(wait * time.Second)
	}
}
