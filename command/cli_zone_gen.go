// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
	"strings"
)

func init() {
	listParam := NewListZoneParam()
	readParam := NewReadZoneParam()

	cliCommand := &cli.Command{
		Name:  "zone",
		Usage: "A manage commands of Zone",
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l", "ls", "find"},
				Usage:   "List Zone",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "from",
						Usage:       "set offset",
						Destination: &listParam.From,
					},
					&cli.Int64SliceFlag{
						Name:  "id",
						Usage: "set filter by id(s)",
					},
					&cli.IntFlag{
						Name:        "max",
						Usage:       "set limit",
						Destination: &listParam.Max,
					},
					&cli.StringSliceFlag{
						Name:  "name",
						Usage: "set filter by name(s)",
					},
					&cli.StringSliceFlag{
						Name:  "sort",
						Usage: "set field(s) for sort",
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, listParam)

					// Set option values for slice
					listParam.Id = c.Int64Slice("id")
					listParam.Name = c.StringSlice("name")
					listParam.Sort = c.StringSlice("sort")

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								ZoneListCompleteArgs(ctx, listParam)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										ZoneListCompleteArgs(ctx, listParam)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									ZoneListCompleteFlags(ctx, listParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							ZoneListCompleteArgs(ctx, listParam)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					listParam.Id = c.Int64Slice("id")
					listParam.Name = c.StringSlice("name")
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
					return ZoneList(ctx, listParam)
				},
			},
			{
				Name:      "read",
				Aliases:   []string{"r"},
				Usage:     "Read Zone",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &readParam.Id,
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, readParam)

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								ZoneReadCompleteArgs(ctx, readParam)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										ZoneReadCompleteArgs(ctx, readParam)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									ZoneReadCompleteFlags(ctx, readParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							ZoneReadCompleteArgs(ctx, readParam)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() > 0 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := readParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), readParam)

					// Run command with params
					return ZoneRead(ctx, readParam)
				},
			},
		},
	}

	// build Category-Resource mapping
	appendResourceCategoryMap("zone", &schema.Category{
		Key:         "information",
		DisplayName: "Service/Product informations",
		Order:       90,
	})

	// build Category-Command mapping

	appendCommandCategoryMap("zone", "list", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("zone", "read", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})

	// build Category-Param mapping

	appendFlagCategoryMap("zone", "list", "from", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("zone", "list", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("zone", "list", "max", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("zone", "list", "name", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("zone", "list", "sort", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("zone", "read", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})

	// append command to GlobalContext
	Commands = append(Commands, cliCommand)
}
