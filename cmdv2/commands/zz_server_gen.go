// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/spf13/cobra"
)

var (
	serverListParam                     = params.NewListServerParam()
	serverBuildParam                    = params.NewBuildServerParam()
	serverReadParam                     = params.NewReadServerParam()
	serverUpdateParam                   = params.NewUpdateServerParam()
	serverDeleteParam                   = params.NewDeleteServerParam()
	serverPlanChangeParam               = params.NewPlanChangeServerParam()
	serverBootParam                     = params.NewBootServerParam()
	serverShutdownParam                 = params.NewShutdownServerParam()
	serverShutdownForceParam            = params.NewShutdownForceServerParam()
	serverResetParam                    = params.NewResetServerParam()
	serverWaitForBootParam              = params.NewWaitForBootServerParam()
	serverWaitForDownParam              = params.NewWaitForDownServerParam()
	serverSSHParam                      = params.NewSSHServerParam()
	serverSSHExecParam                  = params.NewSSHExecServerParam()
	serverScpParam                      = params.NewScpServerParam()
	serverVncParam                      = params.NewVncServerParam()
	serverVncInfoParam                  = params.NewVncInfoServerParam()
	serverVncSendParam                  = params.NewVncSendServerParam()
	serverVncSnapshotParam              = params.NewVncSnapshotServerParam()
	serverRemoteDesktopParam            = params.NewRemoteDesktopServerParam()
	serverRemoteDesktopInfoParam        = params.NewRemoteDesktopInfoServerParam()
	serverDiskInfoParam                 = params.NewDiskInfoServerParam()
	serverDiskConnectParam              = params.NewDiskConnectServerParam()
	serverDiskDisconnectParam           = params.NewDiskDisconnectServerParam()
	serverInterfaceInfoParam            = params.NewInterfaceInfoServerParam()
	serverInterfaceAddForInternetParam  = params.NewInterfaceAddForInternetServerParam()
	serverInterfaceAddForRouterParam    = params.NewInterfaceAddForRouterServerParam()
	serverInterfaceAddForSwitchParam    = params.NewInterfaceAddForSwitchServerParam()
	serverInterfaceAddDisconnectedParam = params.NewInterfaceAddDisconnectedServerParam()
	serverISOInfoParam                  = params.NewISOInfoServerParam()
	serverISOInsertParam                = params.NewISOInsertServerParam()
	serverISOEjectParam                 = params.NewISOEjectServerParam()
	serverMonitorCPUParam               = params.NewMonitorCPUServerParam()
	serverMonitorNicParam               = params.NewMonitorNicServerParam()
	serverMonitorDiskParam              = params.NewMonitorDiskServerParam()
	serverMaintenanceInfoParam          = params.NewMaintenanceInfoServerParam()
)

// serverCmd represents the command to manage SAKURA Cloud Server
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A manage commands of Server",
	Long:  `A manage commands of Server`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var serverListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find", "selector"},
	Short:   "List Server",
	Long:    `List Server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(serverListParam))
		return err
	},
}

func serverListCmdInit() {
	fs := serverListCmd.Flags()
	fs.IntVarP(&serverListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&serverListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringSliceVarP(&serverListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &serverListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&serverListParam.From, "from", "", 0, "set offset")
	fs.StringSliceVarP(&serverListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
}

var serverBuildCmd = &cobra.Command{
	Use: "build",

	Short: "Build Server",
	Long:  `Build Server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverBuildParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("build parameter: \n%s\n", debugMarshalIndent(serverBuildParam))
		return err
	},
}

