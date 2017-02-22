// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package command

import (
	"gopkg.in/urfave/cli.v2"
)

func init() {
	updateParam := NewUpdateDiskParam()
	deleteParam := NewDeleteDiskParam()
	reinstallToBlankParam := NewReinstallToBlankDiskParam()
	serverDisconnectParam := NewServerDisconnectDiskParam()
	listParam := NewListDiskParam()
	readParam := NewReadDiskParam()
	waitForCopyParam := NewWaitForCopyDiskParam()
	reinstallFromArchiveParam := NewReinstallFromArchiveDiskParam()
	reinstallFromDiskParam := NewReinstallFromDiskDiskParam()
	serverConnectParam := NewServerConnectDiskParam()
	createParam := NewCreateDiskParam()
	editParam := NewEditDiskParam()

	cliCommand := &cli.Command{
		Name:  "disk",
		Usage: "A manage commands of Disk",
		Subcommands: []*cli.Command{
			{
				Name:      "update",
				Aliases:   []string{"u"},
				Usage:     "Update Disk",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "connection",
						Usage:       "set disk connection('virtio' or 'ide')",
						Destination: &updateParam.Connection,
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
					return DiskUpdate(ctx, updateParam)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"d", "rm"},
				Usage:     "Delete Disk",
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
					return DiskDelete(ctx, deleteParam)
				},
			},
			{
				Name:      "reinstall-to-blank",
				Usage:     "ReinstallToBlank Disk",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &reinstallToBlankParam.Id,
					},
					&cli.Int64SliceFlag{
						Name:  "distant-from",
						Usage: "set distant from disk IDs",
					},
					&cli.BoolFlag{
						Name:        "async",
						Usage:       "set async flag(if true,return with non block)",
						Destination: &reinstallToBlankParam.Async,
					},
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					reinstallToBlankParam.DistantFrom = c.Int64Slice("distant-from")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() == 1 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := reinstallToBlankParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), reinstallToBlankParam)

					// Run command with params
					return DiskReinstallToBlank(ctx, reinstallToBlankParam)
				},
			},
			{
				Name:      "server-disconnect",
				Usage:     "ServerDisconnect Disk",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &serverDisconnectParam.Id,
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
					if errors := serverDisconnectParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), serverDisconnectParam)

					// Run command with params
					return DiskServerDisconnect(ctx, serverDisconnectParam)
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l", "ls", "find"},
				Usage:   "List Disk",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "scope",
						Usage:       "set filter by scope('user' or 'shared')",
						Destination: &listParam.Scope,
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
					&cli.IntFlag{
						Name:        "max",
						Usage:       "set limit",
						Destination: &listParam.Max,
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
					return DiskList(ctx, listParam)
				},
			},
			{
				Name:      "read",
				Aliases:   []string{"r"},
				Usage:     "Read Disk",
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
					return DiskRead(ctx, readParam)
				},
			},
			{
				Name:      "wait-for-copy",
				Aliases:   []string{"wait"},
				Usage:     "WaitForCopy Disk",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &waitForCopyParam.Id,
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
					if errors := waitForCopyParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), waitForCopyParam)

					// Run command with params
					return DiskWaitForCopy(ctx, waitForCopyParam)
				},
			},
			{
				Name:      "reinstall-from-archive",
				Usage:     "ReinstallFromArchive Disk",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &reinstallFromArchiveParam.Id,
					},
					&cli.Int64Flag{
						Name:        "source-archive-id",
						Usage:       "[Required] set source disk ID",
						Destination: &reinstallFromArchiveParam.SourceArchiveId,
					},
					&cli.Int64SliceFlag{
						Name:  "distant-from",
						Usage: "set distant from disk IDs",
					},
					&cli.BoolFlag{
						Name:        "async",
						Usage:       "set async flag(if true,return with non block)",
						Destination: &reinstallFromArchiveParam.Async,
					},
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					reinstallFromArchiveParam.DistantFrom = c.Int64Slice("distant-from")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() == 1 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := reinstallFromArchiveParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), reinstallFromArchiveParam)

					// Run command with params
					return DiskReinstallFromArchive(ctx, reinstallFromArchiveParam)
				},
			},
			{
				Name:      "reinstall-from-disk",
				Usage:     "ReinstallFromDisk Disk",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "async",
						Usage:       "set async flag(if true,return with non block)",
						Destination: &reinstallFromDiskParam.Async,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &reinstallFromDiskParam.Id,
					},
					&cli.Int64Flag{
						Name:        "source-disk-id",
						Usage:       "[Required] set source disk ID",
						Destination: &reinstallFromDiskParam.SourceDiskId,
					},
					&cli.Int64SliceFlag{
						Name:  "distant-from",
						Usage: "set distant from disk IDs",
					},
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					reinstallFromDiskParam.DistantFrom = c.Int64Slice("distant-from")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() == 1 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := reinstallFromDiskParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), reinstallFromDiskParam)

					// Run command with params
					return DiskReinstallFromDisk(ctx, reinstallFromDiskParam)
				},
			},
			{
				Name:      "server-connect",
				Usage:     "ServerConnect Disk",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &serverConnectParam.Id,
					},
					&cli.Int64Flag{
						Name:        "server-id",
						Usage:       "[Required] set target server ID",
						Destination: &serverConnectParam.ServerId,
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
					if errors := serverConnectParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), serverConnectParam)

					// Run command with params
					return DiskServerConnect(ctx, serverConnectParam)
				},
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create Disk",
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
						Name:        "connection",
						Usage:       "[Required] set disk connection('virtio' or 'ide')",
						Value:       "virtio",
						Destination: &createParam.Connection,
					},
					&cli.Int64Flag{
						Name:        "source-disk-id",
						Usage:       "set source disk ID",
						Destination: &createParam.SourceDiskId,
					},
					&cli.BoolFlag{
						Name:        "async",
						Usage:       "set async flag(if true,return with non block)",
						Destination: &createParam.Async,
					},
					&cli.StringFlag{
						Name:        "name",
						Usage:       "[Required] set resource display name",
						Destination: &createParam.Name,
					},
					&cli.IntFlag{
						Name:        "size",
						Usage:       "[Required] set disk size(GB)",
						Value:       20,
						Destination: &createParam.Size,
					},
					&cli.Int64Flag{
						Name:        "source-archive-id",
						Usage:       "set source disk ID",
						Destination: &createParam.SourceArchiveId,
					},
					&cli.Int64SliceFlag{
						Name:  "distant-from",
						Usage: "set distant from disk IDs",
					},
					&cli.StringFlag{
						Name:        "plan",
						Usage:       "[Required] set disk plan('hdd' or 'ssd')",
						Value:       "ssd",
						Destination: &createParam.Plan,
					},
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					createParam.DistantFrom = c.Int64Slice("distant-from")
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
					return DiskCreate(ctx, createParam)
				},
			},
			{
				Name:      "edit",
				Aliases:   []string{"config"},
				Usage:     "Edit Disk",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "hostname",
						Usage:       "set hostname",
						Destination: &editParam.Hostname,
					},
					&cli.BoolFlag{
						Name:        "disable-password-auth",
						Aliases:     []string{"disable-pw-auth"},
						Usage:       "disable password auth on SSH",
						Destination: &editParam.DisablePasswordAuth,
					},
					&cli.StringFlag{
						Name:        "default-route",
						Aliases:     []string{"gateway"},
						Usage:       "set default gateway",
						Destination: &editParam.DefaultRoute,
					},
					&cli.IntFlag{
						Name:        "nw-masklen",
						Aliases:     []string{"network-masklen"},
						Usage:       "set ipaddress  prefix",
						Value:       24,
						Destination: &editParam.NwMasklen,
					},
					&cli.StringFlag{
						Name:        "ipaddress",
						Aliases:     []string{"ip"},
						Usage:       "set ipaddress",
						Destination: &editParam.Ipaddress,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &editParam.Id,
					},
					&cli.StringFlag{
						Name:        "password",
						Usage:       "set password",
						Destination: &editParam.Password,
					},
					&cli.Int64SliceFlag{
						Name:  "ssh-key-ids",
						Usage: "set ssh-key ID(s)",
					},
					&cli.Int64SliceFlag{
						Name:    "startup-script-ids",
						Aliases: []string{"note-ids"},
						Usage:   "set startup-script ID(s)",
					},
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					editParam.SshKeyIds = c.Int64Slice("ssh-key-ids")
					editParam.StartupScriptIds = c.Int64Slice("startup-script-ids")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() == 1 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := editParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), editParam)

					// Run command with params
					return DiskEdit(ctx, editParam)
				},
			},
		},
	}

	Commands = append(Commands, cliCommand)
}
