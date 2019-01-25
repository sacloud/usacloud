package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/sacloud/go-jmespath"
)

type jsonOutput struct {
	out   io.Writer
	err   io.Writer
	query string
}

func NewJSONOutput(out io.Writer, err io.Writer, query string) Output {
	return &jsonOutput{
		out:   out,
		err:   err,
		query: query,
	}
}

func (o *jsonOutput) Print(targets ...interface{}) error {
	if o.out == nil {
		o.out = os.Stdout
	}
	if o.err == nil {
		o.err = os.Stderr
	}

	if len(targets) == 0 {
		fmt.Fprintf(o.err, "Result is empty\n")
		return nil
	}

	var values interface{} = targets

	if o.query != "" {
		v, err := o.searchByJMESPath(targets)
		if err != nil {
			return fmt.Errorf("JSONOutput:Query: jmespath.Search is Failed: %s", err)
		}
		values = v
	}

	rawArray, err := json.Marshal(values)
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
	o.out.Write(b)
	fmt.Fprintln(o.out, "")
	return nil

}

func (o *jsonOutput) searchByJMESPath(v interface{}) (result interface{}, err error) {

	defer func() {
		ret := recover()
		if ret != nil {
			fmt.Fprintf(o.err, "jmespath.Search failed: parse error\n")
			err = fmt.Errorf("jmespath.Search failed: %s", ret)
		}
	}()
	result, err = jmespath.Search(o.query, v)
	return
}
