package command

import (
	"fmt"
	"os/exec"
)

type Command struct{}

func New() *Command {
	return &Command{}
}

func (c *Command) RunFlushDNS() error {
	cmd := exec.Command("ipconfig", "/flushdns")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Must run flushdns manually\n Error: %v", err)
	}
	return nil
}