func serverBuildCmdInit() {
	fs := serverBuildCmd.Flags()
	fs.StringVarP(&serverBuildParam.Ipaddress, "ipaddress", "", "", "set ipaddress")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &serverBuildParam.StartupScriptIds), "startup-script-ids", "", "set startup script ID(s)")
	fs.StringVarP(&serverBuildParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&serverBuildParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &serverBuildParam.SourceDiskId), "source-disk-id", "", "set source disk ID")
	fs.VarP(newIDValue(0, &serverBuildParam.DiskId), "disk-id", "", "set connect disk ID")
	fs.StringVarP(&serverBuildParam.Password, "password", "", "", "set password")
	fs.VarP(newIDValue(0, &serverBuildParam.ISOImageId), "iso-image-id", "", "set iso-image ID")
	fs.StringVarP(&serverBuildParam.SSHKeyMode, "ssh-key-mode", "", "", "ssh-key mode[none/id/generate/upload]")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &serverBuildParam.SSHKeyIds), "ssh-key-ids", "", "set ssh-key ID(s)")
	fs.StringVarP(&serverBuildParam.Name, "name", "", "", "set resource display name")
	fs.IntVarP(&serverBuildParam.Core, "core", "", 1, "set CPU core count")
	fs.StringVarP(&serverBuildParam.OsType, "os-type", "", "", "set source OS type")
	fs.StringVarP(&serverBuildParam.DiskPlan, "disk-plan", "", "ssd", "set disk plan('hdd' or 'ssd')")
	fs.StringVarP(&serverBuildParam.SSHKeyDescription, "ssh-key-description", "", "", "set ssh-key description")
	fs.StringVarP(&serverBuildParam.SSHKeyPrivateKeyOutput, "ssh-key-private-key-output", "", "", "set ssh-key privatekey output path")
	fs.BoolVarP(&serverBuildParam.DisableBootAfterCreate, "disable-boot-after-create", "", false, "boot after create")
	fs.VarP(newIDValue(0, &serverBuildParam.SourceArchiveId), "source-archive-id", "", "set source disk ID")
	fs.StringVarP(&serverBuildParam.Hostname, "hostname", "", "", "set hostname")
	fs.StringVarP(&serverBuildParam.SSHKeyName, "ssh-key-name", "", "", "set ssh-key name")
	fs.StringSliceVarP(&serverBuildParam.StartupScripts, "startup-scripts", "", []string{}, "set startup script(s)")
	fs.StringVarP(&serverBuildParam.Commitment, "commitment", "", "standard", "set plan of core assignment")
	fs.VarP(newIDValue(0, &serverBuildParam.PrivateHostId), "private-host-id", "", "set private-host-id")
	fs.StringVarP(&serverBuildParam.InterfaceDriver, "interface-driver", "", "virtio", "set interface driver[virtio/e1000]")
	fs.StringVarP(&serverBuildParam.DiskMode, "disk-mode", "", "create", "disk create mode[create/connect/diskless]")
	fs.IntVarP(&serverBuildParam.NwMasklen, "nw-masklen", "", 24, "set ipaddress  prefix")
	fs.BoolVarP(&serverBuildParam.UsKeyboard, "us-keyboard", "", false, "use us-keyboard")
	fs.StringVarP(&serverBuildParam.SSHKeyPassPhrase, "ssh-key-pass-phrase", "", "", "set ssh-key pass phrase")
	fs.StringSliceVarP(&serverBuildParam.SSHKeyPublicKeyFiles, "ssh-key-public-key-files", "", []string{}, "set ssh-key public key file")
	fs.StringVarP(&serverBuildParam.DiskConnection, "disk-connection", "", "virtio", "set disk connection('virtio' or 'ide')")
	fs.StringVarP(&serverBuildParam.NetworkMode, "network-mode", "", "shared", "network connection mode[shared/switch/disconnect/none]")
	fs.BoolVarP(&serverBuildParam.StartupScriptsEphemeral, "startup-scripts-ephemeral", "", true, "set startup script persist mode")
	fs.VarP(newIDValue(0, &serverBuildParam.PacketFilterId), "packet-filter-id", "", "set packet filter ID")
	fs.VarP(newIDValue(0, &serverBuildParam.SwitchId), "switch-id", "", "set connect switch ID")
	fs.BoolVarP(&serverBuildParam.DisablePasswordAuth, "disable-password-auth", "", false, "disable password auth on SSH")
	fs.BoolVarP(&serverBuildParam.SSHKeyEphemeral, "ssh-key-ephemeral", "", true, "set ssh-key persist mode")
	fs.VarP(newIDValue(0, &serverBuildParam.IconId), "icon-id", "", "set Icon ID")
	fs.IntVarP(&serverBuildParam.Memory, "memory", "", 1, "set memory size(GB)")
	fs.IntVarP(&serverBuildParam.DiskSize, "disk-size", "", 20, "set disk size(GB)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &serverBuildParam.DistantFrom), "distant-from", "", "set distant from disk IDs")
	fs.StringVarP(&serverBuildParam.DefaultRoute, "default-route", "", "", "set default gateway")
	fs.StringSliceVarP(&serverBuildParam.SSHKeyPublicKeys, "ssh-key-public-keys", "", []string{}, "set ssh-key public key ")
}

var serverReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read Server",
	Long:  `Read Server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(serverReadParam))
		return err
	},
}

func serverReadCmdInit() {
}

var serverUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update Server",
	Long:  `Update Server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("update parameter: \n%s\n", debugMarshalIndent(serverUpdateParam))
		return err
	},
}

func serverUpdateCmdInit() {
	fs := serverUpdateCmd.Flags()
	fs.StringVarP(&serverUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&serverUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&serverUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &serverUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.StringVarP(&serverUpdateParam.InterfaceDriver, "interface-driver", "", "virtio", "set interface driver[virtio/e1000]")
}

var serverDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete Server",
	Long:    `Delete Server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(serverDeleteParam))
		return err
	},
}

func serverDeleteCmdInit() {
	fs := serverDeleteCmd.Flags()
	fs.BoolVarP(&serverDeleteParam.Force, "force", "f", false, "forced-shutdown flag if server is running")
	fs.BoolVarP(&serverDeleteParam.WithoutDisk, "without-disk", "", false, "don't delete connected disks with server")
}

var serverPlanChangeCmd = &cobra.Command{
	Use: "plan-change",

	Short: "Change server plan(core/memory)",
	Long:  `Change server plan(core/memory)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverPlanChangeParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("plan-change parameter: \n%s\n", debugMarshalIndent(serverPlanChangeParam))
		return err
	},
}

func serverPlanChangeCmdInit() {
	fs := serverPlanChangeCmd.Flags()
	fs.IntVarP(&serverPlanChangeParam.Core, "core", "", 0, "set CPU core count")
	fs.IntVarP(&serverPlanChangeParam.Memory, "memory", "", 0, "set memory size(GB)")
	fs.StringVarP(&serverPlanChangeParam.Commitment, "commitment", "", "standard", "set plan of core assignment")
}

var serverBootCmd = &cobra.Command{
	Use:     "boot",
	Aliases: []string{"power-on"},
	Short:   "Boot Server",
	Long:    `Boot Server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverBootParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("boot parameter: \n%s\n", debugMarshalIndent(serverBootParam))
		return err
	},
}

func serverBootCmdInit() {
}

var serverShutdownCmd = &cobra.Command{
	Use:     "shutdown",
	Aliases: []string{"power-off"},
	Short:   "Shutdown Server",
	Long:    `Shutdown Server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverShutdownParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("shutdown parameter: \n%s\n", debugMarshalIndent(serverShutdownParam))
		return err
	},
}

func serverShutdownCmdInit() {
}

var serverShutdownForceCmd = &cobra.Command{
	Use:     "shutdown-force",
	Aliases: []string{"stop"},
	Short:   "ShutdownForce Server",
	Long:    `ShutdownForce Server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverShutdownForceParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("shutdown-force parameter: \n%s\n", debugMarshalIndent(serverShutdownForceParam))
		return err
	},
}

func serverShutdownForceCmdInit() {
}

var serverResetCmd = &cobra.Command{
	Use: "reset",

	Short: "Reset Server",
	Long:  `Reset Server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverResetParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("reset parameter: \n%s\n", debugMarshalIndent(serverResetParam))
		return err
	},
}

func serverResetCmdInit() {
}

var serverWaitForBootCmd = &cobra.Command{
	Use: "wait-for-boot",

	Short: "Wait until boot is completed",
	Long:  `Wait until boot is completed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverWaitForBootParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("wait-for-boot parameter: \n%s\n", debugMarshalIndent(serverWaitForBootParam))
		return err
	},
}

func serverWaitForBootCmdInit() {
}

var serverWaitForDownCmd = &cobra.Command{
	Use: "wait-for-down",

	Short: "Wait until shutdown is completed",
	Long:  `Wait until shutdown is completed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverWaitForDownParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("wait-for-down parameter: \n%s\n", debugMarshalIndent(serverWaitForDownParam))
		return err
	},
}

