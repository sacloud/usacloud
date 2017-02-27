package main

import (
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/schema"
	"github.com/sacloud/usacloud/version"
	"gopkg.in/urfave/cli.v2"
	"io"
	"os"
	"sort"
	"strings"
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
	//app.ShellComplete = func(c *cli.Context) {
	//	// This will complete if no args are passed
	//	if c.NArg() > 0 {
	//		return
	//	}
	//
	//	for _, command := range command.Commands {
	//		fmt.Println(command.Name)
	//	}
	//}

	cli.AppHelpTemplate = command.TopLevelHelpTemplate
	cli.SubcommandHelpTemplate = command.ResourceLevelHelpTemplate
	cli.CommandHelpTemplate = command.CommandLevelHelpTemplate

	cli.InitCompletionFlag.Hidden = true

	currentHelpPrinter := cli.HelpPrinter
	cli.HelpPrinter = func(w io.Writer, templ string, d interface{}) {

		var value interface{}
		switch args := d.(type) {
		case *cli.App:
			if args.Name == app.Name {
				// usacloud --help
				rawCommands := args.VisibleCommands()

				var sortedCommands = command.SortableResources{}
				for _, c := range rawCommands {
					category, ok := command.ResourceCategoryMap[c.Name]
					if !ok {
						category = schema.DefaultResourceCategory
					}
					sortedCommands = append(sortedCommands, command.SortableResource{
						Category: category,
						Command:  c,
					})
				}

				sort.Sort(sortedCommands)
				resourceHelpValue := []*command.ResourceHelpValue{}
				for _, c := range sortedCommands {
					resourceHelpValue = appendResourceByCategory(resourceHelpValue, c.Category, c.Command)
				}

				// define and create struct which has *cli.App + ResourceCategories
				value = struct {
					*cli.App
					ResourceCategories []*command.ResourceHelpValue
				}{args, resourceHelpValue}
			} else {
				// usacloud [resource] --help
				// keys := helpKeys[1:]
				helpKeys := strings.Split(args.Name, " ")
				r := helpKeys[1]

				rawCommands := args.VisibleCommands()
				resourceHelpValue := []*command.ResourceHelpValue{}

				for _, c := range rawCommands {
					category, ok := command.CommandCategoryMap[r][c.Name]
					if !ok {
						category = schema.DefaultCommandCategory
					}
					resourceHelpValue = appendResourceByCategory(resourceHelpValue, category, c)
				}

				// define and create struct which has *cli.App + ResourceCategories
				value = struct {
					*cli.App
					CommandCategories []*command.ResourceHelpValue
				}{args, resourceHelpValue}

			}

		case *cli.Command:
			var categoryHelpValues []*command.CategoryHelpValue
			// usacloud [resource] [command] --help
			helpKeys := strings.Split(args.FullName(), " ")
			r := helpKeys[0]
			c := helpKeys[1]
			// build flag categories
			rawFlags := args.VisibleFlags()
			for _, flag := range rawFlags {

				found := false
				for _, name := range flag.Names() {

					if category, ok := command.FlagCategoryMap[r][c][name]; ok {
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
				FlagCategories []*command.CategoryHelpValue
			}{args, categoryHelpValues}

			if len(categoryHelpValues) == 1 && categoryHelpValues[0].Name == "default" {
				categoryHelpValues[0].DisplayText = ""
			}
		}

		currentHelpPrinter(w, templ, value)
	}

	app.Run(os.Args)
}

func appendResourceByCategory(values []*command.ResourceHelpValue, category *schema.Category, comm *cli.Command) []*command.ResourceHelpValue {
	exists := false
	for _, catHelp := range values {
		if catHelp.Name == category.Key {
			exists = true
			catHelp.AppendCommand(comm)
		}
	}

	if !exists {
		values = append(values, &command.ResourceHelpValue{
			Name:            category.Key,
			DisplayText:     category.DisplayName,
			VisibleCommands: []*cli.Command{comm},
		})
	}

	return values
}

func appendFlagByCategory(values []*command.CategoryHelpValue, category *schema.Category, flag cli.Flag) []*command.CategoryHelpValue {
	exists := false
	for _, catHelp := range values {
		if catHelp.Name == category.Key {
			exists = true
			catHelp.AppendFlags(flag)
		}
	}
	if !exists {
		values = append(values, &command.CategoryHelpValue{
			Name:         category.Key,
			DisplayText:  category.DisplayName,
			VisibleFlags: []cli.Flag{flag},
		})
	}

	return values
}
