package output

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/flatmap"
	"github.com/bitly/go-simplejson"
	"io"
	"os"
)

type idOutput struct {
	Out io.Writer
	Err io.Writer
}

var idOutputTargetColumns = []string{"ID", "Key", "BillID", "RowNumber"}

func NewIDOutput(out io.Writer, err io.Writer) Output {
	return &idOutput{
		Out: out,
		Err: err,
	}
}

func (o *idOutput) Print(targets ...interface{}) error {
	if o.Out == nil {
		o.Out = os.Stdout
	}
	if o.Err == nil {
		o.Err = os.Stderr
	}

	if len(targets) == 0 {
		fmt.Fprintf(o.Err, "Result is empty\n")
		return nil
	}

	// targets -> byte[] -> []interface{}
	rawArray, err := json.Marshal(targets)
	if err != nil {
		return fmt.Errorf("FreeOutput:Print: json.Marshal is failed: %s", err)
	}

	j, err := simplejson.NewJson(rawArray)
	if err != nil {
		return fmt.Errorf("FreeOutput:Print: create simplejson is failed: %s", err)
	}
	for i := range targets {

		// interface{} -> map[string]interface{}
		v := j.GetIndex(i)
		mapValue, err := v.Map()
		if err != nil {
			return fmt.Errorf("FreeOutput:Print: json format is invalid: %v", err)
		}

		// to flatmap( map[string]string )
		flatMap, err := flatmap.Flatten(mapValue)
		if err != nil {
			return fmt.Errorf("FreeOutput:Print: create flatmap is failed: %v", err)
		}

		value := ""
		for _, key := range idOutputTargetColumns {
			if key == "RowNumber" {
				value = fmt.Sprintf("%d", i+1)
				break
			} else {
				if v, ok := flatMap[key]; ok {
					value = v
					break
				}
			}
		}

		fmt.Fprintf(o.Out, "%s\n", value)
	}

	return nil

}
