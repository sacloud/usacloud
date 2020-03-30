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
	databaseListParam                  = params.NewListDatabaseParam()
	databaseCreateParam                = params.NewCreateDatabaseParam()
	databaseReadParam                  = params.NewReadDatabaseParam()
	databaseUpdateParam                = params.NewUpdateDatabaseParam()
	databaseDeleteParam                = params.NewDeleteDatabaseParam()
	databaseBootParam                  = params.NewBootDatabaseParam()
	databaseShutdownParam              = params.NewShutdownDatabaseParam()
	databaseShutdownForceParam         = params.NewShutdownForceDatabaseParam()
	databaseResetParam                 = params.NewResetDatabaseParam()
	databaseWaitForBootParam           = params.NewWaitForBootDatabaseParam()
	databaseWaitForDownParam           = params.NewWaitForDownDatabaseParam()
	databaseBackupInfoParam            = params.NewBackupInfoDatabaseParam()
	databaseBackupCreateParam          = params.NewBackupCreateDatabaseParam()
	databaseBackupRestoreParam         = params.NewBackupRestoreDatabaseParam()
	databaseBackupLockParam            = params.NewBackupLockDatabaseParam()
	databaseBackupUnlockParam          = params.NewBackupUnlockDatabaseParam()
	databaseBackupRemoveParam          = params.NewBackupRemoveDatabaseParam()
	databaseCloneParam                 = params.NewCloneDatabaseParam()
	databaseReplicaCreateParam         = params.NewReplicaCreateDatabaseParam()
	databaseMonitorCPUParam            = params.NewMonitorCPUDatabaseParam()
	databaseMonitorMemoryParam         = params.NewMonitorMemoryDatabaseParam()
	databaseMonitorNicParam            = params.NewMonitorNicDatabaseParam()
	databaseMonitorSystemDiskParam     = params.NewMonitorSystemDiskDatabaseParam()
	databaseMonitorBackupDiskParam     = params.NewMonitorBackupDiskDatabaseParam()
	databaseMonitorSystemDiskSizeParam = params.NewMonitorSystemDiskSizeDatabaseParam()
	databaseMonitorBackupDiskSizeParam = params.NewMonitorBackupDiskSizeDatabaseParam()
	databaseLogsParam                  = params.NewLogsDatabaseParam()
)

// databaseCmd represents the command to manage SAKURA Cloud Database
var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "A manage commands of Database",
	Long:  `A manage commands of Database`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var databaseListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find", "selector"},
	Short:   "List Database",
	Long:    `List Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(databaseListParam))
		return err
	},
}

func databaseListCmdInit() {
	fs := databaseListCmd.Flags()
	fs.StringSliceVarP(&databaseListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.StringSliceVarP(&databaseListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringSliceVarP(&databaseListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &databaseListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&databaseListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&databaseListParam.Max, "max", "", 0, "set limit")
}

var databaseCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create Database",
	Long:  `Create Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("create parameter: \n%s\n", debugMarshalIndent(databaseCreateParam))
		return err
	},
}

