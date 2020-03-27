// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"fmt"

	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/spf13/cobra"
)

// databaseCmd represents the database command
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
		listParam, err := params.NewListDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(listParam))
		return err
	},
}

var databaseCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create Database",
	Long:  `Create Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		createParam, err := params.NewCreateDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("create parameter: \n%s\n", debugMarshalIndent(createParam))
		return err
	},
}

var databaseReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read Database",
	Long:  `Read Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		readParam, err := params.NewReadDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(readParam))
		return err
	},
}

var databaseUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update Database",
	Long:  `Update Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		updateParam, err := params.NewUpdateDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("update parameter: \n%s\n", debugMarshalIndent(updateParam))
		return err
	},
}

var databaseDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete Database",
	Long:    `Delete Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		deleteParam, err := params.NewDeleteDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(deleteParam))
		return err
	},
}

var databaseBootCmd = &cobra.Command{
	Use:     "boot",
	Aliases: []string{"power-on"},
	Short:   "Boot Database",
	Long:    `Boot Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		bootParam, err := params.NewBootDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("boot parameter: \n%s\n", debugMarshalIndent(bootParam))
		return err
	},
}

var databaseShutdownCmd = &cobra.Command{
	Use:     "shutdown",
	Aliases: []string{"power-off"},
	Short:   "Shutdown Database",
	Long:    `Shutdown Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		shutdownParam, err := params.NewShutdownDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("shutdown parameter: \n%s\n", debugMarshalIndent(shutdownParam))
		return err
	},
}

var databaseShutdownForceCmd = &cobra.Command{
	Use:     "shutdown-force",
	Aliases: []string{"stop"},
	Short:   "ShutdownForce Database",
	Long:    `ShutdownForce Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		shutdownForceParam, err := params.NewShutdownForceDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("shutdown-force parameter: \n%s\n", debugMarshalIndent(shutdownForceParam))
		return err
	},
}

var databaseResetCmd = &cobra.Command{
	Use: "reset",

	Short: "Reset Database",
	Long:  `Reset Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resetParam, err := params.NewResetDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("reset parameter: \n%s\n", debugMarshalIndent(resetParam))
		return err
	},
}

var databaseWaitForBootCmd = &cobra.Command{
	Use: "wait-for-boot",

	Short: "Wait until boot is completed",
	Long:  `Wait until boot is completed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		waitForBootParam, err := params.NewWaitForBootDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("wait-for-boot parameter: \n%s\n", debugMarshalIndent(waitForBootParam))
		return err
	},
}

var databaseWaitForDownCmd = &cobra.Command{
	Use: "wait-for-down",

	Short: "Wait until shutdown is completed",
	Long:  `Wait until shutdown is completed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		waitForDownParam, err := params.NewWaitForDownDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("wait-for-down parameter: \n%s\n", debugMarshalIndent(waitForDownParam))
		return err
	},
}

var databaseBackupInfoCmd = &cobra.Command{
	Use:     "backup-info",
	Aliases: []string{"backups", "backup-list"},
	Short:   "Show information of backup",
	Long:    `Show information of backup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		backupInfoParam, err := params.NewBackupInfoDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("backup-info parameter: \n%s\n", debugMarshalIndent(backupInfoParam))
		return err
	},
}

var databaseBackupCreateCmd = &cobra.Command{
	Use: "backup-create",

	Short: "Make new database backup",
	Long:  `Make new database backup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		backupCreateParam, err := params.NewBackupCreateDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("backup-create parameter: \n%s\n", debugMarshalIndent(backupCreateParam))
		return err
	},
}

var databaseBackupRestoreCmd = &cobra.Command{
	Use: "backup-restore",

	Short: "Restore database from backup",
	Long:  `Restore database from backup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		backupRestoreParam, err := params.NewBackupRestoreDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("backup-restore parameter: \n%s\n", debugMarshalIndent(backupRestoreParam))
		return err
	},
}

var databaseBackupLockCmd = &cobra.Command{
	Use: "backup-lock",

	Short: "Lock backup",
	Long:  `Lock backup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		backupLockParam, err := params.NewBackupLockDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("backup-lock parameter: \n%s\n", debugMarshalIndent(backupLockParam))
		return err
	},
}

var databaseBackupUnlockCmd = &cobra.Command{
	Use: "backup-unlock",

	Short: "Unlock backup",
	Long:  `Unlock backup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		backupUnlockParam, err := params.NewBackupUnlockDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("backup-unlock parameter: \n%s\n", debugMarshalIndent(backupUnlockParam))
		return err
	},
}

var databaseBackupRemoveCmd = &cobra.Command{
	Use: "backup-remove",

	Short: "Remove backup",
	Long:  `Remove backup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		backupRemoveParam, err := params.NewBackupRemoveDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("backup-remove parameter: \n%s\n", debugMarshalIndent(backupRemoveParam))
		return err
	},
}

var databaseCloneCmd = &cobra.Command{
	Use: "clone",

	Short: "Create clone instance",
	Long:  `Create clone instance`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cloneParam, err := params.NewCloneDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("clone parameter: \n%s\n", debugMarshalIndent(cloneParam))
		return err
	},
}

