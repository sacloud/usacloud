package commands

import (
	goContext "context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	libsacloud "github.com/sacloud/libsacloud/v2"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/fake"
	"github.com/sacloud/libsacloud/v2/sacloud/trace"
	"github.com/sacloud/libsacloud/v2/utils/builder"
	"github.com/sacloud/libsacloud/v2/utils/setup"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/version"
)

type Context interface {
	Option() *CLIOptions
	Output() output.Output
	Client() sacloud.APICaller
	Zone() string
	goContext.Context
}

type context struct {
	parentCtx goContext.Context
	option    *CLIOptions
	output    output.Output

	client     sacloud.APICaller
	clientOnce sync.Once
}

func NewContext(ctx goContext.Context, option *CLIOptions, formatter interface{}) Context {
	return &context{
		parentCtx: ctx,
		option:    option,
		output:    getOutputWriter(formatter),
	}
}

func (c *context) Option() *CLIOptions {
	return c.option
}

func (c *context) Output() output.Output {
	return c.output
}

func (c *context) Client() sacloud.APICaller {
	c.clientOnce.Do(func() {
		o := c.Option()

		httpClient := http.DefaultClient
		httpClient.Timeout = time.Duration(o.HTTPRequestTimeout) * time.Second
		httpClient.Transport = &sacloud.RateLimitRoundTripper{RateLimitPerSec: o.HTTPRequestRateLimit}

		retryWaitMax := sacloud.APIDefaultRetryWaitMax
		retryWaitMin := sacloud.APIDefaultRetryWaitMin
		if o.RetryWaitMax > 0 {
			retryWaitMax = time.Duration(o.RetryWaitMax) * time.Second
		}
		if o.RetryWaitMin > 0 {
			retryWaitMin = time.Duration(o.RetryWaitMin) * time.Second
		}

		ua := fmt.Sprintf("Usacloud/ (+https://github.com/sacloud/usacloud) cli/v%s libsacloud/%s", version.Version, libsacloud.Version)

		caller := &sacloud.Client{
			AccessToken:       o.AccessToken,
			AccessTokenSecret: o.AccessTokenSecret,
			UserAgent:         ua,
			AcceptLanguage:    o.AcceptLanguage,
			RetryMax:          o.RetryMax,
			RetryWaitMax:      retryWaitMax,
			RetryWaitMin:      retryWaitMin,
			HTTPClient:        httpClient,
		}
		sacloud.DefaultStatePollingTimeout = 72 * time.Hour

		if o.TraceMode != "" {
			enableAPITrace := true
			enableHTTPTrace := true

			mode := strings.ToLower(o.TraceMode)
			switch mode {
			case "api":
				enableHTTPTrace = false
			case "http":
				enableAPITrace = false
			}

			if enableAPITrace {
				// note: exact once
				trace.AddClientFactoryHooks()
			}
			if enableHTTPTrace {
				caller.HTTPClient.Transport = &sacloud.TracingRoundTripper{
					Transport: caller.HTTPClient.Transport,
				}
			}
		}

		if o.FakeMode {
			if o.FakeStorePath != "" {
				fake.DataStore = fake.NewJSONFileStore(o.FakeStorePath)
			}
			// note: exact once
			fake.SwitchFactoryFuncToFake()

			defaultInterval := 10 * time.Millisecond

			// update default polling intervals: libsacloud/sacloud
			sacloud.DefaultStatePollingInterval = defaultInterval
			// update default polling intervals: libsacloud/utils/setup
			setup.DefaultDeleteWaitInterval = defaultInterval
			setup.DefaultProvisioningWaitInterval = defaultInterval
			setup.DefaultPollingInterval = defaultInterval
			// update default polling intervals: libsacloud/utils/builder
			builder.DefaultNICUpdateWaitDuration = defaultInterval
		}

		zones := o.Zones
		if len(zones) == 0 {
			zones = sacloud.SakuraCloudZones
		}
		if o.APIRootURL != "" {
			if strings.HasSuffix(o.APIRootURL, "/") {
				o.APIRootURL = strings.TrimRight(o.APIRootURL, "/")
			}
			sacloud.SakuraCloudAPIRoot = o.APIRootURL
		}
		c.client = caller
	})

	return c.client
}

func (c *context) Zone() string {
	return c.Option().Zone
}

func (c *context) Deadline() (deadline time.Time, ok bool) {
	return c.parentCtx.Deadline()
}

func (c *context) Done() <-chan struct{} {
	return c.parentCtx.Done()
}

func (c *context) Err() error {
	return c.parentCtx.Err()
}

func (c *context) Value(key interface{}) interface{} {
	return c.parentCtx.Value(key)
}

func getOutputWriter(rawformatter interface{}) output.Output {
	if rawformatter == nil {
		return nil
	}
	formatter, ok := rawformatter.(output.Formatter)
	if !ok {
		return nil
	}

	o := cliIO
	if formatter.GetQuiet() {
		return output.NewIDOutput(o.Out, o.Err)
	}
	if formatter.GetFormat() != "" || formatter.GetFormatFile() != "" {
		return output.NewFreeOutput(o.Out, o.Err, formatter)
	}
	switch formatter.GetOutputType() {
	case "csv":
		return output.NewRowOutput(o.Out, o.Err, ',', formatter)
	case "tsv":
		return output.NewRowOutput(o.Out, o.Err, '\t', formatter)
	case "json":
		query := formatter.GetQuery()
		if query == "" {
			bQuery, _ := ioutil.ReadFile(formatter.GetQueryFile()) // nolint: err was already checked
			query = string(bQuery)
		}
		return output.NewJSONOutput(o.Out, o.Err, query)
	case "yaml":
		return output.NewYAMLOutput(o.Out, o.Err)
	default:
		return output.NewTableOutput(o.Out, o.Err, formatter)
	}
}
