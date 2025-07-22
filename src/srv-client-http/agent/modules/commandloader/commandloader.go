package commandloader

import (
	"log"
)

func LoaderCommand(command string) ([]byte, error) {
	c := ShellCmd(command)

	output, err := c.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return output, err
}
