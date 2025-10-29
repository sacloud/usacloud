// Copyright 2017-2025 The sacloud/usacloud Authors
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

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
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
	client := ctx.Saclient()
	op, err := client.ProfileOp()
	if err != nil {
		return nil, err
	}
	names, err := op.List()
	if err != nil {
		return nil, err
	}
	profile, err := client.Profile()
	if err != nil {
		return nil, err
	}
	current := profile.Name

	formatter := func(out io.Writer, profileName string, current bool) {
		fmt.Fprintln(out, profileName)
	}
	if term.IsTerminal() {
		formatter = func(out io.Writer, profileName string, current bool) {
			format := "%s\n"
			if current {
				format = "* " + format
			}
			fmt.Fprintf(out, format, profileName)
		}
	}

	out := ctx.IO().Out()
	for _, name := range names {
		formatter(out, name, name == current)
	}
	return nil, nil
}
