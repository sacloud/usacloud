package output

import (
	"fmt"
	"io"
	"os"

	"github.com/ghodss/yaml"
)

type yamlOutput struct {
	out io.Writer
	err io.Writer
}

func NewYAMLOutput(out io.Writer, err io.Writer) Output {
	return &yamlOutput{
		out: out,
		err: err,
	}
}

func (o *yamlOutput) Print(targets ...interface{}) error {
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

	b, err := yaml.Marshal(targets)
	if err != nil {
		return fmt.Errorf("YAMLOutput:Print: yaml.Marshal is Failed: %s", err)
	}
	o.out.Write(b)
	fmt.Fprintln(o.out, "")
	return nil

}
