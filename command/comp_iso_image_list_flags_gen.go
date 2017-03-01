// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func ISOImageListCompleteFlags(ctx Context, params *ListISOImageParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "from":
		comp = define.Resources["ISOImage"].Commands["list"].Params["from"].CompleteFunc
	case "id":
		comp = define.Resources["ISOImage"].Commands["list"].Params["id"].CompleteFunc
	case "max":
		comp = define.Resources["ISOImage"].Commands["list"].Params["max"].CompleteFunc
	case "name":
		comp = define.Resources["ISOImage"].Commands["list"].Params["name"].CompleteFunc
	case "scope":
		comp = define.Resources["ISOImage"].Commands["list"].Params["scope"].CompleteFunc
	case "sort":
		comp = define.Resources["ISOImage"].Commands["list"].Params["sort"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
