package funcs

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/command/profile"
	"github.com/sacloud/usacloud/define"
	"strings"
)

func ConfigEdit(ctx command.Context, params *params.EditConfigParam) error {
	inputParams := &profile.ConfigFileValue{
		AccessToken:       params.Token,
		AccessTokenSecret: params.Secret,
		Zone:              params.Zone,
	}
	needAsk := inputParams.IsEmpty()
	out := command.GlobalOption.Out

	// load current config file
	profileName := ""
	if ctx.NArgs() > 0 {
		profileName = ctx.Args()[0]
	}
	if profileName == "" {
		profileName = profile.DefaultProfileName
	}

	// validate
	err := profile.ValidateProfileName(profileName)
	if err != nil {
		return err
	}

	conf, err := profile.LoadConfigFile(profileName)
	if err != nil {
		conf = &profile.ConfigFileValue{}
	}

	// token
	if needAsk {
		msg := "\nSetting SakuraCloud API Token => "
		fmt.Fprintf(out, "%s", msg)

		exists := conf.AccessToken != ""
		if exists {
			fmt.Fprintf(out, "(Current = ")
			color.New(color.FgCyan).Fprintf(out, "%q", conf.AccessToken)
			fmt.Fprintf(out, ")")
		}

		// if token is exists on config file , confirm whether to change value
		doChange := true
		if exists {
			doChange = command.ConfirmContinue("change token setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		if doChange {
			// read input
			var input string
			fmt.Fprintf(out, "\t%s: ", "Enter token")
			fmt.Fscanln(command.GlobalOption.In, &input)
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
			color.New(color.FgCyan).Fprintf(out, "%q", conf.AccessTokenSecret)
			fmt.Fprintf(out, ")")
		}

		// if secret is exists on config file , confirm whether to change value
		doChange := true
		if exists {
			doChange = command.ConfirmContinue("change secret setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		if doChange {
			// read input
			var input string
			fmt.Fprintf(out, "\t%s: ", "Enter secret")
			fmt.Fscanln(command.GlobalOption.In, &input)
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
			color.New(color.FgCyan).Fprintf(out, "%q", conf.Zone)
			fmt.Fprintf(out, ")")
		}

		// if secret is exists on config file , confirm whether to change value
		doChange := true
		if exists {
			doChange = command.ConfirmContinue("change zone setting")
		} else {
			fmt.Fprintf(out, "\n")
		}

		if doChange {
			// read input
			var input string
			for {
				fmt.Fprintf(out, "\n\t%s[%s]: ", "Enter zone", strings.Join(define.AllowZones, "/"))
				fmt.Fscanln(command.GlobalOption.In, &input)

				if errs := validateInStrValues("", input, define.AllowZones...); len(errs) == 0 {
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

	if inputParams.IsEmpty() {
		color.New(color.FgCyan).Fprintf(out, "\nConfig: Values are empty, profile[%q] was not saved\n", profileName)
		return nil
	}

	// write file
	err = profile.SaveConfigFile(profileName, inputParams)
	if err != nil {
		return fmt.Errorf("Config: Writing configFile is failed: %s", err)
	}

	// get file path
	filePath, err := profile.GetConfigFilePath(profileName)
	if err != nil {
		return fmt.Errorf("Config: GetConfigFilePath is failed: %s", err)
	}

	color.New(color.FgHiGreen).Fprintf(out, "\nWritten your settings to %s\n", filePath)
	return nil
}