func serverWaitForDownCmdInit() {
}

var serverSSHCmd = &cobra.Command{
	Use: "ssh",

	Short: "Connect to server by SSH",
	Long:  `Connect to server by SSH`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverSSHParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("ssh parameter: \n%s\n", debugMarshalIndent(serverSSHParam))
		return err
	},
}

func serverSSHCmdInit() {
	fs := serverSSHCmd.Flags()
	fs.StringVarP(&serverSSHParam.Key, "key", "i", "", "private-key file path")
	fs.StringVarP(&serverSSHParam.User, "user", "l", "", "user name")
	fs.IntVarP(&serverSSHParam.Port, "port", "p", 22, "port")
	fs.StringVarP(&serverSSHParam.Password, "password", "", "", "password(or private-key pass phrase)")
	fs.BoolVarP(&serverSSHParam.Quiet, "quiet", "q", false, "disable information messages")
}

var serverSSHExecCmd = &cobra.Command{
	Use: "ssh-exec",

	Short: "Execute command on server connected by SSH",
	Long:  `Execute command on server connected by SSH`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverSSHExecParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("ssh-exec parameter: \n%s\n", debugMarshalIndent(serverSSHExecParam))
		return err
	},
}

func serverSSHExecCmdInit() {
	fs := serverSSHExecCmd.Flags()
	fs.StringVarP(&serverSSHExecParam.User, "user", "l", "", "user name")
	fs.IntVarP(&serverSSHExecParam.Port, "port", "p", 22, "port")
	fs.StringVarP(&serverSSHExecParam.Password, "password", "", "", "password(or private-key pass phrase)")
	fs.BoolVarP(&serverSSHExecParam.Quiet, "quiet", "q", false, "disable information messages")
	fs.StringVarP(&serverSSHExecParam.Key, "key", "i", "", "private-key file path")
}

var serverScpCmd = &cobra.Command{
	Use: "scp",

	Short: "Copy files/directories by SSH",
	Long:  `Copy files/directories by SSH`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverScpParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("scp parameter: \n%s\n", debugMarshalIndent(serverScpParam))
		return err
	},
}

func serverScpCmdInit() {
	fs := serverScpCmd.Flags()
	fs.BoolVarP(&serverScpParam.Recursive, "recursive", "r", false, "set recursive copy flag")
	fs.BoolVarP(&serverScpParam.Quiet, "quiet", "q", false, "disable information messages")
	fs.StringVarP(&serverScpParam.Key, "key", "i", "", "private-key file path")
	fs.StringVarP(&serverScpParam.User, "user", "l", "", "user name")
	fs.IntVarP(&serverScpParam.Port, "port", "p", 22, "port")
	fs.StringVarP(&serverScpParam.Password, "password", "", "", "password(or private-key pass phrase)")
}

var serverVncCmd = &cobra.Command{
	Use: "vnc",

	Short: "Open VNC client using the OS's default application",
	Long:  `Open VNC client using the OS's default application`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverVncParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("vnc parameter: \n%s\n", debugMarshalIndent(serverVncParam))
		return err
	},
}

func serverVncCmdInit() {
	fs := serverVncCmd.Flags()
	fs.BoolVarP(&serverVncParam.WaitForBoot, "wait-for-boot", "", false, "wait until the server starts up")
}

var serverVncInfoCmd = &cobra.Command{
	Use: "vnc-info",

	Short: "Show VNC proxy information",
	Long:  `Show VNC proxy information`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverVncInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("vnc-info parameter: \n%s\n", debugMarshalIndent(serverVncInfoParam))
		return err
	},
}

func serverVncInfoCmdInit() {
	fs := serverVncInfoCmd.Flags()
	fs.BoolVarP(&serverVncInfoParam.WaitForBoot, "wait-for-boot", "", false, "wait until the server starts up")
}

var serverVncSendCmd = &cobra.Command{
	Use: "vnc-send",

	Short: "Send keys over VNC connection",
	Long:  `Send keys over VNC connection`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverVncSendParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("vnc-send parameter: \n%s\n", debugMarshalIndent(serverVncSendParam))
		return err
	},
}

