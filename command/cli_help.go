package command

import (
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
)

type ResourceHelpValue struct {
	Name            string
	DisplayText     string
	VisibleCommands []*cli.Command
}

func (r *ResourceHelpValue) AppendCommand(c *cli.Command) {
	r.VisibleCommands = append(r.VisibleCommands, c)
}

type SortableResource struct {
	Category *schema.Category
	Command  *cli.Command
}

type SortableResources []SortableResource

func (s SortableResources) Len() int {
	return len(s)
}

func (s SortableResources) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortableResources) Less(i, j int) bool {

	if s[i].Category.Order == s[j].Category.Order {

		return s[i].Command.Name < s[j].Command.Name

	}
	return s[i].Category.Order < s[j].Category.Order
}

type SortableCommand struct {
	Category *schema.Category
}

type CategoryHelpValue struct {
	Name         string
	DisplayText  string
	VisibleFlags []cli.Flag
}

func (c *CategoryHelpValue) AppendFlags(f cli.Flag) {
	c.VisibleFlags = append(c.VisibleFlags, f)
}

var TopLevelHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} resource{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}command [command options] [arguments...]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}

VERSION:
   {{.Version}}{{end}}{{end}}{{if .Description}}

DESCRIPTION:
   {{.Description}}{{end}}{{if len .Authors}}

AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
   {{range $index, $author := .Authors}}{{if $index}}
   {{end}}{{$author}}{{end}}{{end}}{{if .ResourceCategories}}

COMMANDS:{{range .ResourceCategories}}{{range .VisibleCommands}}{{if ne .Name "help"}}
   {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}

GLOBAL OPTIONS:
   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{$option}}{{end}}{{end}}{{if .Copyright}}

COPYRIGHT:
   {{.Copyright}}{{end}}
`

// ResourceLevelHelpTemplate is using with `usacloud [resource] --help`
var ResourceLevelHelpTemplate = `NAME:
   {{.HelpName}} - {{.Usage}}

USAGE:
   {{.HelpName}} command{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}

COMMANDS:{{range .CommandCategories}}{{if .DisplayText}}
 === {{.DisplayText}} ==={{end}}{{range .VisibleCommands}}{{if ne .Name "help"}}
   {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}
{{end}}{{if .VisibleFlags}}
OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`

// CommandLevelHelpTemplate is using with `usacloud [resource] [command] --help`
var CommandLevelHelpTemplate = `NAME:
   {{.HelpName}} - {{.Usage}}

USAGE:
   {{.HelpName}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{if .Category}}

CATEGORY:
   {{.Category}}{{end}}{{if .Description}}

DESCRIPTION:
   {{.Description}}{{end}}{{if .VisibleFlags}}

OPTIONS:{{range .FlagCategories}}{{if .DisplayText}}
 === {{.DisplayText}} ===
   {{else}}
   {{end}}{{if .DisplayText}}{{range .VisibleFlags}}{{.}}
   {{end}}{{else}}{{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{end}}{{end}}
`
