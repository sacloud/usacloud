package commands

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/sacloud/libsacloud/v2/sacloud"

	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/pkg/validation"

	"github.com/sacloud/libsacloud/v2/sacloud/profile"
	"github.com/spf13/pflag"
)

var (
	cliOption = &CLIOptions{}
	cliIO     = &IO{
		In:       os.Stdin,
		Out:      os.Stdout,
		Progress: os.Stderr,
		Err:      os.Stderr,
	}
)

// IO TODO 名前
type IO struct {
	In       *os.File
	Out      io.Writer
	Progress io.Writer
	Err      io.Writer
}

// CLIOptions CLIオプション
type CLIOptions struct {
	// Profile プロファイル名
	Profile string

	profile.ConfigValue

	// Foramt コマンド出力時のフォーマット。Goテンプレートを指定
	Format string
	// DefaultOutputType デフォルトアウトプットタイプ
	DefaultOutputType string
	// NoColor ANSIエスケープシーケンスによる色つけを無効化
	NoColor bool
}

func initGlobalFlags(flags *pflag.FlagSet) {
	initCredentialFlags(flags)
	initOutputFlags(flags)
	initDebugFlags(flags)
}

func initCredentialFlags(flags *pflag.FlagSet) {
	fs := pflag.NewFlagSet("Credentials", pflag.ExitOnError)
	fs.StringVarP(&cliOption.Profile, "profile", "", "default", "the name of saved credentials")
	fs.StringVarP(&cliOption.AccessToken, "token", "", "", "the API token used when calling SAKURA Cloud API")
	fs.StringVarP(&cliOption.AccessTokenSecret, "secret", "", "", "the API secret used when calling SAKURA Cloud API")
	fs.StringVarP(&cliOption.Zone, "zone", "", "", "target zone name")
	fs.StringSliceVarP(&cliOption.Zones, "zones", "", []string{}, "permitted zone names")
	flags.AddFlagSet(fs)
}

func initOutputFlags(flags *pflag.FlagSet) {
	fs := pflag.NewFlagSet("Output", pflag.ExitOnError)
	fs.StringVarP(&cliOption.Format, "format", "", "", "the output format with Go template")
	fs.BoolVarP(&cliOption.NoColor, "no-color", "", false, "disable ANSI color output")
	flags.AddFlagSet(fs)
}

func initDebugFlags(flags *pflag.FlagSet) {
	fs := pflag.NewFlagSet("Debug", pflag.ExitOnError)
	fs.StringVarP(&cliOption.TraceMode, "trace", "", "", "enable trace logs for API calling")
	fs.BoolVarP(&cliOption.FakeMode, "fake", "", false, "enable fake API driver")
	fs.StringVarP(&cliOption.FakeStorePath, "fake-store", "", "", "path to file store used by the fake API driver")
	flags.AddFlagSet(fs)
}

func (o *CLIOptions) loadGlobalFlags(flags *pflag.FlagSet) {
	o.loadFromEnv()
	o.loadFromProfile()
	o.loadFromFlags(flags)
	o.fillDefaults()
}

func (o *CLIOptions) fillDefaults() {
	if len(o.Zones) == 0 {
		o.Zones = sacloud.SakuraCloudZones
	}
}

func (o *CLIOptions) loadFromEnv() {
	o.Profile = stringFromEnv("SAKURACLOUD_PROFILE", o.Profile)
	o.AccessToken = stringFromEnv("SAKURACLOUD_ACCESS_TOKEN", o.AccessToken)
	o.AccessTokenSecret = stringFromEnv("SAKURACLOUD_ACCESS_TOKEN_SECRET", o.AccessTokenSecret)
	o.Zone = stringFromEnv("SAKURACLOUD_ZONE", o.Zone)
	o.Zones = stringSliceFromEnv("SAKURACLOUD_ZONES", o.Zones)
	o.AcceptLanguage = stringFromEnv("SAKURACLOUD_ACCEPT_LANGUAGE", o.AcceptLanguage)
	o.RetryMax = intFromEnv("SAKURACLOUD_RETRY_MAX", o.RetryMax)
	o.RetryWaitMax = intFromEnv("SAKURACLOUD_RETRY_WAIT_MAX", o.RetryWaitMax)
	o.RetryWaitMin = intFromEnv("SAKURACLOUD_RETRY_WAIT_MIN", o.RetryWaitMin)
	o.HTTPRequestTimeout = intFromEnv("SAKURACLOUD_API_REQUEST_TIMEOUT", o.HTTPRequestTimeout)
	o.HTTPRequestRateLimit = intFromEnv("SAKURACLOUD_API_REQUEST_RATE_LIMIT", o.HTTPRequestRateLimit)
	o.APIRootURL = stringFromEnv("SAKURACLOUD_API_ROOT_URL", o.APIRootURL)
	o.TraceMode = stringFromEnv("SAKURACLOUD_TRACE", o.TraceMode)
	o.FakeMode = os.Getenv("SAKURACLOUD_FAKE_MODE") != ""
	o.FakeStorePath = stringFromEnv("SAKURACLOUD_FAKE_STORE_PATH", o.FakeStorePath)
}

