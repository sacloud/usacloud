package command

import (
	"github.com/mattn/go-isatty"
	"gopkg.in/urfave/cli.v2"
	"io"
	"io/ioutil"
	"os"
)

type Option struct {
	AccessToken       string
	AccessTokenSecret string
	Zone              string
	TraceMode         bool
	Format            string
	In                io.Reader
	Out               io.Writer
	Progress          io.Writer
	Err               io.Writer
	Validated         bool
	Valid             bool
	ValidationResults []error
}

var GlobalOption = &Option{
	In:       os.Stdin,
	Out:      os.Stdout,
	Progress: os.Stderr,
	Err:      os.Stderr,
}

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
		// TODO 複数ゾーンへの対応
		Name:        "zone",
		Usage:       "Target zone of SakuraCloud",
		EnvVars:     []string{"SAKURACLOUD_ZONE"},
		Value:       "tk1a",
		DefaultText: "tk1a",
		Destination: &GlobalOption.Zone,
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
		errs = append(errs, validateRequired("token", o.AccessToken)...)
		errs = append(errs, validateRequired("secret", o.AccessTokenSecret)...)
		errs = append(errs, validateRequired("zone", o.Zone)...)
	}

	o.Validated = true
	o.Valid = len(errs) == 0
	o.ValidationResults = errs

	return errs
}