func databaseCreateCmdInit() {
	fs := databaseCreateCmd.Flags()
	fs.BoolVarP(&databaseCreateParam.EnableWebUi, "enable-web-ui", "", false, "enable web-ui")
	fs.BoolVarP(&databaseCreateParam.EnableBackup, "enable-backup", "", false, "enable backup")
	fs.StringVarP(&databaseCreateParam.Ipaddress1, "ipaddress-1", "", "", "set ipaddress(#1)")
	fs.IntVarP(&databaseCreateParam.NwMaskLen, "nw-mask-len", "", 0, "set network mask length")
	fs.StringVarP(&databaseCreateParam.Name, "name", "", "", "set resource display name")
	fs.VarP(newIDValue(0, &databaseCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.StringVarP(&databaseCreateParam.Database, "database", "", "", "set database type[postgresql/mariadb]")
	fs.StringVarP(&databaseCreateParam.ReplicaUserPassword, "replica-user-password", "", "", "set database replica user password")
	fs.StringVarP(&databaseCreateParam.DefaultRoute, "default-route", "", "", "set default route")
	fs.StringSliceVarP(&databaseCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.IntVarP(&databaseCreateParam.Plan, "plan", "", 10, "set plan[10/30/90/240/500/1000]")
	fs.StringVarP(&databaseCreateParam.Password, "password", "", "", "set database default user password")
	fs.IntVarP(&databaseCreateParam.Port, "port", "", 0, "set database port")
	fs.VarP(newIDValue(0, &databaseCreateParam.SwitchId), "switch-id", "", "set connect switch ID")
	fs.StringSliceVarP(&databaseCreateParam.SourceNetworks, "source-networks", "", []string{}, "set network of allow connection")
	fs.StringSliceVarP(&databaseCreateParam.BackupWeekdays, "backup-weekdays", "", []string{"all"}, "set backup target weekdays[all or mon/tue/wed/thu/fri/sat/sun]")
	fs.StringVarP(&databaseCreateParam.BackupTime, "backup-time", "", "", "set backup start time")
	fs.StringVarP(&databaseCreateParam.Description, "description", "", "", "set resource description")
	fs.StringVarP(&databaseCreateParam.Username, "username", "", "", "set database default user name")
}

var databaseReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read Database",
	Long:  `Read Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(databaseReadParam))
		return err
	},
}

func databaseReadCmdInit() {
}

var databaseUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update Database",
	Long:  `Update Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("update parameter: \n%s\n", debugMarshalIndent(databaseUpdateParam))
		return err
	},
}

func databaseUpdateCmdInit() {
	fs := databaseUpdateCmd.Flags()
	fs.StringVarP(&databaseUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringSliceVarP(&databaseUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.StringSliceVarP(&databaseUpdateParam.SourceNetworks, "source-networks", "", []string{}, "set network of allow connection")
	fs.StringSliceVarP(&databaseUpdateParam.BackupWeekdays, "backup-weekdays", "", []string{"all"}, "set backup target weekdays[all or mon/tue/wed/thu/fri/sat/sun]")
	fs.BoolVarP(&databaseUpdateParam.EnableReplication, "enable-replication", "", false, "enable replication")
	fs.IntVarP(&databaseUpdateParam.Port, "port", "", 0, "set database port")
	fs.BoolVarP(&databaseUpdateParam.EnableWebUi, "enable-web-ui", "", false, "enable web-ui")
	fs.BoolVarP(&databaseUpdateParam.EnableBackup, "enable-backup", "", false, "enable backup")
	fs.StringVarP(&databaseUpdateParam.BackupTime, "backup-time", "", "", "set backup start time")
	fs.StringVarP(&databaseUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringVarP(&databaseUpdateParam.Password, "password", "", "", "set database default user password")
	fs.StringVarP(&databaseUpdateParam.ReplicaUserPassword, "replica-user-password", "", "", "set database replica user password")
	fs.VarP(newIDValue(0, &databaseUpdateParam.IconId), "icon-id", "", "set Icon ID")
}

var databaseDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete Database",
	Long:    `Delete Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(databaseDeleteParam))
		return err
	},
}

func databaseDeleteCmdInit() {
	fs := databaseDeleteCmd.Flags()
	fs.BoolVarP(&databaseDeleteParam.Force, "force", "f", false, "forced-shutdown flag if database is running")
}

var databaseBootCmd = &cobra.Command{
	Use:     "boot",
	Aliases: []string{"power-on"},
	Short:   "Boot Database",
	Long:    `Boot Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseBootParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("boot parameter: \n%s\n", debugMarshalIndent(databaseBootParam))
		return err
	},
}

func databaseBootCmdInit() {
}

var databaseShutdownCmd = &cobra.Command{
	Use:     "shutdown",
	Aliases: []string{"power-off"},
	Short:   "Shutdown Database",
	Long:    `Shutdown Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseShutdownParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("shutdown parameter: \n%s\n", debugMarshalIndent(databaseShutdownParam))
		return err
	},
}

func databaseShutdownCmdInit() {
}

var databaseShutdownForceCmd = &cobra.Command{
	Use:     "shutdown-force",
	Aliases: []string{"stop"},
	Short:   "ShutdownForce Database",
	Long:    `ShutdownForce Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseShutdownForceParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("shutdown-force parameter: \n%s\n", debugMarshalIndent(databaseShutdownForceParam))
		return err
	},
}

func databaseShutdownForceCmdInit() {
}

var databaseResetCmd = &cobra.Command{
	Use: "reset",

	Short: "Reset Database",
	Long:  `Reset Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseResetParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("reset parameter: \n%s\n", debugMarshalIndent(databaseResetParam))
		return err
	},
}

func databaseResetCmdInit() {
}

var databaseWaitForBootCmd = &cobra.Command{
	Use: "wait-for-boot",

	Short: "Wait until boot is completed",
	Long:  `Wait until boot is completed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseWaitForBootParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("wait-for-boot parameter: \n%s\n", debugMarshalIndent(databaseWaitForBootParam))
		return err
	},
}

