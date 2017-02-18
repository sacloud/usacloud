// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package command

import (
	"gopkg.in/urfave/cli.v2"
)

func init() {
	isoEjectParam := NewIsoEjectServerParam()
	interfaceAddForInternetParam := NewInterfaceAddForInternetServerParam()
	interfaceAddDisconnectedParam := NewInterfaceAddDisconnectedServerParam()
	waitForDownParam := NewWaitForDownServerParam()
	isoInfoParam := NewIsoInfoServerParam()
	updateParam := NewUpdateServerParam()
	deleteParam := NewDeleteServerParam()
	sshParam := NewSshServerParam()
	isoInsertParam := NewIsoInsertServerParam()
	diskConnectParam := NewDiskConnectServerParam()
	interfaceAddForRouterParam := NewInterfaceAddForRouterServerParam()
	listParam := NewListServerParam()
	readParam := NewReadServerParam()
	waitForBootParam := NewWaitForBootServerParam()
	diskDisconnectParam := NewDiskDisconnectServerParam()
	interfaceInfoParam := NewInterfaceInfoServerParam()
	bootParam := NewBootServerParam()
	resetParam := NewResetServerParam()
	planChangeParam := NewPlanChangeServerParam()
	diskInfoParam := NewDiskInfoServerParam()
	interfaceAddForSwitchParam := NewInterfaceAddForSwitchServerParam()
	buildParam := NewBuildServerParam()
	shutdownParam := NewShutdownServerParam()

	cliCommand := &cli.Command{
		Name:  "server",
		Usage: "A manage commands of Server",
		Subcommands: []*cli.Command{
			{
				Name:      "iso-eject",
				Usage:     "IsoEject Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &isoEjectParam.Id,
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
					if errors := isoEjectParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), isoEjectParam)

					// Run command with params
					return ServerIsoEject(ctx, isoEjectParam)
				},
			},
			{
				Name:      "interface-add-for-internet",
				Usage:     "InterfaceAddForInternet Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "without-disk-edit",
						Usage:       "set skip edit-disk flag. if true, don't call DiskEdit API after interface added",
						Destination: &interfaceAddForInternetParam.WithoutDiskEdit,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &interfaceAddForInternetParam.Id,
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
					if errors := interfaceAddForInternetParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), interfaceAddForInternetParam)

					// Run command with params
					return ServerInterfaceAddForInternet(ctx, interfaceAddForInternetParam)
				},
			},
			{
				Name:      "interface-add-disconnected",
				Usage:     "InterfaceAddDisconnected Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &interfaceAddDisconnectedParam.Id,
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
					if errors := interfaceAddDisconnectedParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), interfaceAddDisconnectedParam)

					// Run command with params
					return ServerInterfaceAddDisconnected(ctx, interfaceAddDisconnectedParam)
				},
			},
			{
				Name:      "wait-for-down",
				Usage:     "WaitForDown Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &waitForDownParam.Id,
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
					if errors := waitForDownParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), waitForDownParam)

					// Run command with params
					return ServerWaitForDown(ctx, waitForDownParam)
				},
			},
			{
				Name:      "iso-info",
				Usage:     "IsoInfo Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &isoInfoParam.Id,
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
					if errors := isoInfoParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), isoInfoParam)

					// Run command with params
					return ServerIsoInfo(ctx, isoInfoParam)
				},
			},
			{
				Name:      "update",
				Aliases:   []string{"u"},
				Usage:     "Update Server",
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
					return ServerUpdate(ctx, updateParam)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"d", "rm"},
				Usage:     "Delete Server",
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
					return ServerDelete(ctx, deleteParam)
				},
			},
			{
				Name:      "ssh",
				Usage:     "Ssh Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &sshParam.Id,
					},
					&cli.StringFlag{
						Name:        "key",
						Usage:       "private-key file path",
						Destination: &sshParam.Key,
					},
					&cli.StringFlag{
						Name:        "user",
						Usage:       "user name",
						Destination: &sshParam.User,
					},
					&cli.IntFlag{
						Name:        "port",
						Usage:       "[Required] port",
						Value:       22,
						Destination: &sshParam.Port,
					},
					&cli.StringFlag{
						Name:        "password",
						Usage:       "password(or private-key pass phrase)",
						EnvVars:     []string{"SAKURACLOUD_SSH_PASSWORD"},
						Destination: &sshParam.Password,
					},
					&cli.StringFlag{
						Name:        "proxy",
						Usage:       "proxy server",
						Destination: &sshParam.Proxy,
					},
					&cli.BoolFlag{
						Name:        "open-pty",
						Usage:       "open pty",
						Value:       true,
						Destination: &sshParam.OpenPty,
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
					if errors := sshParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), sshParam)

					// Run command with params
					return ServerSsh(ctx, sshParam)
				},
			},
			{
				Name:      "iso-insert",
				Usage:     "IsoInsert Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "iso-image-id",
						Usage:       "set iso-image ID",
						Destination: &isoInsertParam.IsoImageId,
					},
					&cli.StringFlag{
						Name:        "name",
						Usage:       "set resource display name",
						Destination: &isoInsertParam.Name,
					},
					&cli.StringSliceFlag{
						Name:  "tags",
						Usage: "set resource tags",
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &isoInsertParam.Id,
					},
					&cli.StringFlag{
						Name:        "description",
						Aliases:     []string{"desc"},
						Usage:       "set resource description",
						Destination: &isoInsertParam.Description,
					},
					&cli.Int64Flag{
						Name:        "icon-id",
						Usage:       "set Icon ID",
						Destination: &isoInsertParam.IconId,
					},
					&cli.IntFlag{
						Name:        "size",
						Usage:       "set iso size(GB)",
						Value:       5,
						Destination: &isoInsertParam.Size,
					},
					&cli.StringFlag{
						Name:        "iso-file",
						Usage:       "set iso image file",
						Destination: &isoInsertParam.IsoFile,
					},
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					isoInsertParam.Tags = c.StringSlice("tags")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() == 1 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := isoInsertParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), isoInsertParam)

					// Run command with params
					return ServerIsoInsert(ctx, isoInsertParam)
				},
			},
			{
				Name:      "disk-connect",
				Usage:     "DiskConnect Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &diskConnectParam.Id,
					},
					&cli.Int64Flag{
						Name:        "disk-id",
						Usage:       "[Required] set target disk ID",
						Destination: &diskConnectParam.DiskId,
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
					if errors := diskConnectParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), diskConnectParam)

					// Run command with params
					return ServerDiskConnect(ctx, diskConnectParam)
				},
			},
			{
				Name:      "interface-add-for-router",
				Usage:     "InterfaceAddForRouter Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "ipaddress",
						Aliases:     []string{"ip"},
						Usage:       "set ipaddress",
						Destination: &interfaceAddForRouterParam.Ipaddress,
					},
					&cli.StringFlag{
						Name:        "default-route",
						Aliases:     []string{"gateway"},
						Usage:       "set default gateway",
						Destination: &interfaceAddForRouterParam.DefaultRoute,
					},
					&cli.IntFlag{
						Name:        "nw-masklen",
						Aliases:     []string{"network-masklen"},
						Usage:       "set ipaddress  prefix",
						Value:       24,
						Destination: &interfaceAddForRouterParam.NwMasklen,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &interfaceAddForRouterParam.Id,
					},
					&cli.BoolFlag{
						Name:        "without-disk-edit",
						Usage:       "set skip edit-disk flag. if true, don't call DiskEdit API after interface added",
						Destination: &interfaceAddForRouterParam.WithoutDiskEdit,
					},
					&cli.Int64Flag{
						Name:        "switch-id",
						Usage:       "[Required] set connect switch(connected to router) ID",
						Destination: &interfaceAddForRouterParam.SwitchId,
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
					if errors := interfaceAddForRouterParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), interfaceAddForRouterParam)

					// Run command with params
					return ServerInterfaceAddForRouter(ctx, interfaceAddForRouterParam)
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l", "ls", "find"},
				Usage:   "List Server",
				Flags: []cli.Flag{
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
					&cli.Int64SliceFlag{
						Name:  "id",
						Usage: "set filter by id(s)",
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
					return ServerList(ctx, listParam)
				},
			},
			{
				Name:      "read",
				Aliases:   []string{"r"},
				Usage:     "Read Server",
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
					return ServerRead(ctx, readParam)
				},
			},
			{
				Name:      "wait-for-boot",
				Usage:     "WaitForBoot Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &waitForBootParam.Id,
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
					if errors := waitForBootParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), waitForBootParam)

					// Run command with params
					return ServerWaitForBoot(ctx, waitForBootParam)
				},
			},
			{
				Name:      "disk-disconnect",
				Usage:     "DiskDisconnect Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &diskDisconnectParam.Id,
					},
					&cli.Int64Flag{
						Name:        "disk-id",
						Usage:       "[Required] set target disk ID",
						Destination: &diskDisconnectParam.DiskId,
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
					if errors := diskDisconnectParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), diskDisconnectParam)

					// Run command with params
					return ServerDiskDisconnect(ctx, diskDisconnectParam)
				},
			},
			{
				Name:      "interface-info",
				Usage:     "InterfaceInfo Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &interfaceInfoParam.Id,
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
					if errors := interfaceInfoParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), interfaceInfoParam)

					// Run command with params
					return ServerInterfaceInfo(ctx, interfaceInfoParam)
				},
			},
			{
				Name:      "boot",
				Aliases:   []string{"power-on"},
				Usage:     "Boot Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &bootParam.Id,
					},
					&cli.BoolFlag{
						Name:        "async",
						Usage:       "set async flag(if true,return with non block)",
						Destination: &bootParam.Async,
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
					if errors := bootParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), bootParam)

					// Run command with params
					return ServerBoot(ctx, bootParam)
				},
			},
			{
				Name:      "reset",
				Usage:     "Reset Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &resetParam.Id,
					},
					&cli.BoolFlag{
						Name:        "async",
						Usage:       "set async flag(if true,return with non block)",
						Destination: &resetParam.Async,
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
					if errors := resetParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), resetParam)

					// Run command with params
					return ServerReset(ctx, resetParam)
				},
			},
			{
				Name:      "plan-change",
				Usage:     "PlanChange Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &planChangeParam.Id,
					},
					&cli.IntFlag{
						Name:        "core",
						Usage:       "[Required] set CPU core count",
						Destination: &planChangeParam.Core,
					},
					&cli.IntFlag{
						Name:        "memory",
						Usage:       "[Required] set memory size(GB)",
						Destination: &planChangeParam.Memory,
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
					if errors := planChangeParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), planChangeParam)

					// Run command with params
					return ServerPlanChange(ctx, planChangeParam)
				},
			},
			{
				Name:      "disk-info",
				Usage:     "DiskInfo Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &diskInfoParam.Id,
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
					if errors := diskInfoParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), diskInfoParam)

					// Run command with params
					return ServerDiskInfo(ctx, diskInfoParam)
				},
			},
			{
				Name:      "interface-add-for-switch",
				Usage:     "InterfaceAddForSwitch Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "default-route",
						Aliases:     []string{"gateway"},
						Usage:       "set default gateway",
						Destination: &interfaceAddForSwitchParam.DefaultRoute,
					},
					&cli.IntFlag{
						Name:        "nw-masklen",
						Aliases:     []string{"network-masklen"},
						Usage:       "set ipaddress  prefix",
						Value:       24,
						Destination: &interfaceAddForSwitchParam.NwMasklen,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &interfaceAddForSwitchParam.Id,
					},
					&cli.BoolFlag{
						Name:        "without-disk-edit",
						Usage:       "set skip edit-disk flag. if true, don't call DiskEdit API after interface added",
						Destination: &interfaceAddForSwitchParam.WithoutDiskEdit,
					},
					&cli.Int64Flag{
						Name:        "switch-id",
						Usage:       "[Required] set connect switch ID",
						Destination: &interfaceAddForSwitchParam.SwitchId,
					},
					&cli.StringFlag{
						Name:        "ipaddress",
						Aliases:     []string{"ip"},
						Usage:       "set ipaddress",
						Destination: &interfaceAddForSwitchParam.Ipaddress,
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
					if errors := interfaceAddForSwitchParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), interfaceAddForSwitchParam)

					// Run command with params
					return ServerInterfaceAddForSwitch(ctx, interfaceAddForSwitchParam)
				},
			},
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "Build Server",
				Flags: []cli.Flag{
					&cli.Int64SliceFlag{
						Name:    "startup-script-ids",
						Aliases: []string{"note-ids"},
						Usage:   "set startup script ID(s)",
					},
					&cli.StringFlag{
						Name:        "name",
						Usage:       "[Required] set resource display name",
						Destination: &buildParam.Name,
					},
					&cli.StringSliceFlag{
						Name:  "tags",
						Usage: "set resource tags",
					},
					&cli.BoolFlag{
						Name:        "us-keyboard",
						Usage:       "use us-keyboard",
						Destination: &buildParam.UsKeyboard,
					},
					&cli.StringFlag{
						Name:        "network-mode",
						Usage:       "[Required] network connection mode[shared/switch/disconnect/none]",
						Value:       "shared",
						Destination: &buildParam.NetworkMode,
					},
					&cli.StringFlag{
						Name:        "hostname",
						Usage:       "set hostname",
						Destination: &buildParam.Hostname,
					},
					&cli.BoolFlag{
						Name:        "startup-scripts-ephemeral",
						Usage:       "set startup script persist mode(if true, script will delete after create server)",
						Value:       true,
						Destination: &buildParam.StartupScriptsEphemeral,
					},
					&cli.Int64SliceFlag{
						Name:  "ssh-key-ids",
						Usage: "set ssh-key ID(s)",
					},
					&cli.StringSliceFlag{
						Name:  "ssh-key-public-key-files",
						Usage: "set ssh-key public key file",
					},
					&cli.IntFlag{
						Name:        "core",
						Usage:       "[Required] set CPU core count",
						Value:       1,
						Destination: &buildParam.Core,
					},
					&cli.StringFlag{
						Name:        "os-type",
						Usage:       "set source OS type",
						Destination: &buildParam.OsType,
					},
					&cli.Int64Flag{
						Name:        "disk-id",
						Usage:       "set connect disk ID",
						Destination: &buildParam.DiskId,
					},
					&cli.BoolFlag{
						Name:        "use-nic-virtio",
						Usage:       "use virtio on nic",
						Value:       true,
						Destination: &buildParam.UseNicVirtio,
					},
					&cli.StringSliceFlag{
						Name:    "startup-scripts",
						Aliases: []string{"notes"},
						Usage:   "set startup script(s)",
					},
					&cli.StringFlag{
						Name:        "ssh-key-mode",
						Usage:       "ssh-key mode[none/id/generate/upload]",
						Destination: &buildParam.SshKeyMode,
					},
					&cli.IntFlag{
						Name:        "disk-size",
						Usage:       "set disk size(GB)",
						Value:       20,
						Destination: &buildParam.DiskSize,
					},
					&cli.Int64Flag{
						Name:        "source-disk-id",
						Usage:       "set source disk ID",
						Destination: &buildParam.SourceDiskId,
					},
					&cli.BoolFlag{
						Name:        "disable-password-auth",
						Aliases:     []string{"disable-pw-auth"},
						Usage:       "disable password auth on SSH",
						Destination: &buildParam.DisablePasswordAuth,
					},
					&cli.IntFlag{
						Name:        "nw-masklen",
						Aliases:     []string{"network-masklen"},
						Usage:       "set ipaddress  prefix",
						Value:       24,
						Destination: &buildParam.NwMasklen,
					},
					&cli.StringFlag{
						Name:        "ssh-key-name",
						Usage:       "set ssh-key name",
						Destination: &buildParam.SshKeyName,
					},
					&cli.Int64Flag{
						Name:        "iso-image-id",
						Usage:       "set iso-image ID",
						Destination: &buildParam.IsoImageId,
					},
					&cli.Int64Flag{
						Name:        "icon-id",
						Usage:       "set Icon ID",
						Destination: &buildParam.IconId,
					},
					&cli.StringFlag{
						Name:        "disk-mode",
						Usage:       "[Required] disk create mode[create/connect/diskless]",
						Value:       "create",
						Destination: &buildParam.DiskMode,
					},
					&cli.Int64SliceFlag{
						Name:  "distant-from",
						Usage: "set distant from disk IDs",
					},
					&cli.StringFlag{
						Name:        "ipaddress",
						Aliases:     []string{"ip"},
						Usage:       "set ipaddress",
						Destination: &buildParam.Ipaddress,
					},
					&cli.StringFlag{
						Name:        "default-route",
						Aliases:     []string{"gateway"},
						Usage:       "set default gateway",
						Destination: &buildParam.DefaultRoute,
					},
					&cli.StringFlag{
						Name:        "ssh-key-pass-phrase",
						Usage:       "set ssh-key pass phrase",
						Destination: &buildParam.SshKeyPassPhrase,
					},
					&cli.StringFlag{
						Name:        "ssh-key-private-key-output",
						Usage:       "set ssh-key privatekey output path",
						Destination: &buildParam.SshKeyPrivateKeyOutput,
					},
					&cli.IntFlag{
						Name:        "memory",
						Usage:       "[Required] set memory size(GB)",
						Value:       1,
						Destination: &buildParam.Memory,
					},
					&cli.StringFlag{
						Name:        "password",
						Usage:       "set password",
						Destination: &buildParam.Password,
					},
					&cli.StringFlag{
						Name:        "disk-plan",
						Usage:       "set disk plan('hdd' or 'ssd')",
						Value:       "ssd",
						Destination: &buildParam.DiskPlan,
					},
					&cli.StringFlag{
						Name:        "description",
						Aliases:     []string{"desc"},
						Usage:       "set resource description",
						Destination: &buildParam.Description,
					},
					&cli.Int64Flag{
						Name:        "packet-filter-id",
						Usage:       "set packet filter ID",
						Destination: &buildParam.PacketFilterId,
					},
					&cli.StringSliceFlag{
						Name:  "ssh-key-public-keys",
						Usage: "set ssh-key public key ",
					},
					&cli.BoolFlag{
						Name:        "ssh-key-ephemeral",
						Usage:       "set ssh-key persist mode(if true, script will delete after create server)",
						Value:       true,
						Destination: &buildParam.SshKeyEphemeral,
					},
					&cli.StringFlag{
						Name:        "disk-connection",
						Usage:       "set disk connection('virtio' or 'ide')",
						Value:       "virtio",
						Destination: &buildParam.DiskConnection,
					},
					&cli.Int64Flag{
						Name:        "source-archive-id",
						Usage:       "set source disk ID",
						Destination: &buildParam.SourceArchiveId,
					},
					&cli.BoolFlag{
						Name:        "disable-boot-after-create",
						Usage:       "boot after create",
						Value:       false,
						Destination: &buildParam.DisableBootAfterCreate,
					},
					&cli.Int64Flag{
						Name:        "switch-id",
						Usage:       "set connect switch ID",
						Destination: &buildParam.SwitchId,
					},
					&cli.StringFlag{
						Name:        "ssh-key-description",
						Usage:       "set ssh-key description",
						Destination: &buildParam.SshKeyDescription,
					},
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					buildParam.StartupScriptIds = c.Int64Slice("startup-script-ids")
					buildParam.Tags = c.StringSlice("tags")
					buildParam.SshKeyIds = c.Int64Slice("ssh-key-ids")
					buildParam.SshKeyPublicKeyFiles = c.StringSlice("ssh-key-public-key-files")
					buildParam.StartupScripts = c.StringSlice("startup-scripts")
					buildParam.DistantFrom = c.Int64Slice("distant-from")
					buildParam.SshKeyPublicKeys = c.StringSlice("ssh-key-public-keys")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// Validate specific for each command params
					if errors := buildParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), buildParam)

					// Run command with params
					return ServerBuild(ctx, buildParam)
				},
			},
			{
				Name:      "shutdown",
				Aliases:   []string{"power-off"},
				Usage:     "Shutdown Server",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "force",
						Usage:       "force shutdown flag",
						Destination: &shutdownParam.Force,
					},
					&cli.BoolFlag{
						Name:        "async",
						Usage:       "set async flag(if true,return with non block)",
						Destination: &shutdownParam.Async,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &shutdownParam.Id,
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
					if errors := shutdownParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), shutdownParam)

					// Run command with params
					return ServerShutdown(ctx, shutdownParam)
				},
			},
		},
	}

	Commands = append(Commands, cliCommand)
}
