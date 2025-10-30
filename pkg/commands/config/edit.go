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
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/sacloud/api-client-go/profile"
	"github.com/sacloud/iaas-api-go"
	saht "github.com/sacloud/saclient-go"
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
	Zone              string
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

	reader := func(op saht.ProfileAPI, name string) (*saht.Profile, error) { return op.Read(name) }
	writer := func(op saht.ProfileAPI, p *saht.Profile) (*saht.Profile, error) { return op.Update(p) }
	return __editProfile(ctx, p, reader, writer)
}

func __editProfile(
	ctx cli.Context,
	p *EditParameter,
	reader func(saht.ProfileAPI, string) (*saht.Profile, error),
	writer func(saht.ProfileAPI, *saht.Profile) (*saht.Profile, error),
) ([]interface{}, error) {
	op, err := ctx.Saclient().ProfileOp()
	if err != nil {
		return nil, err
	}

	if len(ctx.Args()) > 0 && p.Name == "" {
		p.Name = ctx.Args()[0]
	}

	if p.Name == "" {
		current, err := op.GetCurrentName()
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

	var currentConfig *config.Config = new(config.Config)
	if loaded, err := reader(op, p.Name); err != nil {
		return nil, err
	} else if err := currentConfig.LoadFromAttributes(loaded); err != nil {
		return nil, err
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
			fmt.Fscanln(in, &input) //nolint:errcheck,gosec
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
			fmt.Fscanln(in, &input) //nolint:errcheck,gosec
			currentConfig.AccessTokenSecret = input
		}
	} else {
		currentConfig.AccessTokenSecret = newConfigValue.AccessTokenSecret
	}

	if newConfigValue.ServicePrincipalID == "" {
		msg := "\nSetting SakuraCloud API Service Principal ID (optional)=> "
		fmt.Fprintf(out, "%s", msg)

		doChange := true
		if currentConfig.ServicePrincipalID != "" {
			fmt.Fprintf(out, "(Current = ")
			msgWriter.Fprintf(out, color.New(color.FgCyan), "%q", currentConfig.ServicePrincipalID)
			fmt.Fprintf(out, ")")
			doChange = cli.ConfirmContinue(in, "change service principal ID setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		if doChange {
			// read input
			var input string
			fmt.Fprintf(out, "\t%s: ", "Enter service principal ID")
			fmt.Fscanln(in, &input) //nolint:errcheck,gosec
			currentConfig.ServicePrincipalID = input
		}
	} else {
		currentConfig.ServicePrincipalID = newConfigValue.ServicePrincipalID
	}

	if newConfigValue.ServicePrincipalKeyID == "" {
		msg := "\nSetting SakuraCloud API Service Principal Key ID (optional)=> "
		fmt.Fprintf(out, "%s", msg)

		doChange := true
		if currentConfig.ServicePrincipalKeyID != "" {
			fmt.Fprintf(out, "(Current = ")
			msgWriter.Fprintf(out, color.New(color.FgCyan), "%q", currentConfig.ServicePrincipalKeyID)
			fmt.Fprintf(out, ")")
			doChange = cli.ConfirmContinue(in, "change service principal key ID setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		if doChange {
			// read input
			var input string
			fmt.Fprintf(out, "\t%s: ", "Enter service principal key ID")
			fmt.Fscanln(in, &input) //nolint:errcheck,gosec
			currentConfig.ServicePrincipalKeyID = input
		}
	} else {
		currentConfig.ServicePrincipalKeyID = newConfigValue.ServicePrincipalKeyID
	}

	if newConfigValue.PrivateKeyPEMPath == "" {
		msg := "\nSetting SakuraCloud API Private Key PEM Path (optional)=> "
		fmt.Fprintf(out, "%s", msg)

		doChange := true
		if currentConfig.PrivateKeyPEMPath != "" {
			fmt.Fprintf(out, "(Current = ")
			msgWriter.Fprintf(out, color.New(color.FgCyan), "%q", currentConfig.PrivateKeyPEMPath)
			fmt.Fprintf(out, ")")
			doChange = cli.ConfirmContinue(in, "change private key PEM path setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		var input string
		if doChange {
			// read input
			fmt.Fprintf(out, "\t%s: ", "Enter private key PEM path")
			fmt.Fscanln(in, &input) //nolint:errcheck,gosec
		}

		if input == "" {
			currentConfig.PrivateKeyPEMPath = input
		} else {
			stat, err := os.Stat(input)
			if err != nil {
				doChange = cli.ConfirmContinue(in, "set nonexistent path")
			} else {
				if doChange {
					if !stat.Mode().IsRegular() {
						doChange = cli.ConfirmContinue(in, "set a non-regular file")
					}
				}
				if doChange {
					if stat.Mode().Perm()&0o077 != 0 {
						doChange = cli.ConfirmContinue(in, "set a file readable by others")
					}
				}
			}
			if doChange {
				currentConfig.PrivateKeyPEMPath = input
			}
		}
	} else {
		currentConfig.PrivateKeyPEMPath = newConfigValue.PrivateKeyPEMPath
	}

	if newConfigValue.TokenEndpoint == "" {
		msg := "\nSetting SakuraCloud API Token Endpoint (optional)=> "
		fmt.Fprintf(out, "%s", msg)

		doChange := true
		if currentConfig.TokenEndpoint != "" {
			fmt.Fprintf(out, "(Current = ")
			msgWriter.Fprintf(out, color.New(color.FgCyan), "%q", currentConfig.TokenEndpoint)
			fmt.Fprintf(out, ")")
			doChange = cli.ConfirmContinue(in, "change token endpoint setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		if doChange {
			// read input
			var input string
			fmt.Fprintf(out, "\t%s: ", "Enter token endpoint")
			fmt.Fscanln(in, &input) //nolint:errcheck,gosec
			currentConfig.TokenEndpoint = input
		}
	} else {
		currentConfig.TokenEndpoint = newConfigValue.TokenEndpoint
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
				fmt.Fscanln(in, &input) //nolint:errcheck,gosec
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
				fmt.Fscanln(in, &input) //nolint:errcheck,gosec
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

	updated, err := currentConfig.IntoAttributes()
	if err != nil {
		return nil, err
	}

	persisted, err := writer(op, updated)
	if err != nil {
		return nil, err
	}
	wrote := persisted.Pathname()
	msgWriter.Fprintf(out, color.New(color.FgHiGreen), "\nWritten your settings to %s\n", wrote)

	current, err := op.GetCurrentName()
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
		return nil, op.SetCurrentName(p.Name)
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
