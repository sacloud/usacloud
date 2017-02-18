package main

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/version"
	"gopkg.in/urfave/cli.v2"
	"os"
	"sort"
)

var (
	appName      = "usacloud"
	appUsage     = "Unofficial 'sacloud' - CLI client of the SakuraCloud"
	appCopyright = "Copyright (C) 2017 Kazumichi Yamamoto."
)

func main() {

	app := &cli.App{}

	app.Name = appName
	app.Usage = appUsage
	app.HelpName = appName
	app.Copyright = appCopyright

	app.Flags = command.GlobalFlags
	app.Commands = command.Commands
	app.Version = version.FullVersion()

	app.EnableShellCompletion = true
	app.ShellComplete = func(c *cli.Context) {
		// This will complete if no args are passed
		if c.NArg() > 0 {
			return
		}

		for _, command := range command.Commands {
			fmt.Println(command.Name)
		}
	}

	// sort each flags
	sort.Sort(cli.FlagsByName(app.Flags))
	for _, command := range app.Commands {
		for _, subCommand := range command.Subcommands {
			sort.Sort(cli.FlagsByName(subCommand.Flags))
		}
		sort.Sort(cli.FlagsByName(command.Flags))
	}

	app.Run(os.Args)
}