func databaseWaitForBootCmdInit() {
}

var databaseWaitForDownCmd = &cobra.Command{
	Use: "wait-for-down",

	Short: "Wait until shutdown is completed",
	Long:  `Wait until shutdown is completed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseWaitForDownParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("wait-for-down parameter: \n%s\n", debugMarshalIndent(databaseWaitForDownParam))
		return err
	},
}

func databaseWaitForDownCmdInit() {
}

var databaseBackupInfoCmd = &cobra.Command{
	Use:     "backup-info",
	Aliases: []string{"backups", "backup-list"},
	Short:   "Show information of backup",
	Long:    `Show information of backup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseBackupInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("backup-info parameter: \n%s\n", debugMarshalIndent(databaseBackupInfoParam))
		return err
	},
}

func databaseBackupInfoCmdInit() {
}

var databaseBackupCreateCmd = &cobra.Command{
	Use: "backup-create",

	Short: "Make new database backup",
	Long:  `Make new database backup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseBackupCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("backup-create parameter: \n%s\n", debugMarshalIndent(databaseBackupCreateParam))
		return err
	},
}

func databaseBackupCreateCmdInit() {
}

var databaseBackupRestoreCmd = &cobra.Command{
	Use: "backup-restore",

	Short: "Restore database from backup",
	Long:  `Restore database from backup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseBackupRestoreParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("backup-restore parameter: \n%s\n", debugMarshalIndent(databaseBackupRestoreParam))
		return err
	},
}

func databaseBackupRestoreCmdInit() {
	fs := databaseBackupRestoreCmd.Flags()
	fs.IntVarP(&databaseBackupRestoreParam.Index, "index", "", 0, "index of target backup")
}

var databaseBackupLockCmd = &cobra.Command{
	Use: "backup-lock",

	Short: "Lock backup",
	Long:  `Lock backup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseBackupLockParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("backup-lock parameter: \n%s\n", debugMarshalIndent(databaseBackupLockParam))
		return err
	},
}

func databaseBackupLockCmdInit() {
	fs := databaseBackupLockCmd.Flags()
	fs.IntVarP(&databaseBackupLockParam.Index, "index", "", 0, "index of target backup")
}

var databaseBackupUnlockCmd = &cobra.Command{
	Use: "backup-unlock",

	Short: "Unlock backup",
	Long:  `Unlock backup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseBackupUnlockParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("backup-unlock parameter: \n%s\n", debugMarshalIndent(databaseBackupUnlockParam))
		return err
	},
}

func databaseBackupUnlockCmdInit() {
	fs := databaseBackupUnlockCmd.Flags()
	fs.IntVarP(&databaseBackupUnlockParam.Index, "index", "", 0, "index of target backup")
}