func serverVncSendCmdInit() {
	fs := serverVncSendCmd.Flags()
	fs.BoolVarP(&serverVncSendParam.UseUsKeyboard, "use-us-keyboard", "", false, "use US Keyboard")
	fs.BoolVarP(&serverVncSendParam.Debug, "debug", "d", false, "write debug info")
	fs.BoolVarP(&serverVncSendParam.WaitForBoot, "wait-for-boot", "", false, "wait until the server starts up")
	fs.StringVarP(&serverVncSendParam.Command, "command", "c", "", "command(compatible with HashiCorp Packer's boot_command)")
	fs.StringVarP(&serverVncSendParam.CommandFile, "command-file", "f", "", "command file(compatible with HashiCorp Packer's boot_command)")
}

var serverVncSnapshotCmd = &cobra.Command{
	Use: "vnc-snapshot",

	Short: "Capture VNC snapshot",
	Long:  `Capture VNC snapshot`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverVncSnapshotParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("vnc-snapshot parameter: \n%s\n", debugMarshalIndent(serverVncSnapshotParam))
		return err
	},
}

func serverVncSnapshotCmdInit() {
	fs := serverVncSnapshotCmd.Flags()
	fs.BoolVarP(&serverVncSnapshotParam.WaitForBoot, "wait-for-boot", "", false, "wait until the server starts up")
	fs.StringVarP(&serverVncSnapshotParam.OutputPath, "output-path", "", "", "snapshot output filepath")
}

var serverRemoteDesktopCmd = &cobra.Command{
	Use:     "remote-desktop",
	Aliases: []string{"rdp"},
	Short:   "Open RDP client using the OS's default application",
	Long:    `Open RDP client using the OS's default application`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverRemoteDesktopParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("remote-desktop parameter: \n%s\n", debugMarshalIndent(serverRemoteDesktopParam))
		return err
	},
}

func serverRemoteDesktopCmdInit() {
	fs := serverRemoteDesktopCmd.Flags()
	fs.StringVarP(&serverRemoteDesktopParam.User, "user", "l", "Administrator", "user name")
	fs.IntVarP(&serverRemoteDesktopParam.Port, "port", "p", 3389, "port")
}

var serverRemoteDesktopInfoCmd = &cobra.Command{
	Use:     "remote-desktop-info",
	Aliases: []string{"rdp-info"},
	Short:   "Show RDP information(.rdp)",
	Long:    `Show RDP information(.rdp)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverRemoteDesktopInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("remote-desktop-info parameter: \n%s\n", debugMarshalIndent(serverRemoteDesktopInfoParam))
		return err
	},
}

func serverRemoteDesktopInfoCmdInit() {
	fs := serverRemoteDesktopInfoCmd.Flags()
	fs.StringVarP(&serverRemoteDesktopInfoParam.User, "user", "l", "Administrator", "user name")
	fs.IntVarP(&serverRemoteDesktopInfoParam.Port, "port", "p", 3389, "port")
}

var serverDiskInfoCmd = &cobra.Command{
	Use:     "disk-info",
	Aliases: []string{"disk-list"},
	Short:   "Show information of disk(s) connected to server",
	Long:    `Show information of disk(s) connected to server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverDiskInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("disk-info parameter: \n%s\n", debugMarshalIndent(serverDiskInfoParam))
		return err
	},
}

func serverDiskInfoCmdInit() {
}

var serverDiskConnectCmd = &cobra.Command{
	Use: "disk-connect",

	Short: "Connect disk to server",
	Long:  `Connect disk to server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverDiskConnectParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("disk-connect parameter: \n%s\n", debugMarshalIndent(serverDiskConnectParam))
		return err
	},
}

func serverDiskConnectCmdInit() {
	fs := serverDiskConnectCmd.Flags()
	fs.VarP(newIDValue(0, &serverDiskConnectParam.DiskId), "disk-id", "", "set target disk ID")
}

var serverDiskDisconnectCmd = &cobra.Command{
	Use: "disk-disconnect",

	Short: "Disconnect disk from server",
	Long:  `Disconnect disk from server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverDiskDisconnectParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("disk-disconnect parameter: \n%s\n", debugMarshalIndent(serverDiskDisconnectParam))
		return err
	},
}

