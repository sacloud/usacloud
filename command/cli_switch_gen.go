// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package command

import (
	"gopkg.in/urfave/cli.v2"
)

func init() {
	bridgeDisconnectParam := NewBridgeDisconnectSwitchParam()
	listParam := NewListSwitchParam()
	createParam := NewCreateSwitchParam()
	readParam := NewReadSwitchParam()
	updateParam := NewUpdateSwitchParam()
	deleteParam := NewDeleteSwitchParam()
	bridgeConnectParam := NewBridgeConnectSwitchParam()

	cliCommand := &cli.Command{
		Name:  "switch",
		Usage: "A manage commands of Switch",
		Subcommands: []*cli.Command{
			{
				Name:      "bridge-disconnect",
				Usage:     "BridgeDisconnect Switch",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &bridgeDisconnectParam.Id,
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
					if errors := bridgeDisconnectParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), bridgeDisconnectParam)

					// Run command with params
					return SwitchBridgeDisconnect(ctx, bridgeDisconnectParam)
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l", "ls", "find"},
				Usage:   "List Switch",
				Flags: []cli.Flag{
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
					&cli.IntFlag{
						Name:        "max",
						Usage:       "set limit",
						Destination: &listParam.Max,
					},
					&cli.StringSliceFlag{
						Name:  "sort",
						Usage: "set field(s) for sort",
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
					return SwitchList(ctx, listParam)
				},
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create Switch",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "description",
						Aliases:     []string{"desc"},
						Usage:       "set resource description",
						Destination: &createParam.Description,
					},
					&cli.StringSliceFlag{
						Name:  "tags",
						Usage: "set resource tags",
					},
					&cli.Int64Flag{
						Name:        "icon-id",
						Usage:       "set Icon ID",
						Destination: &createParam.IconId,
					},
					&cli.StringFlag{
						Name:        "name",
						Usage:       "[Required] set resource display name",
						Destination: &createParam.Name,
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
					return SwitchCreate(ctx, createParam)
				},
			},
			{
				Name:      "read",
				Aliases:   []string{"r"},
				Usage:     "Read Switch",
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
					return SwitchRead(ctx, readParam)
				},
			},
			{
				Name:      "update",
				Aliases:   []string{"u"},
				Usage:     "Update Switch",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "icon-id",
						Usage:       "set Icon ID",
						Destination: &updateParam.IconId,
					},
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
					&cli.StringFlag{
						Name:        "description",
						Aliases:     []string{"desc"},
						Usage:       "set resource description",
						Destination: &updateParam.Description,
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
					return SwitchUpdate(ctx, updateParam)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"d", "rm"},
				Usage:     "Delete Switch",
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
					return SwitchDelete(ctx, deleteParam)
				},
			},
			{
				Name:      "bridge-connect",
				Usage:     "BridgeConnect Switch",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &bridgeConnectParam.Id,
					},
					&cli.Int64Flag{
						Name:        "bridge-id",
						Usage:       "[Required] set bridge ID",
						Destination: &bridgeConnectParam.BridgeId,
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
					if errors := bridgeConnectParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), bridgeConnectParam)

					// Run command with params
					return SwitchBridgeConnect(ctx, bridgeConnectParam)
				},
			},
		},
	}

	Commands = append(Commands, cliCommand)
}