var databaseReplicaCreateCmd = &cobra.Command{
	Use: "replica-create",

	Short: "Create replication slave instance",
	Long:  `Create replication slave instance`,
	RunE: func(cmd *cobra.Command, args []string) error {
		replicaCreateParam, err := params.NewReplicaCreateDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("replica-create parameter: \n%s\n", debugMarshalIndent(replicaCreateParam))
		return err
	},
}

var databaseMonitorCpuCmd = &cobra.Command{
	Use: "monitor-cpu",

	Short: "Collect CPU monitor values",
	Long:  `Collect CPU monitor values`,
	RunE: func(cmd *cobra.Command, args []string) error {
		monitorCpuParam, err := params.NewMonitorCpuDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-cpu parameter: \n%s\n", debugMarshalIndent(monitorCpuParam))
		return err
	},
}

var databaseMonitorMemoryCmd = &cobra.Command{
	Use: "monitor-memory",

	Short: "Collect memory monitor values",
	Long:  `Collect memory monitor values`,
	RunE: func(cmd *cobra.Command, args []string) error {
		monitorMemoryParam, err := params.NewMonitorMemoryDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-memory parameter: \n%s\n", debugMarshalIndent(monitorMemoryParam))
		return err
	},
}

var databaseMonitorNicCmd = &cobra.Command{
	Use: "monitor-nic",

	Short: "Collect NIC(s) monitor values",
	Long:  `Collect NIC(s) monitor values`,
	RunE: func(cmd *cobra.Command, args []string) error {
		monitorNicParam, err := params.NewMonitorNicDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-nic parameter: \n%s\n", debugMarshalIndent(monitorNicParam))
		return err
	},
}

var databaseMonitorSystemDiskCmd = &cobra.Command{
	Use: "monitor-system-disk",

	Short: "Collect system-disk monitor values(IO)",
	Long:  `Collect system-disk monitor values(IO)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		monitorSystemDiskParam, err := params.NewMonitorSystemDiskDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-system-disk parameter: \n%s\n", debugMarshalIndent(monitorSystemDiskParam))
		return err
	},
}

var databaseMonitorBackupDiskCmd = &cobra.Command{
	Use: "monitor-backup-disk",

	Short: "Collect backup-disk monitor values(IO)",
	Long:  `Collect backup-disk monitor values(IO)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		monitorBackupDiskParam, err := params.NewMonitorBackupDiskDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-backup-disk parameter: \n%s\n", debugMarshalIndent(monitorBackupDiskParam))
		return err
	},
}

var databaseMonitorSystemDiskSizeCmd = &cobra.Command{
	Use: "monitor-system-disk-size",

	Short: "Collect system-disk monitor values(usage)",
	Long:  `Collect system-disk monitor values(usage)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		monitorSystemDiskSizeParam, err := params.NewMonitorSystemDiskSizeDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-system-disk-size parameter: \n%s\n", debugMarshalIndent(monitorSystemDiskSizeParam))
		return err
	},
}

var databaseMonitorBackupDiskSizeCmd = &cobra.Command{
	Use: "monitor-backup-disk-size",

	Short: "Collect backup-disk monitor values(usage)",
	Long:  `Collect backup-disk monitor values(usage)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		monitorBackupDiskSizeParam, err := params.NewMonitorBackupDiskSizeDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor-backup-disk-size parameter: \n%s\n", debugMarshalIndent(monitorBackupDiskSizeParam))
		return err
	},
}

var databaseLogsCmd = &cobra.Command{
	Use: "logs",

	Short: "Logs Database",
	Long:  `Logs Database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logsParam, err := params.NewLogsDatabaseParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("logs parameter: \n%s\n", debugMarshalIndent(logsParam))
		return err
	},
}

func init() {
	parent := databaseCmd
	parent.AddCommand(databaseListCmd)
	parent.AddCommand(databaseCreateCmd)
	parent.AddCommand(databaseReadCmd)
	parent.AddCommand(databaseUpdateCmd)
	parent.AddCommand(databaseDeleteCmd)
	parent.AddCommand(databaseBootCmd)
	parent.AddCommand(databaseShutdownCmd)
	parent.AddCommand(databaseShutdownForceCmd)
	parent.AddCommand(databaseResetCmd)
	parent.AddCommand(databaseWaitForBootCmd)
	parent.AddCommand(databaseWaitForDownCmd)
	parent.AddCommand(databaseBackupInfoCmd)
	parent.AddCommand(databaseBackupCreateCmd)
	parent.AddCommand(databaseBackupRestoreCmd)
	parent.AddCommand(databaseBackupLockCmd)
	parent.AddCommand(databaseBackupUnlockCmd)
	parent.AddCommand(databaseBackupRemoveCmd)
	parent.AddCommand(databaseCloneCmd)
	parent.AddCommand(databaseReplicaCreateCmd)
	parent.AddCommand(databaseMonitorCpuCmd)
	parent.AddCommand(databaseMonitorMemoryCmd)
	parent.AddCommand(databaseMonitorNicCmd)
	parent.AddCommand(databaseMonitorSystemDiskCmd)
	parent.AddCommand(databaseMonitorBackupDiskCmd)
	parent.AddCommand(databaseMonitorSystemDiskSizeCmd)
	parent.AddCommand(databaseMonitorBackupDiskSizeCmd)
	parent.AddCommand(databaseLogsCmd)
	rootCmd.AddCommand(databaseCmd)
}
