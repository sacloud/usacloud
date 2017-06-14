package command

import (
	"fmt"
	"github.com/sacloud/usacloud/schema"
	"io/ioutil"
)

type InputOption interface {
	GetParamTemplate() string
	GetParamTemplateFile() string
}

func ValidateInputOption(o InputOption) []error {

	t := o.GetParamTemplate()
	tf := o.GetParamTemplateFile()

	// tmpl and tmpl-file
	if t != "" && tf != "" {
		return []error{fmt.Errorf("%q: can't set with --param-template-file", "--param-template")}
	}

	if tf != "" {
		errs := schema.ValidateFileExists()("--param-template-file", tf)
		if len(errs) > 0 {
			return errs
		}
	}

	return []error{}

}

func GetParamTemplateValue(o InputOption) (string, error) {
	t := o.GetParamTemplate()
	tf := o.GetParamTemplateFile()

	if t == "" && tf == "" {
		return "", nil
	}

	if t != "" {
		return t, nil
	}
	b, err := ioutil.ReadFile(o.GetParamTemplateFile())
	if err != nil {
		return "", fmt.Errorf("Read ParameterTemplateFile[%s] is failed: %s", o.GetParamTemplateFile(), err)
	}
	return string(b), nil
}
