// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package command

import (
	"gopkg.in/urfave/cli.v2"
)

func init() {
	listParam := NewListIconParam()
	createParam := NewCreateIconParam()
	readParam := NewReadIconParam()
	updateParam := NewUpdateIconParam()
	deleteParam := NewDeleteIconParam()

	cliCommand := &cli.Command{
		Name:  "icon",
		Usage: "A manage commands of Icon",
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l", "ls", "find"},
				Usage:   "List Icon",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:  "sort",
						Usage: "set field(s) for sort",
					},
					&cli.StringSliceFlag{
						Name:  "name",
						Usage: "set filter by name(s)",
					},
					&cli.Int64SliceFlag{
						Name:  "id",
						Usage: "set filter by id(s)",
					},
					&cli.IntFlag{
						Name:        "from",
						Usage:       "set offset",
						Destination: &listParam.From,
					},
					&cli.StringFlag{
						Name:        "scope",
						Usage:       "set filter by scope('user' or 'shared')",
						Destination: &listParam.Scope,
					},
					&cli.IntFlag{
						Name:        "max",
						Usage:       "set limit",
						Destination: &listParam.Max,
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
					return IconList(ctx, listParam)
				},
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create Icon",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Usage:       "[Required] set resource display name",
						Destination: &createParam.Name,
					},
					&cli.StringSliceFlag{
						Name:  "tags",
						Usage: "set resource tags",
					},
					&cli.StringFlag{
						Name:        "image",
						Usage:       "[Required] set icon image",
						Destination: &createParam.Image,
					},
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					createParam.Tags = c.StringSlice("tags")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// Validate specific for each command params
					if errors := createParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), createParam)

					// Run command with params
					return IconCreate(ctx, createParam)
				},
			},
			{
				Name:      "read",
				Aliases:   []string{"r"},
				Usage:     "Read Icon",
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
					return IconRead(ctx, readParam)
				},
			},
			{
				Name:      "update",
				Aliases:   []string{"u"},
				Usage:     "Update Icon",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &updateParam.Id,
					},
					&cli.StringFlag{
						Name:        "name",
						Usage:       "set resource display name",
						Destination: &updateParam.Name,
					},
					&cli.StringSliceFlag{
						Name:  "tags",
						Usage: "set resource tags",
					},
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					updateParam.Tags = c.StringSlice("tags")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() == 1 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := updateParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), updateParam)

					// Run command with params
					return IconUpdate(ctx, updateParam)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"d", "rm"},
				Usage:     "Delete Icon",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &deleteParam.Id,
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
					if errors := deleteParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), deleteParam)

					// Run command with params
					return IconDelete(ctx, deleteParam)
				},
			},
		},
	}

	Commands = append(Commands, cliCommand)
}
