// Copyright 2017-2019 The Usacloud Authors
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

package command

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/sacloud/usacloud/define"
	"gopkg.in/urfave/cli.v2"
)

type Option struct {
	AccessToken       string
	AccessTokenSecret string
	Zone              string
	ProfileName       string
	Timeout           int
	AcceptLanguage    string
	RetryMax          int
	RetryIntervalSec  int64
	APIRequestTimeout int
	Zones             []string
	APIRootURL        string
	TraceMode         bool
	Format            string
	DefaultOutputType string
	NoColor           bool
	In                *os.File
	Out               io.Writer
	Progress          io.Writer
	Err               io.Writer
	Validated         bool
	Valid             bool
	ValidationResults []error
}

var GlobalOption = &Option{
	In:       os.Stdin,
	Out:      colorable.NewColorableStdout(),
	Progress: colorable.NewColorableStderr(),
	Err:      colorable.NewColorableStderr(),
}

var (
	DefaultZone       = "tk1a"
	DefaultOutputType = "table"
)

func init() {
	if !(isatty.IsTerminal(os.Stderr.Fd()) || isatty.IsCygwinTerminal(os.Stderr.Fd())) {
		GlobalOption.Progress = ioutil.Discard
	}
}

var GlobalFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "token",
		Usage:       "API Token of SakuraCloud",
		EnvVars:     []string{"SAKURACLOUD_ACCESS_TOKEN"},
		DefaultText: "none",
		Destination: &GlobalOption.AccessToken,
	},
	&cli.StringFlag{
		Name:        "secret",
		Usage:       "API Secret of SakuraCloud",
		EnvVars:     []string{"SAKURACLOUD_ACCESS_TOKEN_SECRET"},
		DefaultText: "none",
		Destination: &GlobalOption.AccessTokenSecret,
	},
	&cli.StringFlag{
		Name:        "zone",
		Usage:       "Target zone of SakuraCloud",
		EnvVars:     []string{"SAKURACLOUD_ZONE"},
		Value:       DefaultZone,
		DefaultText: DefaultZone,
		Destination: &GlobalOption.Zone,
	},
	&cli.StringFlag{
		Name:        "config",
		Aliases:     []string{"profile"},
		Usage:       "Config(Profile) name",
		EnvVars:     []string{"USACLOUD_PROFILE"},
		Destination: &GlobalOption.ProfileName,
	},
	&cli.IntFlag{
		Name:        "timeout",
		Usage:       "Number of timeout minutes for polling functions",
		EnvVars:     []string{"SAKURACLOUD_TIMEOUT"},
		Value:       20,
		Destination: &GlobalOption.Timeout,
	},
	&cli.StringFlag{
		Name:        "accept-language",
		Usage:       "Accept-Language Header",
		EnvVars:     []string{"SAKURACLOUD_ACCEPT_LANGUAGE"},
		Destination: &GlobalOption.AcceptLanguage,
	},
	&cli.IntFlag{
		Name:        "retry-max",
		Usage:       "Number of API-Client retries",
		EnvVars:     []string{"SAKURACLOUD_RETRY_MAX"},
		Destination: &GlobalOption.RetryMax,
	},
	&cli.Int64Flag{
		Name:        "retry-interval",
		Usage:       "API client retry interval seconds",
		EnvVars:     []string{"SAKURACLOUD_RETRY_INTERVAL"},
		Destination: &GlobalOption.RetryIntervalSec,
	},
	&cli.IntFlag{
		Name:        "api-request-timeout",
		Usage:       "Maximum wait time(seconds) for calling SakuraCloud API",
		EnvVars:     []string{"SAKURACLOUD_API_REQUEST_TIMEOUT"},
		Destination: &GlobalOption.APIRequestTimeout,
	},
	&cli.BoolFlag{
		Name:        "no-color",
		Usage:       "Flag of not using ANSI color output",
		EnvVars:     []string{"NO_COLOR"},
		Destination: &GlobalOption.NoColor,
	},
	&cli.StringFlag{
		Name:        "api-root-url",
		EnvVars:     []string{"USACLOUD_API_ROOT_URL"},
		Destination: &GlobalOption.APIRootURL,
		Hidden:      true,
	},
	&cli.StringFlag{
		Name:        "default-output-type",
		EnvVars:     []string{"USACLOUD_DEFAULT_OUTPUT_TYPE"},
		Destination: &GlobalOption.DefaultOutputType,
		Value:       DefaultOutputType,
		Hidden:      true,
	},
	&cli.StringSliceFlag{
		Name:   "zones",
		Hidden: true,
	},
	&cli.BoolFlag{
		Name:        "trace",
		Usage:       "Flag of SakuraCloud debug-mode",
		EnvVars:     []string{"SAKURACLOUD_TRACE_MODE"},
		Destination: &GlobalOption.TraceMode,
		Value:       false,
		Hidden:      true,
	},
}

func (o *Option) Validate(skipAuth bool) []error {
	var errs []error

	// token/secret
	needAuth := !skipAuth
	if needAuth {
		errs = append(errs, ValidateRequired("token", o.AccessToken)...)
		errs = append(errs, ValidateRequired("secret", o.AccessTokenSecret)...)
		errs = append(errs, ValidateRequired("zone", o.Zone)...)
		errs = append(errs, ValidateInStrValues("zone", o.Zone, define.AllowZones...)...)
		errs = append(errs, ValidateInStrValues("default-output-type", o.DefaultOutputType, define.AllowOutputTypes...)...)
	}

	o.Validated = true
	o.Valid = len(errs) == 0
	o.ValidationResults = errs

	return errs
}
