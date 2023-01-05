// Copyright 2017-2023 The sacloud/usacloud Authors
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
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/sacloud/api-client-go/profile"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/config"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/printer"
	"github.com/sacloud/usacloud/pkg/term"
	"github.com/sacloud/usacloud/pkg/validate"
)

var editCommand = &core.Command{
	Name:       "edit",
	Category:   "basic",
	Order:      25,
	NoProgress: true,

	ParameterInitializer: func() interface{} {
		return newEditParameter()
	},

	Func: editProfile,
	ValidateFunc: func(ctx cli.Context, parameter interface{}) error {
		p, ok := parameter.(*EditParameter)
		if !ok {
			return fmt.Errorf("invalid parameter: %v", parameter)
		}
		if len(ctx.Args()) > 0 && p.Name == "" {
			p.Name = ctx.Args()[0]
		}

		if err := validate.Exec(p); err != nil {
			return err
		}

		if !term.IsTerminal() && !p.hasValue() {
			return errors.New("stdin or stdout is not a terminal. please specify values via flags")
		}

		return nil
	},
	CustomCompletionFunc: profileCompletion,
}

type EditParameter struct {
	Name              string `validate:"omitempty,profile_name"`
	AccessToken       string `cli:"access-token,aliases=token"`
	AccessTokenSecret string `cli:"access-token-secret,aliases=secret"`
	Zone              string `validate:"omitempty,zone"`
	DefaultOutputType string `validate:"omitempty,output_type"`
	NoColor           bool
	Use               bool
}

func (p *EditParameter) hasValue() bool {
	return p.AccessToken != "" || p.AccessTokenSecret != "" || p.Zone != "" || p.DefaultOutputType != ""
}

func newEditParameter() *EditParameter {
	return &EditParameter{}
}

func init() {
	Resource.AddCommand(editCommand)
}