func serverDiskDisconnectCmdInit() {
	fs := serverDiskDisconnectCmd.Flags()
	fs.VarP(newIDValue(0, &serverDiskDisconnectParam.DiskId), "disk-id", "", "set target disk ID")
}

var serverInterfaceInfoCmd = &cobra.Command{
	Use:     "interface-info",
	Aliases: []string{"interface-list"},
	Short:   "Show information of NIC(s) connected to server",
	Long:    `Show information of NIC(s) connected to server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverInterfaceInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-info parameter: \n%s\n", debugMarshalIndent(serverInterfaceInfoParam))
		return err
	},
}

func serverInterfaceInfoCmdInit() {
}

var serverInterfaceAddForInternetCmd = &cobra.Command{
	Use: "interface-add-for-internet",

	Short: "Create and connect NIC connected to the internet",
	Long:  `Create and connect NIC connected to the internet`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverInterfaceAddForInternetParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-add-for-internet parameter: \n%s\n", debugMarshalIndent(serverInterfaceAddForInternetParam))
		return err
	},
}

func serverInterfaceAddForInternetCmdInit() {
	fs := serverInterfaceAddForInternetCmd.Flags()
	fs.BoolVarP(&serverInterfaceAddForInternetParam.WithoutDiskEdit, "without-disk-edit", "", false, "set skip edit-disk flag. if true, don't call DiskEdit API after interface added")
}

var serverInterfaceAddForRouterCmd = &cobra.Command{
	Use: "interface-add-for-router",

	Short: "Create and connect NIC connected to the router",
	Long:  `Create and connect NIC connected to the router`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverInterfaceAddForRouterParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-add-for-router parameter: \n%s\n", debugMarshalIndent(serverInterfaceAddForRouterParam))
		return err
	},
}

func serverInterfaceAddForRouterCmdInit() {
	fs := serverInterfaceAddForRouterCmd.Flags()
	fs.StringVarP(&serverInterfaceAddForRouterParam.DefaultRoute, "default-route", "", "", "set default gateway")
	fs.IntVarP(&serverInterfaceAddForRouterParam.NwMasklen, "nw-masklen", "", 24, "set ipaddress  prefix")
	fs.VarP(newIDValue(0, &serverInterfaceAddForRouterParam.SwitchId), "switch-id", "", "set connect switch(connected to router) ID")
	fs.BoolVarP(&serverInterfaceAddForRouterParam.WithoutDiskEdit, "without-disk-edit", "", false, "set skip edit-disk flag. if true, don't call DiskEdit API after interface added")
	fs.StringVarP(&serverInterfaceAddForRouterParam.Ipaddress, "ipaddress", "", "", "set ipaddress")
}

var serverInterfaceAddForSwitchCmd = &cobra.Command{
	Use: "interface-add-for-switch",

	Short: "Create and connect NIC connected to the switch",
	Long:  `Create and connect NIC connected to the switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverInterfaceAddForSwitchParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-add-for-switch parameter: \n%s\n", debugMarshalIndent(serverInterfaceAddForSwitchParam))
		return err
	},
}

func serverInterfaceAddForSwitchCmdInit() {
	fs := serverInterfaceAddForSwitchCmd.Flags()
	fs.BoolVarP(&serverInterfaceAddForSwitchParam.WithoutDiskEdit, "without-disk-edit", "", false, "set skip edit-disk flag. if true, don't call DiskEdit API after interface added")
	fs.StringVarP(&serverInterfaceAddForSwitchParam.Ipaddress, "ipaddress", "", "", "set ipaddress")
	fs.StringVarP(&serverInterfaceAddForSwitchParam.DefaultRoute, "default-route", "", "", "set default gateway")
	fs.IntVarP(&serverInterfaceAddForSwitchParam.NwMasklen, "nw-masklen", "", 24, "set ipaddress  prefix")
	fs.VarP(newIDValue(0, &serverInterfaceAddForSwitchParam.SwitchId), "switch-id", "", "set connect switch ID")
}

var serverInterfaceAddDisconnectedCmd = &cobra.Command{
	Use: "interface-add-disconnected",

	Short: "Create and connect a disconnected NIC",
	Long:  `Create and connect a disconnected NIC`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverInterfaceAddDisconnectedParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-add-disconnected parameter: \n%s\n", debugMarshalIndent(serverInterfaceAddDisconnectedParam))
		return err
	},
}

