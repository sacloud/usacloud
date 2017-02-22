// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package command

import (
	"gopkg.in/urfave/cli.v2"
)

func init() {
	updateParam := NewUpdatePacketFilterParam()
	deleteParam := NewDeletePacketFilterParam()
	ruleListParam := NewRuleListPacketFilterParam()
	ruleAddParam := NewRuleAddPacketFilterParam()
	ruleUpdateParam := NewRuleUpdatePacketFilterParam()
	interfaceConnectParam := NewInterfaceConnectPacketFilterParam()
	interfaceDisconnectParam := NewInterfaceDisconnectPacketFilterParam()
	createParam := NewCreatePacketFilterParam()
	readParam := NewReadPacketFilterParam()
	ruleDeleteParam := NewRuleDeletePacketFilterParam()
	listParam := NewListPacketFilterParam()

	cliCommand := &cli.Command{
		Name:  "packet-filter",
		Usage: "A manage commands of PacketFilter",
		Subcommands: []*cli.Command{
			{
				Name:      "update",
				Aliases:   []string{"u"},
				Usage:     "Update PacketFilter",
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
					&cli.StringFlag{
						Name:        "description",
						Aliases:     []string{"desc"},
						Usage:       "set resource description",
						Destination: &updateParam.Description,
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
					if errors := updateParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), updateParam)

					// Run command with params
					return PacketFilterUpdate(ctx, updateParam)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"d", "rm"},
				Usage:     "Delete PacketFilter",
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
					return PacketFilterDelete(ctx, deleteParam)
				},
			},
			{
				Name:      "rule-list",
				Aliases:   []string{"rules"},
				Usage:     "RuleList PacketFilter",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &ruleListParam.Id,
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
					if errors := ruleListParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), ruleListParam)

					// Run command with params
					return PacketFilterRuleList(ctx, ruleListParam)
				},
			},
			{
				Name:      "rule-add",
				Usage:     "RuleAdd PacketFilter",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "protocol",
						Usage:       "set target protocol[tcp/udp/icmp/fragment/ip]",
						Destination: &ruleAddParam.Protocol,
					},
					&cli.StringFlag{
						Name:        "source-network",
						Usage:       "set source network[A.A.A.A] or [A.A.A.A/N (N=1..31)] or [A.A.A.A/M.M.M.M]",
						Destination: &ruleAddParam.SourceNetwork,
					},
					&cli.StringFlag{
						Name:        "source-port",
						Usage:       "set source port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]",
						Destination: &ruleAddParam.SourcePort,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &ruleAddParam.Id,
					},
					&cli.StringFlag{
						Name:        "destination-port",
						Aliases:     []string{"dest-port"},
						Usage:       "set destination port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]",
						Destination: &ruleAddParam.DestinationPort,
					},
					&cli.StringFlag{
						Name:        "action",
						Usage:       "set action[allow/deny]",
						Destination: &ruleAddParam.Action,
					},
					&cli.StringFlag{
						Name:        "description",
						Aliases:     []string{"desc"},
						Usage:       "set resource description",
						Destination: &ruleAddParam.Description,
					},
					&cli.IntFlag{
						Name:        "index",
						Usage:       "index to insert rule into",
						Value:       1,
						Destination: &ruleAddParam.Index,
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
					if errors := ruleAddParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), ruleAddParam)

					// Run command with params
					return PacketFilterRuleAdd(ctx, ruleAddParam)
				},
			},
			{
				Name:      "rule-update",
				Usage:     "RuleUpdate PacketFilter",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "protocol",
						Usage:       "set target protocol[tcp/udp/icmp/fragment/ip]",
						Destination: &ruleUpdateParam.Protocol,
					},
					&cli.StringFlag{
						Name:        "source-network",
						Usage:       "set source network[A.A.A.A] or [A.A.A.A/N (N=1..31)] or [A.A.A.A/M.M.M.M]",
						Destination: &ruleUpdateParam.SourceNetwork,
					},
					&cli.StringFlag{
						Name:        "destination-port",
						Aliases:     []string{"dest-port"},
						Usage:       "set destination port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]",
						Destination: &ruleUpdateParam.DestinationPort,
					},
					&cli.StringFlag{
						Name:        "description",
						Aliases:     []string{"desc"},
						Usage:       "set resource description",
						Destination: &ruleUpdateParam.Description,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &ruleUpdateParam.Id,
					},
					&cli.StringFlag{
						Name:        "source-port",
						Usage:       "set source port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]",
						Destination: &ruleUpdateParam.SourcePort,
					},
					&cli.StringFlag{
						Name:        "action",
						Usage:       "set action[allow/deny]",
						Destination: &ruleUpdateParam.Action,
					},
					&cli.IntFlag{
						Name:        "index",
						Usage:       "[Required] index of target rule",
						Destination: &ruleUpdateParam.Index,
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
					if errors := ruleUpdateParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), ruleUpdateParam)

					// Run command with params
					return PacketFilterRuleUpdate(ctx, ruleUpdateParam)
				},
			},
			{
				Name:      "interface-connect",
				Usage:     "InterfaceConnect PacketFilter",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &interfaceConnectParam.Id,
					},
					&cli.Int64Flag{
						Name:        "interface-id",
						Usage:       "[Required] set interface ID",
						Destination: &interfaceConnectParam.InterfaceId,
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
					if errors := interfaceConnectParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), interfaceConnectParam)

					// Run command with params
					return PacketFilterInterfaceConnect(ctx, interfaceConnectParam)
				},
			},
			{
				Name:      "interface-disconnect",
				Usage:     "InterfaceDisconnect PacketFilter",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &interfaceDisconnectParam.Id,
					},
					&cli.Int64Flag{
						Name:        "interface-id",
						Usage:       "[Required] set interface ID",
						Destination: &interfaceDisconnectParam.InterfaceId,
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
					if errors := interfaceDisconnectParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), interfaceDisconnectParam)

					// Run command with params
					return PacketFilterInterfaceDisconnect(ctx, interfaceDisconnectParam)
				},
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create PacketFilter",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "description",
						Aliases:     []string{"desc"},
						Usage:       "set resource description",
						Destination: &createParam.Description,
					},
					&cli.StringFlag{
						Name:        "name",
						Usage:       "[Required] set resource display name",
						Destination: &createParam.Name,
					},
				},
				Action: func(c *cli.Context) error {

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
					return PacketFilterCreate(ctx, createParam)
				},
			},
			{
				Name:      "read",
				Aliases:   []string{"r"},
				Usage:     "Read PacketFilter",
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
					return PacketFilterRead(ctx, readParam)
				},
			},
			{
				Name:      "rule-delete",
				Usage:     "RuleDelete PacketFilter",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &ruleDeleteParam.Id,
					},
					&cli.IntFlag{
						Name:        "index",
						Usage:       "[Required] index of target rule",
						Destination: &ruleDeleteParam.Index,
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
					if errors := ruleDeleteParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), ruleDeleteParam)

					// Run command with params
					return PacketFilterRuleDelete(ctx, ruleDeleteParam)
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l", "ls", "find"},
				Usage:   "List PacketFilter",
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
					listParam.Name = c.StringSlice("name")
					listParam.Id = c.Int64Slice("id")
					listParam.Sort = c.StringSlice("sort")

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
					return PacketFilterList(ctx, listParam)
				},
			},
		},
	}

	Commands = append(Commands, cliCommand)
}
