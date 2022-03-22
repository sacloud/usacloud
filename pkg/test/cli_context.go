// Copyright 2017-2022 The Usacloud Authors
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

package test

import (
	"context"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/config"
	"github.com/sacloud/usacloud/pkg/output"
)

type DummyCLIContextValue struct {
	Context      context.Context
	IO           cli.IO
	Option       *config.Config
	Output       output.Output
	ResourceName string
	CommandName  string
	ID           types.ID
	Zone         string
	Resource     interface{}
	Args         []string
}

type DummyCLIContext struct {
	DummyValue *DummyCLIContextValue
}

func (c *DummyCLIContext) IO() cli.IO {
	return c.DummyValue.IO
}

func (c *DummyCLIContext) Option() *config.Config {
	return c.DummyValue.Option
}

func (c *DummyCLIContext) Output() output.Output {
	return c.DummyValue.Output
}

func (c *DummyCLIContext) ResourceName() string {
	return c.DummyValue.ResourceName
}

func (c *DummyCLIContext) CommandName() string {
	return c.DummyValue.CommandName
}

func (c *DummyCLIContext) ID() types.ID {
	return c.DummyValue.ID
}

func (c *DummyCLIContext) Zone() string {
	return c.DummyValue.Zone
}

func (c *DummyCLIContext) Resource() interface{} {
	return c.DummyValue.Resource
}

func (c *DummyCLIContext) WithResource(id types.ID, zone string, resource interface{}) cli.Context {
	return &DummyCLIContext{
		DummyValue: &DummyCLIContextValue{
			Context:      c.DummyValue.Context,
			IO:           c.DummyValue.IO,
			Option:       c.DummyValue.Option,
			Output:       c.DummyValue.Output,
			ResourceName: c.DummyValue.ResourceName,
			CommandName:  c.DummyValue.CommandName,
			Args:         c.DummyValue.Args,
			ID:           id,
			Zone:         zone,
			Resource:     resource,
		},
	}
}

func (c *DummyCLIContext) Client() iaas.APICaller {
	return APICaller()
}

func (c *DummyCLIContext) Deadline() (deadline time.Time, ok bool) {
	return c.DummyValue.Context.Deadline()
}

func (c *DummyCLIContext) Done() <-chan struct{} {
	return c.DummyValue.Context.Done()
}

func (c *DummyCLIContext) Err() error {
	return c.DummyValue.Context.Err()
}

func (c *DummyCLIContext) Value(key interface{}) interface{} {
	return c.DummyValue.Context.Value(key)
}

func (c *DummyCLIContext) Args() []string {
	return c.DummyValue.Args
}
