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

package main

import (
	"bytes"
	"text/template"

	"github.com/sacloud/usacloud/tools"
)

func GenerateManipulateCommand(ctx *tools.GenerateContext, command *tools.Command) (string, error) {
	b := bytes.NewBufferString("")
	t := template.New("c")
	template.Must(t.Parse(manipulateCommandTemplate))

	setParamActions, err := generateSetParamActions(ctx, command)
	if err != nil {
		return "", err
	}

	err = t.Execute(b, map[string]interface{}{
		"ResourceName":    command.ResourceName(),
		"SetParamActions": setParamActions,
		"FunctionName":    command.FunctionName(),
	})
	return b.String(), err
}

var manipulateCommandTemplate = `
	client := sacloud.New{{ .TargetAPIName }}Op(ctx.Client())
	res, err := client.Read(ctx{{ if not .IsGlobal }}, ctx.Zone(){{ end }}, params.Id)
	if err != nil {
		return err
	}

	// TODO Remove and implements here!!
	return fmt.Errorf("Not Implements {{.FunctionName}} : %#v" , p)
`