func serverInterfaceAddDisconnectedCmdInit() {
}

var serverISOInfoCmd = &cobra.Command{
	Use: "iso-info",

	Short: "Show information of ISO-Image inserted to server",
	Long:  `Show information of ISO-Image inserted to server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverISOInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("iso-info parameter: \n%s\n", debugMarshalIndent(serverISOInfoParam))
		return err
	},
}

func serverISOInfoCmdInit() {
}

var serverISOInsertCmd = &cobra.Command{
	Use: "iso-insert",

	Short: "Insert ISO-Image to server",
	Long:  `Insert ISO-Image to server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverISOInsertParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("iso-insert parameter: \n%s\n", debugMarshalIndent(serverISOInsertParam))
		return err
	},
}

func serverISOInsertCmdInit() {
	fs := serverISOInsertCmd.Flags()
	fs.StringVarP(&serverISOInsertParam.ISOFile, "iso-file", "", "", "set iso image file")
	fs.StringVarP(&serverISOInsertParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&serverISOInsertParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&serverISOInsertParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &serverISOInsertParam.IconId), "icon-id", "", "set Icon ID")
	fs.VarP(newIDValue(0, &serverISOInsertParam.ISOImageId), "iso-image-id", "", "set iso-image ID")
	fs.IntVarP(&serverISOInsertParam.Size, "size", "", 5, "set iso size(GB)")
}

var serverISOEjectCmd = &cobra.Command{
	Use: "iso-eject",

	Short: "Eject ISO-Image from server",
	Long:  `Eject ISO-Image from server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverISOEjectParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("iso-eject parameter: \n%s\n", debugMarshalIndent(serverISOEjectParam))
		return err
	},
}

func serverISOEjectCmdInit() {
}

var serverMonitorCPUCmd = &cobra.Command{
	Use: "monitor-cpu",

	Short: "Collect CPU monitor values",
	Long:  `Collect CPU monitor values`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverMonitorCPUParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-cpu parameter: \n%s\n", debugMarshalIndent(serverMonitorCPUParam))
		return err
	},
}

func serverMonitorCPUCmdInit() {
	fs := serverMonitorCPUCmd.Flags()
	fs.StringVarP(&serverMonitorCPUParam.Start, "start", "", "", "set start-time")
	fs.StringVarP(&serverMonitorCPUParam.End, "end", "", "", "set end-time")
	fs.StringVarP(&serverMonitorCPUParam.KeyFormat, "key-format", "", "sakuracloud.server.{{.ID}}.cpu", "set monitoring value key-format")
}

var serverMonitorNicCmd = &cobra.Command{
	Use: "monitor-nic",

	Short: "Collect NIC(s) monitor values",
	Long:  `Collect NIC(s) monitor values`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverMonitorNicParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-nic parameter: \n%s\n", debugMarshalIndent(serverMonitorNicParam))
		return err
	},
}

func serverMonitorNicCmdInit() {
	fs := serverMonitorNicCmd.Flags()
	fs.StringVarP(&serverMonitorNicParam.End, "end", "", "", "set end-time")
	fs.IntSliceVarP(&serverMonitorNicParam.Index, "index", "", []int{}, "target index(es)")
	fs.StringVarP(&serverMonitorNicParam.KeyFormat, "key-format", "", "sakuracloud.server.{{.ID}}.nic.{{.Index}}", "set monitoring value key-format")
	fs.StringVarP(&serverMonitorNicParam.Start, "start", "", "", "set start-time")
}

var serverMonitorDiskCmd = &cobra.Command{
	Use: "monitor-disk",

	Short: "Collect Disk(s) monitor values",
	Long:  `Collect Disk(s) monitor values`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverMonitorDiskParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-disk parameter: \n%s\n", debugMarshalIndent(serverMonitorDiskParam))
		return err
	},
}

func serverMonitorDiskCmdInit() {
	fs := serverMonitorDiskCmd.Flags()
	fs.StringVarP(&serverMonitorDiskParam.Start, "start", "", "", "set start-time")
	fs.StringVarP(&serverMonitorDiskParam.End, "end", "", "", "set end-time")
	fs.IntSliceVarP(&serverMonitorDiskParam.Index, "index", "", []int{}, "target index(es)")
	fs.StringVarP(&serverMonitorDiskParam.KeyFormat, "key-format", "", "sakuracloud.server.{{.ID}}.disk.{{.Index}}", "set monitoring value key-format")
}

var serverMaintenanceInfoCmd = &cobra.Command{
	Use: "maintenance-info",

	Short: "MaintenanceInfo Server",
	Long:  `MaintenanceInfo Server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serverMaintenanceInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("maintenance-info parameter: \n%s\n", debugMarshalIndent(serverMaintenanceInfoParam))
		return err
	},
}

