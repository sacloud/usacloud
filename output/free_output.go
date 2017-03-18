package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/flatmap"
	"github.com/bitly/go-simplejson"
	"io"
	"os"
	"text/template"
)

type freeOutput struct {
	Out    io.Writer
	Err    io.Writer
	Format string
}

func NewFreeOutput(out io.Writer, err io.Writer, option FormatOption) Output {
	return &freeOutput{
		Out:    out,
		Err:    err,
		Format: option.GetFormat(),
	}
}

func (o *freeOutput) Print(targets ...interface{}) error {
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

		flatMap["RowNumber"] = fmt.Sprintf("%d", i+1)

		buf := bytes.NewBufferString("")
		t := template.New("t")
		_, err = t.Parse(o.Format)
		if err != nil {
			return fmt.Errorf("Output format is invalid: %s", err)
		}

		err = t.Execute(buf, flatMap)
		if err != nil {
			return err
		}

		o.Out.Write(buf.Bytes())
		fmt.Fprintln(o.Out, "")

	}

	return nil

}
