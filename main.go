package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/sacloud/usacloud/command"
	usacloud_cli "github.com/sacloud/usacloud/command/cli"
	"github.com/sacloud/usacloud/schema"
	"github.com/sacloud/usacloud/version"
	"gopkg.in/urfave/cli.v2"
)

var (
	appName      = "usacloud"
	appUsage     = "CLI client for SakuraCloud"
	appCopyright = "Copyright (C) 2017-2019 The Usacloud Authors"
)

func main() {

	app := &cli.App{
		Name:                  appName,
		Usage:                 appUsage,
		HelpName:              appName,
		Copyright:             appCopyright,
		EnableShellCompletion: true,
		Version:               version.FullVersion(),
		CommandNotFound:       cmdNotFound,
		Flags:                 command.GlobalFlags,
		Commands:              usacloud_cli.Commands,
	}

	cli.AppHelpTemplate = usacloud_cli.TopLevelHelpTemplate
	cli.SubcommandHelpTemplate = usacloud_cli.ResourceLevelHelpTemplate
	cli.CommandHelpTemplate = usacloud_cli.CommandLevelHelpTemplate
	cli.InitCompletionFlag.Hidden = true
	cli.HelpPrinter = getHelpPrinter(app, cli.HelpPrinter)

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}

func getHelpPrinter(app *cli.App, currentHelpPrinter func(io.Writer, string, interface{})) func(io.Writer, string, interface{}) {

	return func(w io.Writer, tmpl string, d interface{}) {
		var value interface{}
		switch args := d.(type) {
		case *cli.App:
			if args.Name == app.Name {
				// usacloud --help
				rawCommands := args.VisibleCommands()

				var sortedCommands = usacloud_cli.SortableResources{}
				for _, c := range rawCommands {
					category, ok := usacloud_cli.ResourceCategoryMap[c.Name]
					if !ok {
						category = schema.DefaultResourceCategory
					}
					sortedCommands = append(sortedCommands, usacloud_cli.SortableResource{
						Category: category,
						Command:  c,
					})
				}

				sort.Sort(sortedCommands)
				resourceHelpValue := []*usacloud_cli.ResourceHelpValue{}
				for _, c := range sortedCommands {
					resourceHelpValue = appendResourceByCategory(resourceHelpValue, c.Category, c.Command)
				}

				// define and create struct which has *cli.App + ResourceCategories
				value = struct {
					*cli.App
					ResourceCategories []*usacloud_cli.ResourceHelpValue
				}{args, resourceHelpValue}
			} else {
				// usacloud [resource] --help
				// keys := helpKeys[1:]
				helpKeys := strings.Split(args.Name, " ")
				r := helpKeys[1]

				rawCommands := args.VisibleCommands()
				resourceHelpValue := []*usacloud_cli.ResourceHelpValue{}

				for _, c := range rawCommands {
					category, ok := usacloud_cli.CommandCategoryMap[r][c.Name]
					if !ok {
						category = schema.DefaultCommandCategory
					}
					resourceHelpValue = appendResourceByCategory(resourceHelpValue, category, c)
				}

				// define and create struct which has *cli.App + ResourceCategories
				value = struct {
					*cli.App
					CommandCategories []*usacloud_cli.ResourceHelpValue
				}{args, resourceHelpValue}

			}

		case *cli.Command:
			var categoryHelpValues []*usacloud_cli.CategoryHelpValue
			// usacloud [resource] [command] --help
			helpKeys := strings.Split(args.FullName(), " ")
			r := helpKeys[0]
			c := ""
			if len(helpKeys) > 1 {
				c = helpKeys[1]
			}
			// build flag categories
			rawFlags := args.VisibleFlags()
			for _, flag := range rawFlags {

				found := false
				for _, name := range flag.Names() {

					if category, ok := usacloud_cli.FlagCategoryMap[r][c][name]; ok {
						categoryHelpValues = appendFlagByCategory(categoryHelpValues, category, flag)
						found = true
					}
					if found {
						break
					}
				}

				if !found {
					categoryHelpValues = appendFlagByCategory(categoryHelpValues, schema.DefaultParamCategory, flag)
				}
			}

			// define and create struct which has *cli.Command + FlagCategories
			value = struct {
				*cli.Command
				FlagCategories []*usacloud_cli.CategoryHelpValue
			}{args, categoryHelpValues}

			if len(categoryHelpValues) == 1 && categoryHelpValues[0].Name == "default" {
				categoryHelpValues[0].DisplayText = ""
			}
		}

		currentHelpPrinter(w, tmpl, value)
	}
}

func appendResourceByCategory(values []*usacloud_cli.ResourceHelpValue, category *schema.Category, comm *cli.Command) []*usacloud_cli.ResourceHelpValue {
	exists := false
	for _, catHelp := range values {
		if catHelp.Name == category.Key {
			exists = true
			catHelp.AppendCommand(comm)
		}
	}

	if !exists {
		values = append(values, &usacloud_cli.ResourceHelpValue{
			Name:            category.Key,
			DisplayText:     category.DisplayName,
			VisibleCommands: []*cli.Command{comm},
		})
	}

	return values
}

func appendFlagByCategory(values []*usacloud_cli.CategoryHelpValue, category *schema.Category, flag cli.Flag) []*usacloud_cli.CategoryHelpValue {
	exists := false
	for _, catHelp := range values {
		if catHelp.Name == category.Key {
			exists = true
			catHelp.AppendFlags(flag)
		}
	}
	if !exists {
		values = append(values, &usacloud_cli.CategoryHelpValue{
			Name:         category.Key,
			DisplayText:  category.DisplayName,
			VisibleFlags: []cli.Flag{flag},
		})
	}

	return values
}

func cmdNotFound(c *cli.Context, command string) {
	fmt.Fprintf(
		os.Stderr,
		"%s: '%s' is not a %s command. See '%s --help'\n",
		c.App.Name,
		command,
		c.App.Name,
		c.App.Name,
	)
	os.Exit(1)
}
