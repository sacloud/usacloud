package command

import (
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/usacloud/output"
)

type context struct {
	flagContext FlagContext
	client      *api.Client
	output      output.Output
}
type Context interface {
	GetOutput() output.Output
	GetAPIClient() *api.Client
	FlagContext
}

type FlagContext interface {
	IsSet(name string) bool
}

func NewContext(flagContext FlagContext, formater output.OutputFormater) Context {

	return &context{
		flagContext: flagContext,
		client:      createAPIClient(),
		output:      getOutputWriter(formater),
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
