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

package config

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/sacloud/libsacloud/v2/sacloud/profile"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/config"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/printer"
)

func Edit(ctx cli.Context, params *params.EditConfigParam) error {
	inputParams := &config.Config{
		ConfigValue: profile.ConfigValue{
			AccessToken:       params.Token,
			AccessTokenSecret: params.Secret,
			Zone:              params.Zone,
		},
		DefaultOutputType: params.DefaultOutputType,
	}
	needAsk := inputParams.IsEmpty()
	in := ctx.IO().In()
	out := ctx.IO().Out()
	printer := printer.Printer{NoColor: ctx.Option().NoColor}

	// load current config file
	args := ctx.Args()
	profileName := ""
	if len(args) > 0 {
		profileName = args[0]
	}
	if profileName == "" {
		profileName = profile.DefaultProfileName
	}

	// validate
	err := profile.ValidateName(profileName)
	if err != nil {
		return err
	}

	conf := &config.Config{}
	if err := profile.Load(profileName, conf); err != nil {
		conf = &config.Config{}
	}

	// token
	if needAsk {
		msg := "\nSetting SakuraCloud API Token => "
		fmt.Fprintf(out, "%s", msg)

		exists := conf.AccessToken != ""
		if exists {
			fmt.Fprintf(out, "(Current = ")
			printer.Fprintf(out, color.New(color.FgCyan), "%q", conf.AccessToken)
			fmt.Fprintf(out, ")")
		}

		// if token is exists on config file , confirm whether to change value
		doChange := true
		if exists {
			doChange = cli.ConfirmContinue(in, "change token setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		if doChange {
			// read input
			var input string
			fmt.Fprintf(out, "\t%s: ", "Enter token")
			fmt.Fscanln(in, &input)
			inputParams.AccessToken = input
		} else {
			inputParams.AccessToken = conf.AccessToken
		}
	} else {
		if inputParams.AccessToken == "" {
			inputParams.AccessToken = conf.AccessToken
		}
	}

	// secret
	if needAsk {
		msg := "\nSetting SakuraCloud API Secret => "
		fmt.Fprintf(out, "%s", msg)

		exists := conf.AccessTokenSecret != ""
		if exists {
			fmt.Fprintf(out, "(Current = ")
			printer.Fprintf(out, color.New(color.FgCyan), "%q", conf.AccessTokenSecret)
			fmt.Fprintf(out, ")")
		}

		// if secret is exists on config file , confirm whether to change value
		doChange := true
		if exists {
			doChange = cli.ConfirmContinue(in, "change secret setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		if doChange {
			// read input
			var input string
			fmt.Fprintf(out, "\t%s: ", "Enter secret")
			fmt.Fscanln(in, &input)
			inputParams.AccessTokenSecret = input
		} else {
			inputParams.AccessTokenSecret = conf.AccessTokenSecret
		}
	} else {
		if inputParams.AccessTokenSecret == "" {
			inputParams.AccessTokenSecret = conf.AccessTokenSecret
		}
	}

	// zone
	if needAsk {
		msg := "\nSetting SakuraCloud Zone => "
		fmt.Fprintf(out, "%s", msg)

		exists := conf.Zone != ""
		if exists {
			fmt.Fprintf(out, "(Current = ")
			printer.Fprintf(out, color.New(color.FgCyan), "%q", conf.Zone)
			fmt.Fprintf(out, ")")
		}

		// if secret is exists on config file , confirm whether to change value
		doChange := true
		if exists {
			doChange = cli.ConfirmContinue(in, "change zone setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		if doChange {
			// read input
			var input string
			for {
				fmt.Fprintf(out, "\n\t%s[%s]: ", "Enter zone", strings.Join(config.AllowZones, "/"))
				fmt.Fscanln(in, &input)

				if errs := cli.ValidateInStrValues("", input, config.AllowZones...); len(errs) == 0 {
					break
				}

				fmt.Fprintf(out, "Invalid value. Please input again\n")
			}
			inputParams.Zone = input
		} else {
			inputParams.Zone = conf.Zone
		}
	} else {
		if inputParams.Zone == "" {
			inputParams.Zone = conf.Zone
		}
	}

	// default output type
	if needAsk {
		msg := "\nSetting Default Output Type => "
		fmt.Fprintf(out, "%s", msg)

		exists := conf.DefaultOutputType != ""
		if exists {
			fmt.Fprintf(out, "(Current = ")
			printer.Fprintf(out, color.New(color.FgCyan), "%q", conf.DefaultOutputType)
			fmt.Fprintf(out, ")")
		}

		// if value is exists on config file , confirm whether to change value
		doChange := true
		if exists {
			doChange = cli.ConfirmContinue(in, "change output setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		if doChange {
			// read input
			var input string
			for {
				fmt.Fprintf(out, "\n\t%s[%s]: ", "Enter default-output-type", strings.Join(config.AllowOutputTypes, "/"))
				fmt.Fscanln(in, &input)

				if errs := cli.ValidateInStrValues("", input, config.AllowOutputTypes...); len(errs) == 0 {
					break
				}

				fmt.Fprintf(out, "Invalid value. Please input again\n")
			}
			inputParams.DefaultOutputType = input
		} else {
			inputParams.DefaultOutputType = conf.DefaultOutputType
		}
	} else {
		if inputParams.DefaultOutputType == "" {
			inputParams.DefaultOutputType = conf.DefaultOutputType
		}
	}

	if inputParams.IsEmpty() {
		printer.Fprintf(out, color.New(color.FgCyan), "\nConfig: Values are empty, profile[%q] was not saved\n", profileName)
		return nil
	}

	// write file
	err = profile.Save(profileName, inputParams)
	if err != nil {
		return fmt.Errorf("writing condif file failed: %s", err)
	}

	// get file path
	filePath, err := profile.ConfigFilePath(profileName)
	if err != nil {
		return err
	}

	printer.Fprintf(out, color.New(color.FgHiGreen), "\nWritten your settings to %s\n", filePath)
	return nil
}
