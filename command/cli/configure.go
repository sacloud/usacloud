package cli

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/define"
	"gopkg.in/urfave/cli.v2"
	"io/ioutil"
	"strings"
)

func init() {

	// create config(APIKey) file command
	initParam := &command.ConfigFileValue{}
	command := &cli.Command{
		Name:  "config",
		Usage: "A manage command of APIKey settings",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "token",
				Usage:       "API Token of SakuraCloud",
				Destination: &initParam.AccessToken,
			},
			&cli.StringFlag{
				Name:        "secret",
				Usage:       "API Secret of SakuraCloud",
				Destination: &initParam.AccessTokenSecret,
			},
			&cli.StringFlag{
				Name:        "zone",
				Usage:       "Target zone of SakuraCloud",
				Destination: &initParam.Zone,
			},
			&cli.BoolFlag{
				Name:  "show",
				Usage: "Show current config",
			},
		},
		Action: func(c *cli.Context) error {
			return configAction(c, initParam)
		},
	}
	// build Category-Resource mapping
	AppendResourceCategoryMap("config", &define.CategoryConfig)
	Commands = append(Commands, command)
}

func configAction(c *cli.Context, inputParams *command.ConfigFileValue) error {

	needAsk := inputParams.IsEmpty()
	out := command.GlobalOption.Out

	// load current config file
	conf, err := command.LoadConfigFile()
	if err != nil {
		return fmt.Errorf("Config: Loading configFile is failed: %s", err)
	}

	if c.Bool("show") {
		fmt.Fprintf(out, "\n")
		fmt.Fprintf(out, "token=%s\n", conf.AccessToken)
		fmt.Fprintf(out, "secret=%s\n", conf.AccessTokenSecret)
		fmt.Fprintf(out, "zone=%s\n", conf.Zone)
		fmt.Fprintf(out, "\n")
		return nil
	}

	// config value errors
	errs := []error{}

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
				fmt.Fprintf(out, "\n\t%s[%s](default:tk1a): ", "Enter zone", strings.Join(command.AllowZones, "/"))
				fmt.Fscanln(command.GlobalOption.In, &input)

				if errs := validateInStrValues("", input, command.AllowZones...); len(errs) == 0 {
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

	// validate zone
	errs = append(errs, validateInStrValues("--zone", inputParams.Zone, command.AllowZones...)...)
	if len(errs) > 0 {
		return command.FlattenErrors(errs)
	}

	// write file
	filePath, err := command.GetConfigFilePath()
	if err != nil {
		return fmt.Errorf("Config: Getting configFilePath is failed: %s", err)
	}
	rawBody, err := json.MarshalIndent(inputParams, "", "\t")
	if err != nil {
		return fmt.Errorf("Config: Creating configFile body is failed: %s", err)
	}

	err = ioutil.WriteFile(filePath, rawBody, 0600)
	if err != nil {
		return fmt.Errorf("Config: Writing configFile is failed: %s", err)
	}
	color.New(color.FgHiGreen).Fprintf(out, "\nWritten your settings to %s\n", filePath)
	return nil
}
