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

package self

import (
	"fmt"
	"strings"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/self"
)

var idCommand = &core.Command{
	Name:       "id",
	Category:   "basic",
	Order:      10,
	NoProgress: true,

	ParameterInitializer: func() interface{} {
		return newIDParameter()
	},

	Func: idFunc,
}

type idParameter struct {
	NoNewLine bool `cli:",short=n"`
}

func newIDParameter() *idParameter {
	return &idParameter{}
}

func init() {
	Resource.AddCommand(idCommand)
}

func idFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	id, err := self.ID()
	if err != nil {
		return nil, err
	}
	id = strings.TrimRight(id, "\r\n")

	p := parameter.(*idParameter)
	suffix := "\n"
	if p.NoNewLine {
		suffix = ""
	}
	fmt.Fprintf(ctx.IO().Out(), "%s%s", id, suffix) // nolint
	return nil, nil
}
