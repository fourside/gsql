package gsql

import (
	"fmt"
	"github.com/mitchellh/cli"
)

type createCommand struct {
	ui cli.Ui
}

func (c *createCommand) Run(args []string) int {
	c.ui.Info(fmt.Sprintf("create command %v", args))
	if len(args) != 2 {
		c.ui.Error(c.Help())
		return 1
	}
	return 0
}

func (c *createCommand) Synopsis() string {
	return "create an account"
}

func (c *createCommand) Help() string {
	return "Usage: gsql create {account} {password}"
}
