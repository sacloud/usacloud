package funcs

import (
	"io"
	"log"

	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/usacloud/output"
)

// dummyCommandContext is a structure for making funcs package testable
type dummyCommandContext struct {
	outputDest io.Writer
	// args is Command line arguments excluding $0
	args  []string
	flags map[string]interface{}
}

func (c *dummyCommandContext) GetOutput() output.Output {
	return c
}

func (c *dummyCommandContext) Print(v ...interface{}) error {
	log.Print(v...)
	return nil
}

func (c *dummyCommandContext) GetAPIClient() *api.Client {
	return dummyContext.GetAPIClient()
}

func (c *dummyCommandContext) Args() []string {
	return c.args
}

func (c *dummyCommandContext) NArgs() int {
	return len(c.args)
}

func (c *dummyCommandContext) IsSet(name string) bool {
	_, ok := c.flags[name]
	return ok
}
