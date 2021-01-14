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

package core

import (
	"encoding/json"
	"os"
	"reflect"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/spf13/cobra"

	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/util"
)

func (c *Command) loadParameters(ctx cli.Context, cmd *cobra.Command, parameters cflag.CommonParameterValueHolder) error {
	p := parameters.ParametersFlagValue()
	if p == "" {
		return nil
	}

	// c.currentParameterの実体をParameterInitializerで再度初期化
	// Note: ポインタごと置き換えるとコマンドラインフラグのパースがうまく動かないため実体を差し替える
	reflect.ValueOf(c.currentParameter).Elem().Set(reflect.ValueOf(c.ParameterInitializer()).Elem())

	data, err := util.BytesFromPathOrContent(p)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, c.currentParameter); err != nil {
		return nil
	}

	// os.Argsを元にもう一度フラグをパースする
	// 参照: cobra.Command#ExecuteC
	var flags []string
	if cmd.TraverseChildren {
		_, flags, err = cmd.Traverse(os.Args[1:])
	} else {
		_, flags, err = cmd.Find(os.Args[1:])
	}
	if err != nil {
		return err
	}
	return cmd.ParseFlags(flags)
}
