// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func PrivateHostListCompleteFlags(ctx command.Context, params *params.ListPrivateHostParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["PrivateHost"].Commands["list"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["PrivateHost"].Commands["list"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags", "selector":
		param := define.Resources["PrivateHost"].Commands["list"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "from", "offset":
		param := define.Resources["PrivateHost"].Commands["list"].BuildedParams().Get("from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "max", "limit":
		param := define.Resources["PrivateHost"].Commands["list"].BuildedParams().Get("max")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sort":
		param := define.Resources["PrivateHost"].Commands["list"].BuildedParams().Get("sort")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PrivateHostCreateCompleteFlags(ctx command.Context, params *params.CreatePrivateHostParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["PrivateHost"].Commands["create"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["PrivateHost"].Commands["create"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["PrivateHost"].Commands["create"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "icon-id":
		param := define.Resources["PrivateHost"].Commands["create"].BuildedParams().Get("icon-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PrivateHostReadCompleteFlags(ctx command.Context, params *params.ReadPrivateHostParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["PrivateHost"].Commands["read"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["PrivateHost"].Commands["read"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PrivateHostUpdateCompleteFlags(ctx command.Context, params *params.UpdatePrivateHostParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["PrivateHost"].Commands["update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "name":
		param := define.Resources["PrivateHost"].Commands["update"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["PrivateHost"].Commands["update"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["PrivateHost"].Commands["update"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "icon-id":
		param := define.Resources["PrivateHost"].Commands["update"].BuildedParams().Get("icon-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["PrivateHost"].Commands["update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PrivateHostDeleteCompleteFlags(ctx command.Context, params *params.DeletePrivateHostParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["PrivateHost"].Commands["delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["PrivateHost"].Commands["delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PrivateHostServerInfoCompleteFlags(ctx command.Context, params *params.ServerInfoPrivateHostParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["PrivateHost"].Commands["server-info"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["PrivateHost"].Commands["server-info"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PrivateHostServerAddCompleteFlags(ctx command.Context, params *params.ServerAddPrivateHostParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "server-id":
		param := define.Resources["PrivateHost"].Commands["server-add"].BuildedParams().Get("server-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["PrivateHost"].Commands["server-add"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["PrivateHost"].Commands["server-add"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func PrivateHostServerDeleteCompleteFlags(ctx command.Context, params *params.ServerDeletePrivateHostParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "server-id":
		param := define.Resources["PrivateHost"].Commands["server-delete"].BuildedParams().Get("server-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["PrivateHost"].Commands["server-delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["PrivateHost"].Commands["server-delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
