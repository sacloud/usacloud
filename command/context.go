// Copyright 2017-2019 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

type OutputTypeHolder interface {
	GetOutputType() string
	SetOutputType(string)
}

func NewContext(flagContext FlagContext, args []string, formater interface{}) Context {

	var out output.Output
	if formater != nil {
		if o, ok := formater.(output.Formatter); ok {
			out = getOutputWriter(o)
		}
	}

	return &context{
		flagContext: flagContext,
		client:      createAPIClient(),
		output:      out,
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
