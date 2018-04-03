package command

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/version"
)

var numericZeros = []interface{}{
	int(0),
	int8(0),
	int16(0),
	int32(0),
	int64(0),
	uint(0),
	uint8(0),
	uint16(0),
	uint32(0),
	uint64(0),
	float32(0),
	float64(0),
}

// IsEmpty is copied from github.com/stretchr/testify/assert/assetions.go
func IsEmpty(object interface{}) bool {

	if object == nil {
		return true
	} else if object == "" {
		return true
	} else if object == false {
		return true
	}

	for _, v := range numericZeros {
		if object == v {
			return true
		}
	}

	objValue := reflect.ValueOf(object)

	switch objValue.Kind() {
	case reflect.Map:
		fallthrough
	case reflect.Slice, reflect.Chan:
		{
			return (objValue.Len() == 0)
		}
	case reflect.Struct:
		switch object.(type) {
		case time.Time:
			return object.(time.Time).IsZero()
		}
	case reflect.Ptr:
		{
			if objValue.IsNil() {
				return true
			}
			switch object.(type) {
			case *time.Time:
				return object.(*time.Time).IsZero()
			default:
				return false
			}
		}
	}
	return false
}

func FlattenErrors(errors []error) error {
	if len(errors) == 0 {
		return nil
	}
	var list = make([]string, 0)
	for _, str := range errors {
		list = append(list, str.Error())
	}
	return fmt.Errorf(strings.Join(list, "\n"))
}

func FlattenErrorsWithPrefix(errors []error, pref string) error {
	var list = make([]string, 0)
	for _, str := range errors {
		list = append(list, fmt.Sprintf("[%s] : %s", pref, str.Error()))
	}
	return fmt.Errorf(strings.Join(list, "\n"))

}

func createAPIClient() *api.Client {
	c := api.NewClient(GlobalOption.AccessToken, GlobalOption.AccessTokenSecret, GlobalOption.Zone)
	c.UserAgent = fmt.Sprintf("usacloud-%s", version.Version)
	c.TraceMode = GlobalOption.TraceMode

	if GlobalOption.Timeout > 0 {
		c.DefaultTimeoutDuration = time.Duration(GlobalOption.Timeout) * time.Minute
	}

	if GlobalOption.AcceptLanguage != "" {
		c.AcceptLanguage = GlobalOption.AcceptLanguage
	}
	if GlobalOption.RetryMax >= 0 {
		c.RetryMax = GlobalOption.RetryMax
	}
	if GlobalOption.RetryIntervalSec >= 0 {
		c.RetryInterval = time.Duration(GlobalOption.RetryIntervalSec) * time.Second
	}

	return c
}

func getOutputWriter(formatter output.Formatter) output.Output {
	o := GlobalOption
	if formatter.GetQuiet() {
		return output.NewIDOutput(o.Out, o.Err)
	}
	if formatter.GetFormat() != "" || formatter.GetFormatFile() != "" {
		return output.NewFreeOutput(o.Out, o.Err, formatter)
	}
	switch formatter.GetOutputType() {
	case "csv":
		return output.NewRowOutput(o.Out, o.Err, ',', formatter)
	case "tsv":
		return output.NewRowOutput(o.Out, o.Err, '\t', formatter)
	case "json":
		return output.NewJSONOutput(o.Out, o.Err)
	default:
		return output.NewTableOutput(o.Out, o.Err, formatter)
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

func Confirm(msg string) bool {

	fi, err := GlobalOption.In.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Size() > 0 {
		return true
	}
	fmt.Printf("\n%s(y/n) [n]: ", msg)

	var input string
	fmt.Fscanln(GlobalOption.In, &input)
	return input == "y" || input == "yes"
}

func ConfirmContinue(target string, ids ...int64) bool {
	if len(ids) == 0 {
		return Confirm(fmt.Sprintf("Are you sure you want to %s?", target))
	}

	strIDs := StringIDs(ids)
	msg := fmt.Sprintf("Target resource IDs => [\n\t%s\n]", strings.Join(strIDs, ",\n\t"))
	return Confirm(fmt.Sprintf("%s\nAre you sure you want to %s?", msg, target))
}

func UniqIDs(elements []int64) []int64 {
	encountered := map[int64]bool{}
	result := []int64{}
	for v := range elements {
		if !encountered[elements[v]] {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}
