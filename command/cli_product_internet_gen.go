// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package command

import (
	"gopkg.in/urfave/cli.v2"
)

func init() {
	readParam := NewReadProductInternetParam()
	listParam := NewListProductInternetParam()

	cliCommand := &cli.Command{
		Name:    "product-internet",
		Aliases: []string{"internet-plan"},
		Usage:   "A manage commands of ProductInternet",
		Subcommands: []*cli.Command{
			{
				Name:      "read",
				Aliases:   []string{"r"},
				Usage:     "Read ProductInternet",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &readParam.Id,
					},
				},
				Action: func(c *cli.Context) error {

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() == 1 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := readParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), readParam)

					// Run command with params
					return ProductInternetRead(ctx, readParam)
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l", "ls", "find"},
				Usage:   "List ProductInternet",
				Flags: []cli.Flag{
					&cli.Int64SliceFlag{
						Name:  "id",
						Usage: "set filter by id(s)",
					},
					&cli.IntFlag{
						Name:        "from",
						Usage:       "set offset",
						Destination: &listParam.From,
					},
					&cli.IntFlag{
						Name:        "max",
						Usage:       "set limit",
						Destination: &listParam.Max,
					},
					&cli.StringSliceFlag{
						Name:  "sort",
						Usage: "set field(s) for sort",
					},
					&cli.StringSliceFlag{
						Name:  "name",
						Usage: "set filter by name(s)",
					},
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					listParam.Sort = c.StringSlice("sort")
					listParam.Name = c.StringSlice("name")
					listParam.Id = c.Int64Slice("id")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// Validate specific for each command params
					if errors := listParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), listParam)

					// Run command with params
					return ProductInternetList(ctx, listParam)
				},
			},
		},
	}

	Commands = append(Commands, cliCommand)
}
