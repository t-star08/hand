package main

import (
	"github.com/spf13/cobra"
	"github.com/t-star08/hand/cmd/initcmd"
	"github.com/t-star08/hand/cmd/insertcmd"
	"github.com/t-star08/hand/cmd/keepcmd"
	"github.com/t-star08/hand/cmd/listcmd"
	"github.com/t-star08/hand/cmd/modifycmd"
	"github.com/t-star08/hand/cmd/removecmd"
	"github.com/t-star08/hand/cmd/reportcmd"
	"github.com/t-star08/hand/cmd/statuscmd"
)

var cmd = &cobra.Command {
	Use: "hand",
	Version: "v0.0.2",
}

func init() {
	cmd.AddCommand (
		initcmd.CMD,
		insertcmd.CMD,
		keepcmd.CMD,
		listcmd.CMD,
		modifycmd.CMD,
		removecmd.CMD,
		reportcmd.CMD,
		statuscmd.CMD,
	)
}

func main() {
	cmd.Execute()
}
