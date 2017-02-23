package command

import (
	"fmt"
	"gopkg.in/urfave/cli.v2"
	"io"
	"os"
)

type Option struct {
	AccessToken       string
	AccessTokenSecret string
	Zone              string
	TraceMode         bool
	Format            string
	Out               io.Writer
	Err               io.Writer
}

var GlobalOption = &Option{
	Out: os.Stdout,
	Err: os.Stderr,
}

var GlobalFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "token",
		Aliases:     []string{"sakuracloud-access-token"},
		Usage:       "API Token of SakuraCloud",
		EnvVars:     []string{"SAKURACLOUD_ACCESS_TOKEN"},
		DefaultText: "none",
		Destination: &GlobalOption.AccessToken,
	},
	&cli.StringFlag{
		Name:        "secret",
		Aliases:     []string{"sakuracloud-access-token-secret"},
		Usage:       "API Secret of SakuraCloud",
		EnvVars:     []string{"SAKURACLOUD_ACCESS_TOKEN_SECRET"},
		DefaultText: "none",
		Destination: &GlobalOption.AccessTokenSecret,
	},
	&cli.StringFlag{
		// TODO 複数ゾーンへの対応
		Name:        "zone",
		Aliases:     []string{"sakuracloud-default-zone"},
		Usage:       "Target zone of SakuraCloud",
		EnvVars:     []string{"SAKURACLOUD_ZONE"},
		Value:       "tk1a",
		DefaultText: "tk1a",
		Destination: &GlobalOption.Zone,
	},
	&cli.BoolFlag{
		Name:        "trace",
		Aliases:     []string{"sakuracloud-trace-mode"},
		Usage:       "Flag of SakuraCloud debug-mode",
		EnvVars:     []string{"SAKURACLOUD_TRACE_MODE"},
		Destination: &GlobalOption.TraceMode,
		Value:       false,
	},
	&cli.StringFlag{ // TODO 移動、schema.commandで指定するように
		Name:        "format",
		Aliases:     []string{"output-format"},
		Usage:       "Output format",
		Value:       "table",
		DefaultText: "table",
		Destination: &GlobalOption.Format,
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

	// format
	formatAllows := []string{"table", "json", "tsv", "csv"}
	isValidFormat := false
	for _, f := range formatAllows {
		if f == o.Format {
			isValidFormat = true
		}
	}
	if !isValidFormat {
		errs = append(errs, fmt.Errorf("%q: is invalid. Either [table/json/tsv/csv] is required", "format"))
	}

	return errs
}
