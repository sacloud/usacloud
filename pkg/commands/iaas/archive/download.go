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

package archive

import (
	"fmt"
	"os"

	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/validate"
)

var downloadCommand = &core.Command{
	Name:         "download",
	Category:     "operation",
	Order:        20,
	SelectorType: core.SelectorTypeRequireSingle, // Note: 現状はダウンロード先を指定してもらう形のため、複数だとファイルを上書きしてしまう

	ValidateFunc: validateDownloadParameter,
	ParameterInitializer: func() interface{} {
		return newDownloadParameter()
	},
}

type downloadParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`

	Destination string `cli:",aliases=dest,category=download,order=10" mapconv:"Writer,omitempty,filters=path_to_writer"` // 省略時は標準出力
	Force       bool   `cli:",short=f,category=download,order=20,desc=overwrite file when --destination file is already exist"`
}

func validateDownloadParameter(ctx cli.Context, parameter interface{}) error {
	if err := validate.Exec(parameter); err != nil {
		return err
	}
	p := parameter.(*downloadParameter)

	var errs []error

	if p.Destination != "" {
		exists := false
		_, err := os.Stat(p.Destination)

		if err == nil {
			exists = true
		} else if !os.IsNotExist(err) {
			errs = append(errs, validate.NewFlagError("--destination", fmt.Sprintf("opening file %q failed", p.Destination)))
		}

		if exists && !p.Force {
			errs = append(errs, validate.NewFlagError("--destination", "file already exists"))
		}
	}

	return validate.NewValidationError(errs...)
}

func newDownloadParameter() *downloadParameter {
	return &downloadParameter{}
}

func init() {
	Resource.AddCommand(downloadCommand)
}
