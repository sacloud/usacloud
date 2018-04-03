package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/bitly/go-simplejson"
)

type jsonOutput struct {
	Out io.Writer
	Err io.Writer
}

func NewJSONOutput(out io.Writer, err io.Writer) Output {
	return &jsonOutput{
		Out: out,
		Err: err,
	}
}

func (o *jsonOutput) Print(targets ...interface{}) error {
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

	rawArray, err := json.Marshal(targets)
	if err != nil {
		return fmt.Errorf("JSONOutput:Print: json.Marshal is Failed: %s", err)
	}

	j, err := simplejson.NewJson(rawArray)

	if err != nil {
		return fmt.Errorf("JSONOutput:Print: Create SimpleJSON object is failed: %s", err)
	}

	b, err := j.EncodePretty()
	if err != nil {
		return fmt.Errorf("JSONOutput:Print: Print pretty JSON is failed: %s", err)
	}
	o.Out.Write(b)
	fmt.Fprintln(o.Out, "")
	return nil

}
