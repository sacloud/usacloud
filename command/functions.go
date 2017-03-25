package command

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/version"
	"gopkg.in/urfave/cli.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func flattenErrors(errors []error) error {
	if len(errors) == 0 {
		return nil
	}
	var list = make([]string, 0)
	for _, str := range errors {
		list = append(list, str.Error())
	}
	return fmt.Errorf(strings.Join(list, "\n"))
}

func flattenErrorsWithPrefix(errors []error, pref string) error {
	var list = make([]string, 0)
	for _, str := range errors {
		list = append(list, fmt.Sprintf("[%s] : %s", pref, str.Error()))
	}
	return fmt.Errorf(strings.Join(list, "\n"))

}

func setSortBy(target sortable, key string) {
	reverse := strings.HasPrefix(key, "-")
	key = strings.Replace(key, "-", "", -1)
	target.SetSortBy(key, reverse)
}

type sortable interface {
	SetSortBy(key string, reverse bool)
}

func createAPIClient() *api.Client {
	c := api.NewClient(GlobalOption.AccessToken, GlobalOption.AccessTokenSecret, GlobalOption.Zone)
	c.UserAgent = fmt.Sprintf("usacloud-%s", version.Version)
	c.TraceMode = GlobalOption.TraceMode
	return c
}

func getOutputWriter(formatter output.OutputFormatter) output.Output {
	o := GlobalOption
	switch formatter.GetOutputType() {
	case "csv":
		return output.NewRowOutput(o.Out, o.Err, ',', formatter)
	case "tsv":
		return output.NewRowOutput(o.Out, o.Err, '\t', formatter)
	case "json":
		return output.NewJSONOutput(o.Out, o.Err)
	default:

		if formatter.GetQuiet() {
			return output.NewIDOutput(o.Out, o.Err)
		} else {
			if formatter.GetFormat() == "" {
				return output.NewTableOutput(o.Out, o.Err, formatter)
			} else {
				return output.NewFreeOutput(o.Out, o.Err, formatter)
			}
		}

	}
}

func StringIDs(ids []int64) []string {
	var strIDs []string

	for _, v := range ids {
		if v != 0 {
			strIDs = append(strIDs, fmt.Sprintf("%d", v))
		}
	}

	return strIDs
}

func GetConfigFilePath() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("getting HomeDir is failed:%s", err)
	}
	return filepath.Join(homeDir, ".usacloud_config"), nil
}

func LoadConfigFile() (*ConfigFileValue, error) {
	v := &ConfigFileValue{}
	filePath, err := GetConfigFilePath()
	if err != nil {
		return v, err
	}

	// file exists?
	if _, err := os.Stat(filePath); err == nil {
		// read file
		buf, err := ioutil.ReadFile(filePath)
		if err != nil {
			return v, err
		}
		if err := json.Unmarshal(buf, v); err != nil {
			return v, err
		}
	}

	return v, nil
}

func getSSHPrivateKeyStorePath(serverID int64) (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("getting HomeDir is failed:%s", err)
	}
	return filepath.Join(homeDir, ".ssh", fmt.Sprintf("sacloud_pkey_%d", serverID)), nil
}

func completionFlagNames(c *cli.Context, commandName string) {
	comm := c.App.Command(commandName)
	if comm == nil {
		return
	}
	for _, f := range comm.VisibleFlags() {
		for _, n := range f.Names() {
			format := "--%s\n"
			if len(n) == 1 {
				format = "-%s\n"
			}
			fmt.Printf(format, n)
		}
	}
}

func getSSHDefaultUserName(client *api.Client, serverID int64) (string, error) {

	// read server
	server, err := client.GetServerAPI().Read(serverID)
	if err != nil {
		return "", err
	}

	if len(server.Disks) == 0 {
		return "", nil
	}

	return getSSHDefaultUserNameDiskRec(client, server.Disks[0].ID)
}

func getSSHDefaultUserNameDiskRec(client *api.Client, diskID int64) (string, error) {

	disk, err := client.GetDiskAPI().Read(diskID)
	if err != nil {
		return "", err
	}

	if disk.SourceDisk != nil {
		return getSSHDefaultUserNameDiskRec(client, disk.SourceDisk.ID)
	}

	if disk.SourceArchive != nil {
		return getSSHDefaultUserNameArchiveRec(client, disk.SourceArchive.ID)

	}

	return "", nil
}

func getSSHDefaultUserNameArchiveRec(client *api.Client, archiveID int64) (string, error) {
	// read archive
	archive, err := client.GetArchiveAPI().Read(archiveID)
	if err != nil {
		return "", err
	}

	if archive.Scope == string(sacloud.ESCopeShared) {

		// has ubuntu/coreos tag?
		if archive.HasTag("distro-ubuntu") {
			return "ubuntu", nil
		}

		if archive.HasTag("distro-vyos") {
			return "vyos", nil
		}

		if archive.HasTag("distro-coreos") {
			return "core", nil
		}
	}
	if archive.SourceDisk != nil {
		return getSSHDefaultUserNameDiskRec(client, archive.SourceDisk.ID)
	}

	if archive.SourceArchive != nil {
		return getSSHDefaultUserNameArchiveRec(client, archive.SourceArchive.ID)
	}
	return "", nil

}

func confirm(msg string) bool {
	fmt.Printf("\n%s(y/n) [n]: ", msg)
	var input string
	fmt.Fscanln(GlobalOption.In, &input)
	return input == "y"
}

func confirmContinue(target string, ids ...int64) bool {
	if len(ids) == 0 {
		return confirm(fmt.Sprintf("Are you sure you want to %s?", target))
	} else {

		strIDs := StringIDs(ids)
		msg := fmt.Sprintf("Target resource IDs => [\n\t%s\n]", strings.Join(strIDs, ",\n\t"))
		return confirm(fmt.Sprintf("%s\nAre you sure you want to %s?", msg, target))
	}
}

func parseDateTimeString(strDateTime string) time.Time {
	allowDatetimeFormatList := []string{
		time.RFC3339,
	}

	if strDateTime != "" {
		for _, format := range allowDatetimeFormatList {
			d, err := time.Parse(format, strDateTime)
			if err == nil {
				// success
				return d
			}
		}
	}

	return time.Now()
}
