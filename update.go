package gsql

import (
	"fmt"
	"github.com/mitchellh/cli"
)

type updateCommand struct {
	ui cli.Ui
}

func (c *updateCommand) Run(args []string) int {
	c.ui.Info(fmt.Sprintf("update command %v", args))
	if len(args) != 2 {
		c.ui.Error(c.Help())
		return 1
	}
	return 0
}

func (c *updateCommand) Synopsis() string {
	return "update an account password"
}

func (c *updateCommand) Help() string {
	return "Usage: gsql update {account} {password}"
}
