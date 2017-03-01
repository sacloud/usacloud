package command

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/version"
	"gopkg.in/urfave/cli.v2"
	"path/filepath"
	"strings"
)

func flattenErrors(errors []error) error {
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

func getOutputWriter(formater output.OutputFormater) output.Output {
	o := GlobalOption
	switch o.Format { // TODO CSV/TSVサポート
	case "json":
		return output.NewJSONOutput(o.Out, o.Err)
	default:
		return output.NewTableOutput(o.Out, o.Err, formater)
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

func getSSHPrivateKeyStorePath(serverID int64) (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("ServerCreate is failed: getting HomeDir is failed:%s", err)
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
