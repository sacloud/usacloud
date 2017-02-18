package command

import (
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/usacloud/output"
)

type context struct {
	flagContext FlagContext
	client      *api.Client
	output      output.Output
	nargs       int
	args        []string
}
type Context interface {
	GetOutput() output.Output
	GetAPIClient() *api.Client
	Args() []string
	NArgs() int
	FlagContext
}

type FlagContext interface {
	IsSet(name string) bool
}

func NewContext(flagContext FlagContext, args []string, formater output.OutputFormater) Context {

	return &context{
		flagContext: flagContext,
		client:      createAPIClient(),
		output:      getOutputWriter(formater),
		args:        args,
		nargs:       len(args),
	}

}

func (c *context) GetOutput() output.Output {
	return c.output
}

func (c *context) GetAPIClient() *api.Client {
	return c.client
}

func (c *context) IsSet(name string) bool {
	return c.flagContext.IsSet(name)
}

func (c *context) NArgs() int {
	return c.nargs
}

func (c *context) Args() []string {
	return c.args
}
