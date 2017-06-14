package params

import "github.com/sacloud/usacloud/command"

func isEmpty(v interface{}) bool {
	return command.IsEmpty(v)
}