func editProfile(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*EditParameter)
	if !ok {
		return nil, fmt.Errorf("invalid parameter: %v", parameter)
	}

	if len(ctx.Args()) > 0 && p.Name == "" {
		p.Name = ctx.Args()[0]
	}

	if p.Name == "" {
		current, err := profile.CurrentName()
		if err != nil {
			return nil, err
		}
		p.Name = current
	}

	newConfigValue := &config.Config{
		ConfigValue: profile.ConfigValue{
			AccessToken:       p.AccessToken,
			AccessTokenSecret: p.AccessTokenSecret,
			Zone:              p.Zone,
		},
		Profile:           p.Name,
		DefaultOutputType: p.DefaultOutputType,
	}

	currentConfig, err := getProfileConfigValue(p.Name)
	if err != nil {
		return nil, err
	}
	if currentConfig == nil {
		currentConfig = &config.Config{}
	}

	out := ctx.IO().Out()
	in := ctx.IO().In()
	msgWriter := &printer.Printer{NoColor: p.NoColor}

	// Note: currentConfigには未知のキーが入る可能性がある&それを保持しておく必要があるため、
	//       ここではnewConfigValueまたは入力値からcurrentConfigへ上書きする形で実装する

	// access token
	if newConfigValue.AccessToken == "" {
		msg := "\nSetting SakuraCloud API Token => "
		fmt.Fprintf(out, "%s", msg)

		doChange := true
		if currentConfig.AccessToken != "" {
			fmt.Fprintf(out, "(Current = ")
			msgWriter.Fprintf(out, color.New(color.FgCyan), "%q", currentConfig.AccessToken)
			fmt.Fprintf(out, ")")
			doChange = cli.ConfirmContinue(in, "change token setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		if doChange {
			// read input
			var input string
			fmt.Fprintf(out, "\t%s: ", "Enter token")
			fmt.Fscanln(in, &input)
			currentConfig.AccessToken = input
		}
	} else {
		currentConfig.AccessToken = newConfigValue.AccessToken
	}

	if newConfigValue.AccessTokenSecret == "" {
		msg := "\nSetting SakuraCloud API Secret=> "
		fmt.Fprintf(out, "%s", msg)

		doChange := true
		if currentConfig.AccessTokenSecret != "" {
			fmt.Fprintf(out, "(Current = ")
			msgWriter.Fprintf(out, color.New(color.FgCyan), "%q", currentConfig.AccessTokenSecret)
			fmt.Fprintf(out, ")")
			doChange = cli.ConfirmContinue(in, "change secret setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		if doChange {
			// read input
			var input string
			fmt.Fprintf(out, "\t%s: ", "Enter secret")
			fmt.Fscanln(in, &input)
			currentConfig.AccessTokenSecret = input
		}
	} else {
		currentConfig.AccessTokenSecret = newConfigValue.AccessTokenSecret
	}

	if newConfigValue.Zone == "" {
		msg := "\nSetting SakuraCloud Zone=> "
		fmt.Fprintf(out, "%s", msg)

		doChange := true
		if currentConfig.Zone != "" {
			fmt.Fprintf(out, "(Current = ")
			msgWriter.Fprintf(out, color.New(color.FgCyan), "%q", currentConfig.Zone)
			fmt.Fprintf(out, ")")
			doChange = cli.ConfirmContinue(in, "change zone setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		if doChange {
			zones := currentConfig.Zones
			if len(zones) == 0 {
				zones = iaas.SakuraCloudZones
			}
			// read input
			var input string
			for {
				fmt.Fprintf(out, "\t%s[%s]: ", "Enter Zone", strings.Join(zones, "/"))
				fmt.Fscanln(in, &input)
				if input == "" || containsString(zones, input) {
					break
				}
				fmt.Fprintf(out, "Invalid value. Please input again\n")
			}
			currentConfig.Zone = input
		}
	} else {
		currentConfig.Zone = newConfigValue.Zone
	}

	if newConfigValue.DefaultOutputType == "" {
		msg := "\nSetting Default Output Type=> "
		fmt.Fprintf(out, "%s", msg)

		doChange := true
		if currentConfig.DefaultOutputType != "" {
			fmt.Fprintf(out, "(Current = ")
			msgWriter.Fprintf(out, color.New(color.FgCyan), "%q", currentConfig.DefaultOutputType)
			fmt.Fprintf(out, ")")
			doChange = cli.ConfirmContinue(in, "change default output type setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		if doChange {
			outputTypes := []string{"table", "json", "yaml"}
			// read input
			var input string
			for {
				fmt.Fprintf(out, "\t%s[%s]: ", "Enter Default Output Type", strings.Join(outputTypes, "/"))
				fmt.Fscanln(in, &input)
				if input == "" || containsString(outputTypes, input) {
					break
				}
				fmt.Fprintf(out, "Invalid value. Please input again\n")
			}
			currentConfig.DefaultOutputType = input
		}
	} else {
		currentConfig.DefaultOutputType = newConfigValue.DefaultOutputType
	}

	if err := profile.Save(p.Name, currentConfig); err != nil {
		return nil, err
	}

	wrote, err := profile.ConfigFilePath(p.Name)
	if err != nil {
		return nil, err
	}
	msgWriter.Fprintf(out, color.New(color.FgHiGreen), "\nWritten your settings to %s\n", wrote)

	current, err := profile.CurrentName()
	if err != nil {
		return nil, err
	}

	// 編集したプロファイルが現在使われていない場合は切り替え
	if current != p.Name {
		if !p.Use {
			if !cli.Confirm(in, fmt.Sprintf("Would you like to switch to profile %q?", p.Name)) {
				return nil, nil
			}
		}
		return nil, profile.SetCurrentName(p.Name)
	}
	return nil, nil
}

func containsString(values []string, v string) bool {
	for _, value := range values {
		if v == value {
			return true
		}
	}
	return false
}
