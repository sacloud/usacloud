// Copyright 2017-2020 The Usacloud Authors
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
