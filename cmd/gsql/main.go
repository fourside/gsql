package main

import (
	"fmt"
	"github.com/fourside/gsql"
	"github.com/mitchellh/cli"
	"os"
)

func main() {
	c := cli.NewCLI("gsql", gsql.Version())
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"check": gsql.CheckCommandFactory,
	}
	exitCode, err := c.Run()
	if err != nil {
		fmt.Print("Error: %s", err.Error())
	}
	os.Exit(exitCode)
}
