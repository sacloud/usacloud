// Copyright 2017-2021 The Usacloud Authors
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

package config

import (
	"fmt"
	"io"

	"github.com/sacloud/libsacloud/v2/sacloud/profile"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/term"
)

var listCommand = &core.Command{
	Name:       "list",
	Aliases:    []string{"ls"},
	Category:   "basic",
	Order:      10,
	NoProgress: true,

	ParameterInitializer: func() interface{} {
		return newListParameter()
	},

	Func: listFunc,
}

type listParameter struct{}

func newListParameter() *listParameter {
	return &listParameter{}
}

func init() {
	Resource.AddCommand(listCommand)
}

func listFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	names, err := profile.List()
	if err != nil {
		return nil, err
	}
	current, err := profile.CurrentName()
	if err != nil {
		return nil, err
	}

	formatter := func(out io.Writer, profileName string, current bool) {
		fmt.Fprintln(out, profileName) // nolint
	}
	if term.IsTerminal() {
		formatter = func(out io.Writer, profileName string, current bool) {
			format := "%s\n"
			if current {
				format = "* " + format
			}
			fmt.Fprintf(out, format, profileName) // nolint
		}
	}

	out := ctx.IO().Out()
	for _, name := range names {
		formatter(out, name, name == current)
	}
	return nil, nil
}
