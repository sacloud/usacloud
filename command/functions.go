package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/version"
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
