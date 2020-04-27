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
	"fmt"

	"github.com/sacloud/usacloud/pkg/config"

	"github.com/sacloud/libsacloud/v2/sacloud/profile"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func ConfigShow(ctx cli.Context, params *params.ShowConfigParam) error {

	profileName := ""
	if ctx.NArgs() == 0 {
		n, err := profile.CurrentName()
		if err != nil {
			return err
		}
		profileName = n
	} else {
		profileName = ctx.Args()[0]
	}

	conf := &config.Config{}
	if err := profile.Load(profileName, conf); err != nil {
		return err
	}

	out := ctx.IO().Out()
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "token               = %s\n", conf.AccessToken)
	fmt.Fprintf(out, "secret              = %s\n", conf.AccessTokenSecret)
	fmt.Fprintf(out, "zone                = %s\n", conf.Zone)
	fmt.Fprintf(out, "default-output-type = %s\n", conf.DefaultOutputType)
	fmt.Fprintf(out, "\n")
	return nil
}
