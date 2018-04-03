package migration

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/profile"
	"github.com/sacloud/usacloud/helper/printer"
)

const OldConfigFileName = ".usacloud_config"

var (
	confirmMigrateFunc   = confirmMigrate
	confirmOverwriteFunc = confirmOverwrite
)

var (
	migrateInfoMsg = `
================================================================================
Warning: old-style configuration file(%q) found.
         please run 'usacloud config migrate' command.
================================================================================

`
	confirmMigrateMsg = `Are you sure you want to migrate old-style configuration file?`

	confirmOverwriteMsg = `Default profile config[%q] already exists.
Do you want to replace it?`
)

func getOldConfigPath() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, OldConfigFileName), nil
}

func confirmMigrate() bool {
	return command.Confirm(confirmMigrateMsg)
}

func confirmOverwrite(configFilePath string) bool {
	return command.Confirm(fmt.Sprintf(confirmOverwriteMsg, configFilePath))
}

func CheckConfigVersion() error {
	src, err := getOldConfigPath()
	if err != nil {
		return fmt.Errorf("Getting old config path is failed: %s", err)
	}

	// exists?
	if _, err := os.Stat(src); err == nil {
		fmt.Fprintf(command.GlobalOption.Err, migrateInfoMsg, src)
	}

	return nil
}

func MigrateConfig() error {
	src, err := getOldConfigPath()
	if err != nil {
		return fmt.Errorf("Getting old config path is failed: %s", err)
	}

	// exists?
	if _, err := os.Stat(src); err != nil {
		// not exists, noop
		return nil
	}

	if !confirmMigrateFunc() {
		return nil
	}

	// profile[default] exists?
	dest, err := profile.GetConfigFilePath(profile.DefaultProfileName)
	if err != nil {
		return fmt.Errorf("Getting new config path is failed: %s", err)
	}
	if _, err := os.Stat(dest); err == nil {
		// exists, show confirm dialog
		if !confirmOverwriteFunc(dest) {
			return nil
		}
	}

	// prepare dest dir
	destDir := filepath.Dir(dest)
	if _, err := os.Stat(destDir); err != nil {
		err := os.MkdirAll(destDir, 0755)
		if err != nil {
			return fmt.Errorf("Can't create destination dir[%q]: %s", destDir, err)
		}
	}

	// move src to dest
	if err := os.Rename(src, dest); err != nil {
		return fmt.Errorf("Migrating [%q] to [%q] is failed: %s", src, dest, err)
	}

	printer.Fprintf(command.GlobalOption.Out, color.New(color.FgGreen), "\nMigrated: [%q] to [%q]\n", src, dest)
	return nil
}