var databaseBackupRemoveCmd = &cobra.Command{
	Use: "backup-remove",

	Short: "Remove backup",
	Long:  `Remove backup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseBackupRemoveParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("backup-remove parameter: \n%s\n", debugMarshalIndent(databaseBackupRemoveParam))
		return err
	},
}

func databaseBackupRemoveCmdInit() {
	fs := databaseBackupRemoveCmd.Flags()
	fs.IntVarP(&databaseBackupRemoveParam.Index, "index", "", 0, "index of target backup")
}

var databaseCloneCmd = &cobra.Command{
	Use: "clone",

	Short: "Create clone instance",
	Long:  `Create clone instance`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseCloneParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("clone parameter: \n%s\n", debugMarshalIndent(databaseCloneParam))
		return err
	},
}

func databaseCloneCmdInit() {
	fs := databaseCloneCmd.Flags()
	fs.StringSliceVarP(&databaseCloneParam.SourceNetworks, "source-networks", "", []string{}, "set network of allow connection")
	fs.BoolVarP(&databaseCloneParam.EnableBackup, "enable-backup", "", false, "enable backup")
	fs.StringVarP(&databaseCloneParam.BackupTime, "backup-time", "", "", "set backup start time")
	fs.IntVarP(&databaseCloneParam.NwMaskLen, "nw-mask-len", "", 0, "set network mask length")
	fs.IntVarP(&databaseCloneParam.Plan, "plan", "", 10, "set plan[10/30/90/240/500/1000]")
	fs.StringVarP(&databaseCloneParam.DefaultRoute, "default-route", "", "", "set default route")
	fs.StringVarP(&databaseCloneParam.Name, "name", "", "", "set resource display name")
	fs.VarP(newIDValue(0, &databaseCloneParam.SwitchId), "switch-id", "", "set connect switch ID")
	fs.StringVarP(&databaseCloneParam.ReplicaUserPassword, "replica-user-password", "", "", "set database replica user password")
	fs.StringSliceVarP(&databaseCloneParam.BackupWeekdays, "backup-weekdays", "", []string{"all"}, "set backup target weekdays[all or mon/tue/wed/thu/fri/sat/sun]")
	fs.StringVarP(&databaseCloneParam.Ipaddress1, "ipaddress-1", "", "", "set ipaddress(#1)")
	fs.StringVarP(&databaseCloneParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&databaseCloneParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.BoolVarP(&databaseCloneParam.EnableWebUi, "enable-web-ui", "", false, "enable web-ui")
	fs.IntVarP(&databaseCloneParam.Port, "port", "", 0, "set database port")
	fs.VarP(newIDValue(0, &databaseCloneParam.IconId), "icon-id", "", "set Icon ID")
}

var databaseReplicaCreateCmd = &cobra.Command{
	Use: "replica-create",

	Short: "Create replication slave instance",
	Long:  `Create replication slave instance`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseReplicaCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("replica-create parameter: \n%s\n", debugMarshalIndent(databaseReplicaCreateParam))
		return err
	},
}

func databaseReplicaCreateCmdInit() {
	fs := databaseReplicaCreateCmd.Flags()
	fs.StringVarP(&databaseReplicaCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&databaseReplicaCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &databaseReplicaCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.VarP(newIDValue(0, &databaseReplicaCreateParam.SwitchId), "switch-id", "", "set connect switch ID")
	fs.StringVarP(&databaseReplicaCreateParam.Ipaddress1, "ipaddress-1", "", "", "set ipaddress(#1)")
	fs.IntVarP(&databaseReplicaCreateParam.NwMaskLen, "nw-mask-len", "", 0, "set network mask length")
	fs.StringVarP(&databaseReplicaCreateParam.DefaultRoute, "default-route", "", "", "set default route")
	fs.StringVarP(&databaseReplicaCreateParam.Name, "name", "", "", "set resource display name")
}

var databaseMonitorCPUCmd = &cobra.Command{
	Use: "monitor-cpu",

	Short: "Collect CPU monitor values",
	Long:  `Collect CPU monitor values`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseMonitorCPUParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-cpu parameter: \n%s\n", debugMarshalIndent(databaseMonitorCPUParam))
		return err
	},
}

func databaseMonitorCPUCmdInit() {
	fs := databaseMonitorCPUCmd.Flags()
	fs.StringVarP(&databaseMonitorCPUParam.Start, "start", "", "", "set start-time")
	fs.StringVarP(&databaseMonitorCPUParam.End, "end", "", "", "set end-time")
	fs.StringVarP(&databaseMonitorCPUParam.KeyFormat, "key-format", "", "sakuracloud.database.{{.ID}}.cpu", "set monitoring value key-format")
}

var databaseMonitorMemoryCmd = &cobra.Command{
	Use: "monitor-memory",

	Short: "Collect memory monitor values",
	Long:  `Collect memory monitor values`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseMonitorMemoryParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-memory parameter: \n%s\n", debugMarshalIndent(databaseMonitorMemoryParam))
		return err
	},
}

func databaseMonitorMemoryCmdInit() {
	fs := databaseMonitorMemoryCmd.Flags()
	fs.StringVarP(&databaseMonitorMemoryParam.Start, "start", "", "", "set start-time")
	fs.StringVarP(&databaseMonitorMemoryParam.End, "end", "", "", "set end-time")
	fs.StringVarP(&databaseMonitorMemoryParam.KeyFormat, "key-format", "", "sakuracloud.database.{{.ID}}.memory", "set monitoring value key-format")
}

var databaseMonitorNicCmd = &cobra.Command{
	Use: "monitor-nic",

	Short: "Collect NIC(s) monitor values",
	Long:  `Collect NIC(s) monitor values`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseMonitorNicParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-nic parameter: \n%s\n", debugMarshalIndent(databaseMonitorNicParam))
		return err
	},
}

func databaseMonitorNicCmdInit() {
	fs := databaseMonitorNicCmd.Flags()
	fs.StringVarP(&databaseMonitorNicParam.Start, "start", "", "", "set start-time")
	fs.StringVarP(&databaseMonitorNicParam.End, "end", "", "", "set end-time")
	fs.StringVarP(&databaseMonitorNicParam.KeyFormat, "key-format", "", "sakuracloud.database.{{.ID}}.nic", "set monitoring value key-format")
}

var databaseMonitorSystemDiskCmd = &cobra.Command{
	Use: "monitor-system-disk",

	Short: "Collect system-disk monitor values(IO)",
	Long:  `Collect system-disk monitor values(IO)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseMonitorSystemDiskParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-system-disk parameter: \n%s\n", debugMarshalIndent(databaseMonitorSystemDiskParam))
		return err
	},
}

