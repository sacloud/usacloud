// Copyright 2017-2021 The Usacloud Authors
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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-commands'; DO NOT EDIT

package simplemonitor

import (
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *createParameter) CleanupEmptyValue(fs *pflag.FlagSet) {

}

func (p *createParameter) buildFlags(fs *pflag.FlagSet) {

	fs.StringVarP(&p.Parameters, "parameters", "", p.Parameters, "Input parameters in JSON format")
	fs.BoolVarP(&p.GenerateSkeleton, "generate-skeleton", "", p.GenerateSkeleton, "Output skeleton of parameters with JSON format (aliases: --skeleton)")
	fs.BoolVarP(&p.Example, "example", "", p.Example, "Output example parameters with JSON format")
	fs.BoolVarP(&p.AssumeYes, "assumeyes", "y", p.AssumeYes, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&p.OutputType, "output-type", "o", p.OutputType, "Output format options: [table/json/yaml] (aliases: --out)")
	fs.BoolVarP(&p.Quiet, "quiet", "q", p.Quiet, "Output IDs only")
	fs.StringVarP(&p.Format, "format", "", p.Format, "Output format in Go templates (aliases: --fmt)")
	fs.StringVarP(&p.Query, "query", "", p.Query, "Query for JSON output")
	fs.StringVarP(&p.QueryDriver, "query-driver", "", p.QueryDriver, "Name of the driver that handles queries to JSON output options: [jmespath/jq]")
	fs.StringVarP(&p.Target, "target", "", p.Target, "(*required) ")
	fs.StringVarP(&p.Description, "description", "", p.Description, "")
	fs.StringSliceVarP(&p.Tags, "tags", "", p.Tags, "")
	fs.VarP(core.NewIDFlag(&p.IconID, &p.IconID), "icon-id", "", "")
	fs.IntVarP(&p.DelayLoop, "delay-loop", "", p.DelayLoop, "")
	fs.IntVarP(&p.Timeout, "timeout", "", p.Timeout, "")
	fs.BoolVarP(&p.Enabled, "enabled", "", p.Enabled, "")
	fs.StringVarP(&p.HealthCheck.Protocol, "health-check-protocol", "", p.HealthCheck.Protocol, "(*required) options: [http/https/ping/tcp/dns/ssh/smtp/pop3/snmp/sslcertificate/ftp]")
	fs.IntVarP(&p.HealthCheck.Port, "health-check-port", "", p.HealthCheck.Port, "")
	fs.StringVarP(&p.HealthCheck.Path, "health-check-path", "", p.HealthCheck.Path, "")
	fs.IntVarP(&p.HealthCheck.Status, "health-check-status", "", p.HealthCheck.Status, "")
	fs.StringVarP(&p.HealthCheck.ContainsString, "health-check-contains-string", "", p.HealthCheck.ContainsString, "")
	fs.BoolVarP(&p.HealthCheck.SNI, "health-check-sni", "", p.HealthCheck.SNI, "")
	fs.StringVarP(&p.HealthCheck.Host, "health-check-host", "", p.HealthCheck.Host, "")
	fs.StringVarP(&p.HealthCheck.BasicAuthUsername, "health-check-basic-auth-username", "", p.HealthCheck.BasicAuthUsername, "")
	fs.StringVarP(&p.HealthCheck.BasicAuthPassword, "health-check-basic-auth-password", "", p.HealthCheck.BasicAuthPassword, "")
	fs.StringVarP(&p.HealthCheck.QName, "health-check-q-name", "", p.HealthCheck.QName, "")
	fs.StringVarP(&p.HealthCheck.ExpectedData, "health-check-expected-data", "", p.HealthCheck.ExpectedData, "")
	fs.StringVarP(&p.HealthCheck.Community, "health-check-community", "", p.HealthCheck.Community, "")
	fs.StringVarP(&p.HealthCheck.SNMPVersion, "health-check-snmp-version", "", p.HealthCheck.SNMPVersion, "")
	fs.StringVarP(&p.HealthCheck.OID, "health-check-oid", "", p.HealthCheck.OID, "")
	fs.IntVarP(&p.HealthCheck.RemainingDays, "health-check-remaining-days", "", p.HealthCheck.RemainingDays, "")
	fs.BoolVarP(&p.HealthCheck.HTTP2, "health-check-http2", "", p.HealthCheck.HTTP2, "")
	fs.StringVarP(&p.HealthCheck.FTPS, "health-check-ftps", "", p.HealthCheck.FTPS, "options: [explicit/implicit]")
	fs.BoolVarP(&p.HealthCheck.VerifySNI, "health-check-verify-sni", "", p.HealthCheck.VerifySNI, "")
	fs.BoolVarP(&p.NotifyEmailEnabled, "notify-email-enabled", "", p.NotifyEmailEnabled, "")
	fs.BoolVarP(&p.NotifyEmailHTML, "notify-email-html", "", p.NotifyEmailHTML, "")
	fs.BoolVarP(&p.NotifySlackEnabled, "notify-slack-enabled", "", p.NotifySlackEnabled, "")
	fs.StringVarP(&p.SlackWebhooksURL, "slack-webhooks-url", "", p.SlackWebhooksURL, "")
	fs.IntVarP(&p.NotifyInterval, "notify-interval", "", p.NotifyInterval, "")
	fs.SetNormalizeFunc(p.normalizeFlagName)
}

func (p *createParameter) normalizeFlagName(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "skeleton":
		name = "generate-skeleton"
	case "out":
		name = "output-type"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func (p *createParameter) buildFlagsUsage(cmd *cobra.Command) {
	var sets []*core.FlagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("common", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("description"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("icon-id"))
		sets = append(sets, &core.FlagSet{
			Title: "Common options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("simple-monitor", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("delay-loop"))
		fs.AddFlag(cmd.LocalFlags().Lookup("enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-basic-auth-password"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-basic-auth-username"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-community"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-contains-string"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-expected-data"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-ftps"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-host"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-http2"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-oid"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-path"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-port"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-protocol"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-q-name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-remaining-days"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-sni"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-snmp-version"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-status"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-verify-sni"))
		fs.AddFlag(cmd.LocalFlags().Lookup("notify-email-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("notify-email-html"))
		fs.AddFlag(cmd.LocalFlags().Lookup("notify-interval"))
		fs.AddFlag(cmd.LocalFlags().Lookup("notify-slack-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("slack-webhooks-url"))
		fs.AddFlag(cmd.LocalFlags().Lookup("target"))
		fs.AddFlag(cmd.LocalFlags().Lookup("timeout"))
		sets = append(sets, &core.FlagSet{
			Title: "Simple-Monitor-specific options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("input", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		sets = append(sets, &core.FlagSet{
			Title: "Input options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-driver"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		sets = append(sets, &core.FlagSet{
			Title: "Output options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("example", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("example"))
		sets = append(sets, &core.FlagSet{
			Title: "Parameter example",
			Flags: fs,
		})
	}

	core.BuildFlagsUsage(cmd, sets)
}

func (p *createParameter) setCompletionFunc(cmd *cobra.Command) {
	cmd.RegisterFlagCompletionFunc("health-check-protocol", util.FlagCompletionFunc("http", "https", "ping", "tcp", "dns", "ssh", "smtp", "pop3", "snmp", "sslcertificate", "ftp"))
	cmd.RegisterFlagCompletionFunc("health-check-ftps", util.FlagCompletionFunc("explicit", "implicit"))

}

func (p *createParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