func serverMaintenanceInfoCmdInit() {
}

func init() {
	parent := serverCmd

	serverListCmdInit()
	parent.AddCommand(serverListCmd)

	serverBuildCmdInit()
	parent.AddCommand(serverBuildCmd)

	serverReadCmdInit()
	parent.AddCommand(serverReadCmd)

	serverUpdateCmdInit()
	parent.AddCommand(serverUpdateCmd)

	serverDeleteCmdInit()
	parent.AddCommand(serverDeleteCmd)

	serverPlanChangeCmdInit()
	parent.AddCommand(serverPlanChangeCmd)

	serverBootCmdInit()
	parent.AddCommand(serverBootCmd)

	serverShutdownCmdInit()
	parent.AddCommand(serverShutdownCmd)

	serverShutdownForceCmdInit()
	parent.AddCommand(serverShutdownForceCmd)

	serverResetCmdInit()
	parent.AddCommand(serverResetCmd)

	serverWaitForBootCmdInit()
	parent.AddCommand(serverWaitForBootCmd)

	serverWaitForDownCmdInit()
	parent.AddCommand(serverWaitForDownCmd)

	serverSSHCmdInit()
	parent.AddCommand(serverSSHCmd)

	serverSSHExecCmdInit()
	parent.AddCommand(serverSSHExecCmd)

	serverScpCmdInit()
	parent.AddCommand(serverScpCmd)

	serverVncCmdInit()
	parent.AddCommand(serverVncCmd)

	serverVncInfoCmdInit()
	parent.AddCommand(serverVncInfoCmd)

	serverVncSendCmdInit()
	parent.AddCommand(serverVncSendCmd)

	serverVncSnapshotCmdInit()
	parent.AddCommand(serverVncSnapshotCmd)

	serverRemoteDesktopCmdInit()
	parent.AddCommand(serverRemoteDesktopCmd)

	serverRemoteDesktopInfoCmdInit()
	parent.AddCommand(serverRemoteDesktopInfoCmd)

	serverDiskInfoCmdInit()
	parent.AddCommand(serverDiskInfoCmd)

	serverDiskConnectCmdInit()
	parent.AddCommand(serverDiskConnectCmd)

	serverDiskDisconnectCmdInit()
	parent.AddCommand(serverDiskDisconnectCmd)

	serverInterfaceInfoCmdInit()
	parent.AddCommand(serverInterfaceInfoCmd)

	serverInterfaceAddForInternetCmdInit()
	parent.AddCommand(serverInterfaceAddForInternetCmd)

	serverInterfaceAddForRouterCmdInit()
	parent.AddCommand(serverInterfaceAddForRouterCmd)

	serverInterfaceAddForSwitchCmdInit()
	parent.AddCommand(serverInterfaceAddForSwitchCmd)

	serverInterfaceAddDisconnectedCmdInit()
	parent.AddCommand(serverInterfaceAddDisconnectedCmd)

	serverISOInfoCmdInit()
	parent.AddCommand(serverISOInfoCmd)

	serverISOInsertCmdInit()
	parent.AddCommand(serverISOInsertCmd)

	serverISOEjectCmdInit()
	parent.AddCommand(serverISOEjectCmd)

	serverMonitorCPUCmdInit()
	parent.AddCommand(serverMonitorCPUCmd)

	serverMonitorNicCmdInit()
	parent.AddCommand(serverMonitorNicCmd)

	serverMonitorDiskCmdInit()
	parent.AddCommand(serverMonitorDiskCmd)

	serverMaintenanceInfoCmdInit()
	parent.AddCommand(serverMaintenanceInfoCmd)

	rootCmd.AddCommand(parent)
}
