package gsql

import (
	"fmt"
	"github.com/mitchellh/cli"
)

type checkCommand struct {
	ui cli.Ui
}

func (c *checkCommand) Run(args []string) int {
	c.ui.Info(fmt.Sprintf("check command %v", args))
	return 0
}

func (c *checkCommand) Synopsis() string {
	return "Check if exists"
}

func (c *checkCommand) Help() string {
	return "Usage: gsql check {account}"
}
