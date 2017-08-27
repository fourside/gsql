package gsql

import (
	"github.com/mitchellh/cli"
	"os"
)

var Ui cli.Ui

func init() {
	Ui = &cli.PrefixedUi{
		InfoPrefix:  "INFO: ",
		WarnPrefix:  "WARN: ",
		ErrorPrefix: "ERROR: ",
		Ui: &cli.BasicUi{
			Writer: os.Stdout,
		},
	}
}

func CheckCommandFactory() (cli.Command, error) {
	return &checkCommand{
		ui: Ui,
	}, nil
}
