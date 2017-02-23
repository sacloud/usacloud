// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package command

import (
	"gopkg.in/urfave/cli.v2"
)

func init() {
	listParam := NewListDNSParam()
	readParam := NewReadDNSParam()
	updateParam := NewUpdateDNSParam()
	recordAddParam := NewRecordAddDNSParam()
	createParam := NewCreateDNSParam()
	deleteParam := NewDeleteDNSParam()
	recordListParam := NewRecordListDNSParam()
	recordUpdateParam := NewRecordUpdateDNSParam()
	recordDeleteParam := NewRecordDeleteDNSParam()

	cliCommand := &cli.Command{
		Name:  "dns",
		Usage: "A manage commands of DNS",
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l", "ls", "find"},
				Usage:   "List DNS",
				Flags: []cli.Flag{
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
					&cli.Int64SliceFlag{
						Name:  "id",
						Usage: "set filter by id(s)",
					},
					&cli.IntFlag{
						Name:        "from",
						Usage:       "set offset",
						Destination: &listParam.From,
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
					return DNSList(ctx, listParam)
				},
			},
			{
				Name:      "read",
				Aliases:   []string{"r"},
				Usage:     "Read DNS",
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
					return DNSRead(ctx, readParam)
				},
			},
			{
				Name:      "update",
				Aliases:   []string{"u"},
				Usage:     "Update DNS",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &updateParam.Id,
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
					&cli.Int64Flag{
						Name:        "icon-id",
						Usage:       "set Icon ID",
						Destination: &updateParam.IconId,
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
					return DNSUpdate(ctx, updateParam)
				},
			},
			{
				Name:      "record-add",
				Usage:     "RecordAdd DNS",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Usage:       "[Required] set name",
						Destination: &recordAddParam.Name,
					},
					&cli.StringFlag{
						Name:        "value",
						Usage:       "set record data",
						Destination: &recordAddParam.Value,
					},
					&cli.IntFlag{
						Name:        "ttl",
						Usage:       "set ttl",
						Value:       3600,
						Destination: &recordAddParam.Ttl,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &recordAddParam.Id,
					},
					&cli.IntFlag{
						Name:        "mx-priority",
						Usage:       "set MX priority",
						Value:       10,
						Destination: &recordAddParam.MxPriority,
					},
					&cli.IntFlag{
						Name:        "srv-priority",
						Usage:       "set SRV priority",
						Value:       0,
						Destination: &recordAddParam.SrvPriority,
					},
					&cli.IntFlag{
						Name:        "srv-weight",
						Usage:       "set SRV priority",
						Value:       0,
						Destination: &recordAddParam.SrvWeight,
					},
					&cli.IntFlag{
						Name:        "srv-port",
						Usage:       "set SRV priority",
						Value:       0,
						Destination: &recordAddParam.SrvPort,
					},
					&cli.StringFlag{
						Name:        "srv-target",
						Usage:       "set SRV priority",
						Destination: &recordAddParam.SrvTarget,
					},
					&cli.StringFlag{
						Name:        "type",
						Usage:       "[Required] set record type[A/AAAA/NS/CNAME/MX/TXT/SRV]",
						Destination: &recordAddParam.Type,
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
					if errors := recordAddParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), recordAddParam)

					// Run command with params
					return DNSRecordAdd(ctx, recordAddParam)
				},
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create DNS",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Usage:       "[Required] set DNS zone name",
						Destination: &createParam.Name,
					},
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
					return DNSCreate(ctx, createParam)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"d", "rm"},
				Usage:     "Delete DNS",
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
					return DNSDelete(ctx, deleteParam)
				},
			},
			{
				Name:      "record-list",
				Usage:     "RecordList DNS",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &recordListParam.Id,
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
					if errors := recordListParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), recordListParam)

					// Run command with params
					return DNSRecordList(ctx, recordListParam)
				},
			},
			{
				Name:      "record-update",
				Usage:     "RecordUpdate DNS",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "srv-port",
						Usage:       "set SRV priority",
						Destination: &recordUpdateParam.SrvPort,
					},
					&cli.IntFlag{
						Name:        "index",
						Usage:       "[Required] index of target record",
						Destination: &recordUpdateParam.Index,
					},
					&cli.StringFlag{
						Name:        "name",
						Usage:       "set name",
						Destination: &recordUpdateParam.Name,
					},
					&cli.StringFlag{
						Name:        "value",
						Usage:       "set record data",
						Destination: &recordUpdateParam.Value,
					},
					&cli.IntFlag{
						Name:        "mx-priority",
						Usage:       "set MX priority",
						Destination: &recordUpdateParam.MxPriority,
					},
					&cli.IntFlag{
						Name:        "srv-weight",
						Usage:       "set SRV priority",
						Destination: &recordUpdateParam.SrvWeight,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &recordUpdateParam.Id,
					},
					&cli.StringFlag{
						Name:        "type",
						Usage:       "set record type[A/AAAA/NS/CNAME/MX/TXT/SRV]",
						Destination: &recordUpdateParam.Type,
					},
					&cli.IntFlag{
						Name:        "ttl",
						Usage:       "set ttl",
						Destination: &recordUpdateParam.Ttl,
					},
					&cli.IntFlag{
						Name:        "srv-priority",
						Usage:       "set SRV priority",
						Destination: &recordUpdateParam.SrvPriority,
					},
					&cli.StringFlag{
						Name:        "srv-target",
						Usage:       "set SRV priority",
						Destination: &recordUpdateParam.SrvTarget,
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
					if errors := recordUpdateParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), recordUpdateParam)

					// Run command with params
					return DNSRecordUpdate(ctx, recordUpdateParam)
				},
			},
			{
				Name:      "record-delete",
				Usage:     "RecordDelete DNS",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &recordDeleteParam.Id,
					},
					&cli.IntFlag{
						Name:        "index",
						Usage:       "[Required] index of target record",
						Destination: &recordDeleteParam.Index,
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
					if errors := recordDeleteParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), recordDeleteParam)

					// Run command with params
					return DNSRecordDelete(ctx, recordDeleteParam)
				},
			},
		},
	}

	Commands = append(Commands, cliCommand)
}