func databaseMonitorSystemDiskCmdInit() {
	fs := databaseMonitorSystemDiskCmd.Flags()
	fs.StringVarP(&databaseMonitorSystemDiskParam.Start, "start", "", "", "set start-time")
	fs.StringVarP(&databaseMonitorSystemDiskParam.End, "end", "", "", "set end-time")
	fs.StringVarP(&databaseMonitorSystemDiskParam.KeyFormat, "key-format", "", "sakuracloud.database.{{.ID}}.disk1", "set monitoring value key-format")
}

var databaseMonitorBackupDiskCmd = &cobra.Command{
	Use: "monitor-backup-disk",

	Short: "Collect backup-disk monitor values(IO)",
	Long:  `Collect backup-disk monitor values(IO)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseMonitorBackupDiskParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-backup-disk parameter: \n%s\n", debugMarshalIndent(databaseMonitorBackupDiskParam))
		return err
	},
}

func databaseMonitorBackupDiskCmdInit() {
	fs := databaseMonitorBackupDiskCmd.Flags()
	fs.StringVarP(&databaseMonitorBackupDiskParam.Start, "start", "", "", "set start-time")
	fs.StringVarP(&databaseMonitorBackupDiskParam.End, "end", "", "", "set end-time")
	fs.StringVarP(&databaseMonitorBackupDiskParam.KeyFormat, "key-format", "", "sakuracloud.database.{{.ID}}.disk2", "set monitoring value key-format")
}

var databaseMonitorSystemDiskSizeCmd = &cobra.Command{
	Use: "monitor-system-disk-size",

	Short: "Collect system-disk monitor values(usage)",
	Long:  `Collect system-disk monitor values(usage)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseMonitorSystemDiskSizeParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-system-disk-size parameter: \n%s\n", debugMarshalIndent(databaseMonitorSystemDiskSizeParam))
		return err
	},
}

func databaseMonitorSystemDiskSizeCmdInit() {
	fs := databaseMonitorSystemDiskSizeCmd.Flags()
	fs.StringVarP(&databaseMonitorSystemDiskSizeParam.Start, "start", "", "", "set start-time")
	fs.StringVarP(&databaseMonitorSystemDiskSizeParam.End, "end", "", "", "set end-time")
	fs.StringVarP(&databaseMonitorSystemDiskSizeParam.KeyFormat, "key-format", "", "sakuracloud.database.{{.ID}}.disk1", "set monitoring value key-format")
}