func (o *CLIOptions) loadFromProfile() {
	if o.Profile != "" {
		if err := profile.Load(o.Profile, o); err != nil {
			fmt.Fprintf(cliIO.Err, "[WARN] loading profile %q is failed: %s", o.Profile, err) // nolint
			return
		}
	}
}

func (o *CLIOptions) loadFromFlags(flags *pflag.FlagSet) {
	out := cliIO.Err
	if flags.Changed("token") {
		v, err := flags.GetString("token")
		if err != nil {
			fmt.Fprintf(out, "[WARN] reading value of %q flag is failed: %s", "token", err) // nolint
			return
		}
		o.AccessToken = v
	}
	if flags.Changed("secret") {
		v, err := flags.GetString("secret")
		if err != nil {
			fmt.Fprintf(out, "[WARN] reading value of %q flag is failed: %s", "secret", err) // nolint
			return
		}
		o.AccessTokenSecret = v
	}
	if flags.Changed("zone") {
		v, err := flags.GetString("zone")
		if err != nil {
			fmt.Fprintf(out, "[WARN] reading value of %q flag is failed: %s", "zone", err) // nolint
			return
		}
		o.Zone = v
	}
	if flags.Changed("zones") {
		v, err := flags.GetStringSlice("zones")
		if err != nil {
			fmt.Fprintf(out, "[WARN] reading value of %q flag is failed: %s", "zones", err) // nolint
			return
		}
		o.Zones = v
	}
	if flags.Changed("format") {
		v, err := flags.GetString("format")
		if err != nil {
			fmt.Fprintf(out, "[WARN] reading value of %q flag is failed: %s", "format", err) // nolint
			return
		}
		o.Format = v
	}
	if flags.Changed("no-color") {
		v, err := flags.GetBool("no-color")
		if err != nil {
			fmt.Fprintf(out, "[WARN] reading value of %q flag is failed: %s", "no-color", err) // nolint
			return
		}
		o.NoColor = v
	}
	if flags.Changed("trace") {
		v, err := flags.GetString("trace")
		if err != nil {
			fmt.Fprintf(out, "[WARN] reading value of %q flag is failed: %s", "trace", err) // nolint
			return
		}
		o.TraceMode = v
	}
	if flags.Changed("fake") {
		v, err := flags.GetBool("fake")
		if err != nil {
			fmt.Fprintf(out, "[WARN] reading value of %q flag is failed: %s", "fake", err) // nolint
			return
		}
		o.FakeMode = v
	}
	if flags.Changed("fake-store") {
		v, err := flags.GetString("fake-store")
		if err != nil {
			fmt.Fprintf(out, "[WARN] reading value of %q flag is failed: %s", "fake-store", err) // nolint
			return
		}
		o.FakeStorePath = v
	}
}

func stringFromEnv(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}

func stringSliceFromEnv(key string, defaultValue []string) []string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	values := strings.Split(v, ",")
	for i := range values {
		values[i] = strings.Trim(values[i], " ")
	}
	return values
}

func intFromEnv(key string, defaultValue int) int {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return defaultValue
	}
	return int(i)
}

func (o *CLIOptions) Validate(skipCred bool) []error {
	var errs []error

	if !skipCred {
		errs = append(errs, validation.Required("token", o.AccessToken)...)
		errs = append(errs, validation.Required("secret", o.AccessTokenSecret)...)
		errs = append(errs, validation.Required("zone", o.Zone)...)
		errs = append(errs, validation.StringInSlice("zone", o.Zone, o.Zones)...)
	}
	errs = append(errs, validation.StringInSlice("default-output-type", o.DefaultOutputType, define.AllowOutputTypes)...)

	return errs
}