var databaseMonitorBackupDiskSizeCmd = &cobra.Command{
	Use: "monitor-backup-disk-size",

	Short: "Collect backup-disk monitor values(usage)",
	Long:  `Collect backup-disk monitor values(usage)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseMonitorBackupDiskSizeParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-backup-disk-size parameter: \n%s\n", debugMarshalIndent(databaseMonitorBackupDiskSizeParam))
		return err
	},
}

func databaseMonitorBackupDiskSizeCmdInit() {
	fs := databaseMonitorBackupDiskSizeCmd.Flags()
	fs.StringVarP(&databaseMonitorBackupDiskSizeParam.Start, "start", "", "", "set start-time")
	fs.StringVarP(&databaseMonitorBackupDiskSizeParam.End, "end", "", "", "set end-time")
	fs.StringVarP(&databaseMonitorBackupDiskSizeParam.KeyFormat, "key-format", "", "sakuracloud.database.{{.ID}}.disk2", "set monitoring value key-format")
}

var databaseLogsCmd = &cobra.Command{
	Use: "logs",

	Short: "Logs Database",
	Long:  `Logs Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := databaseLogsParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("logs parameter: \n%s\n", debugMarshalIndent(databaseLogsParam))
		return err
	},
}

func databaseLogsCmdInit() {
	fs := databaseLogsCmd.Flags()
	fs.StringVarP(&databaseLogsParam.LogName, "log-name", "", "all", "set target logfile name")
	fs.BoolVarP(&databaseLogsParam.Follow, "follow", "f", false, "follow log output")
	fs.VarP(newIDValue(0, &databaseLogsParam.RefreshInterval), "refresh-interval", "", "log refresh interval second")
	fs.BoolVarP(&databaseLogsParam.ListLogNames, "list-log-names", "", false, "show log-name list")
}

func init() {
	parent := databaseCmd

	databaseListCmdInit()
	parent.AddCommand(databaseListCmd)

	databaseCreateCmdInit()
	parent.AddCommand(databaseCreateCmd)

	databaseReadCmdInit()
	parent.AddCommand(databaseReadCmd)

	databaseUpdateCmdInit()
	parent.AddCommand(databaseUpdateCmd)

	databaseDeleteCmdInit()
	parent.AddCommand(databaseDeleteCmd)

	databaseBootCmdInit()
	parent.AddCommand(databaseBootCmd)

	databaseShutdownCmdInit()
	parent.AddCommand(databaseShutdownCmd)

	databaseShutdownForceCmdInit()
	parent.AddCommand(databaseShutdownForceCmd)

	databaseResetCmdInit()
	parent.AddCommand(databaseResetCmd)

	databaseWaitForBootCmdInit()
	parent.AddCommand(databaseWaitForBootCmd)

	databaseWaitForDownCmdInit()
	parent.AddCommand(databaseWaitForDownCmd)

	databaseBackupInfoCmdInit()
	parent.AddCommand(databaseBackupInfoCmd)

	databaseBackupCreateCmdInit()
	parent.AddCommand(databaseBackupCreateCmd)

	databaseBackupRestoreCmdInit()
	parent.AddCommand(databaseBackupRestoreCmd)

	databaseBackupLockCmdInit()
	parent.AddCommand(databaseBackupLockCmd)

	databaseBackupUnlockCmdInit()
	parent.AddCommand(databaseBackupUnlockCmd)

	databaseBackupRemoveCmdInit()
	parent.AddCommand(databaseBackupRemoveCmd)

	databaseCloneCmdInit()
	parent.AddCommand(databaseCloneCmd)

	databaseReplicaCreateCmdInit()
	parent.AddCommand(databaseReplicaCreateCmd)

	databaseMonitorCPUCmdInit()
	parent.AddCommand(databaseMonitorCPUCmd)

	databaseMonitorMemoryCmdInit()
	parent.AddCommand(databaseMonitorMemoryCmd)

	databaseMonitorNicCmdInit()
	parent.AddCommand(databaseMonitorNicCmd)

	databaseMonitorSystemDiskCmdInit()
	parent.AddCommand(databaseMonitorSystemDiskCmd)

	databaseMonitorBackupDiskCmdInit()
	parent.AddCommand(databaseMonitorBackupDiskCmd)

	databaseMonitorSystemDiskSizeCmdInit()
	parent.AddCommand(databaseMonitorSystemDiskSizeCmd)

	databaseMonitorBackupDiskSizeCmdInit()
	parent.AddCommand(databaseMonitorBackupDiskSizeCmd)

	databaseLogsCmdInit()
	parent.AddCommand(databaseLogsCmd)

	rootCmd.AddCommand(parent)
}
