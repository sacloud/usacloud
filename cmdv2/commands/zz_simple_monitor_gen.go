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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/spf13/cobra"
)

// simpleMonitorCmd represents the command to manage SAKURA Cloud SimpleMonitor
func simpleMonitorCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "simple-monitor",
		Short: "A manage commands of SimpleMonitor",
		Long:  `A manage commands of SimpleMonitor`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
}

func simpleMonitorListCmd() *cobra.Command {
	simpleMonitorListParam := params.NewListSimpleMonitorParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find", "selector"},
		Short:   "List SimpleMonitor",
		Long:    `List SimpleMonitor`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return simpleMonitorListParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), simpleMonitorListParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("list local parameter: \n%s\n", debugMarshalIndent(simpleMonitorListParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&simpleMonitorListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &simpleMonitorListParam.Id), "id", "", "set filter by id(s)")
	fs.StringSliceVarP(&simpleMonitorListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.StringVarP(&simpleMonitorListParam.Health, "health", "", "", "set filter by HealthCheck Status('up' or 'down' or 'unknown')")
	fs.IntVarP(&simpleMonitorListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&simpleMonitorListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&simpleMonitorListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&simpleMonitorListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simpleMonitorListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simpleMonitorListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simpleMonitorListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simpleMonitorListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&simpleMonitorListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&simpleMonitorListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&simpleMonitorListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&simpleMonitorListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&simpleMonitorListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&simpleMonitorListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&simpleMonitorListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	return cmd
}

func simpleMonitorCreateCmd() *cobra.Command {
	simpleMonitorCreateParam := params.NewCreateSimpleMonitorParam()
	cmd := &cobra.Command{
		Use: "create",

		Short: "Create SimpleMonitor",
		Long:  `Create SimpleMonitor`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return simpleMonitorCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), simpleMonitorCreateParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("create local parameter: \n%s\n", debugMarshalIndent(simpleMonitorCreateParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&simpleMonitorCreateParam.Target, "target", "", "", "set monitoring target IP or Hostname")
	fs.StringVarP(&simpleMonitorCreateParam.Protocol, "protocol", "", "ping", "set monitoring protocol[http/https/ping/tcp/dns/ssh/smtp/pop3/ssl-certificate]")
	fs.IntVarP(&simpleMonitorCreateParam.Port, "port", "", 0, "set port of tcp monitoring")
	fs.IntVarP(&simpleMonitorCreateParam.DelayLoop, "delay-loop", "", 1, "set delay-loop of monitoring(minute)")
	fs.BoolVarP(&simpleMonitorCreateParam.Disabled, "disabled", "", false, "set monitoring disable")
	fs.StringVarP(&simpleMonitorCreateParam.HostHeader, "host-header", "", "", "set host header of http/https monitoring request")
	fs.StringVarP(&simpleMonitorCreateParam.Path, "path", "", "", "set path of http/https monitoring request")
	fs.IntVarP(&simpleMonitorCreateParam.ResponseCode, "response-code", "", 0, "set response-code of http/https monitoring request")
	fs.BoolVarP(&simpleMonitorCreateParam.Sni, "sni", "", false, "enable SNI support for https monitoring")
	fs.StringVarP(&simpleMonitorCreateParam.Username, "username", "", "", "set Basic Auth user name")
	fs.StringVarP(&simpleMonitorCreateParam.Password, "password", "", "", "set Basic Auth password")
	fs.StringVarP(&simpleMonitorCreateParam.DNSQname, "dns-qname", "", "", "set DNS query target name")
	fs.StringVarP(&simpleMonitorCreateParam.DNSExcepted, "dns-excepted", "", "", "set DNS query excepted value")
	fs.IntVarP(&simpleMonitorCreateParam.RemainingDays, "remaining-days", "", 30, "set SSL-Certificate remaining days")
	fs.BoolVarP(&simpleMonitorCreateParam.NotifyEmail, "notify-email", "", true, "enable e-mail notification")
	fs.StringVarP(&simpleMonitorCreateParam.EmailType, "email-type", "", "text", "set e-mail type")
	fs.StringVarP(&simpleMonitorCreateParam.SlackWebhook, "slack-webhook", "", "", "set slack-webhook URL")
	fs.IntVarP(&simpleMonitorCreateParam.NotifyInterval, "notify-interval", "", 2, "set notify-interval(hours)")
	fs.StringVarP(&simpleMonitorCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&simpleMonitorCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &simpleMonitorCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&simpleMonitorCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&simpleMonitorCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simpleMonitorCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simpleMonitorCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simpleMonitorCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simpleMonitorCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&simpleMonitorCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&simpleMonitorCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&simpleMonitorCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&simpleMonitorCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&simpleMonitorCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&simpleMonitorCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&simpleMonitorCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	return cmd
}

func simpleMonitorReadCmd() *cobra.Command {
	simpleMonitorReadParam := params.NewReadSimpleMonitorParam()
	cmd := &cobra.Command{
		Use: "read",

		Short: "Read SimpleMonitor",
		Long:  `Read SimpleMonitor`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return simpleMonitorReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), simpleMonitorReadParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("read local parameter: \n%s\n", debugMarshalIndent(simpleMonitorReadParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&simpleMonitorReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&simpleMonitorReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simpleMonitorReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simpleMonitorReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simpleMonitorReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simpleMonitorReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&simpleMonitorReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&simpleMonitorReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&simpleMonitorReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&simpleMonitorReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&simpleMonitorReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&simpleMonitorReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&simpleMonitorReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &simpleMonitorReadParam.Id), "id", "", "Set target ID")
	return cmd
}

func simpleMonitorUpdateCmd() *cobra.Command {
	simpleMonitorUpdateParam := params.NewUpdateSimpleMonitorParam()
	cmd := &cobra.Command{
		Use: "update",

		Short: "Update SimpleMonitor",
		Long:  `Update SimpleMonitor`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return simpleMonitorUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), simpleMonitorUpdateParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("update local parameter: \n%s\n", debugMarshalIndent(simpleMonitorUpdateParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&simpleMonitorUpdateParam.Protocol, "protocol", "", "", "set monitoring protocol[http/https/ping/tcp/dns/ssh/smtp/pop3/ssl-certificate]")
	fs.IntVarP(&simpleMonitorUpdateParam.Port, "port", "", 0, "set port of tcp monitoring")
	fs.IntVarP(&simpleMonitorUpdateParam.DelayLoop, "delay-loop", "", 0, "set delay-loop of monitoring(minute)")
	fs.BoolVarP(&simpleMonitorUpdateParam.Disabled, "disabled", "", false, "set monitoring enable/disable")
	fs.StringVarP(&simpleMonitorUpdateParam.HostHeader, "host-header", "", "", "set host header of http/https monitoring request")
	fs.StringVarP(&simpleMonitorUpdateParam.Path, "path", "", "", "set path of http/https monitoring request")
	fs.IntVarP(&simpleMonitorUpdateParam.ResponseCode, "response-code", "", 0, "set response-code of http/https monitoring request")
	fs.BoolVarP(&simpleMonitorUpdateParam.Sni, "sni", "", false, "enable SNI support for https monitoring")
	fs.StringVarP(&simpleMonitorUpdateParam.Username, "username", "", "", "set Basic Auth user name")
	fs.StringVarP(&simpleMonitorUpdateParam.Password, "password", "", "", "set Basic Auth password")
	fs.StringVarP(&simpleMonitorUpdateParam.DNSQname, "dns-qname", "", "", "set DNS query target name")
	fs.StringVarP(&simpleMonitorUpdateParam.DNSExcepted, "dns-excepted", "", "", "set DNS query excepted value")
	fs.IntVarP(&simpleMonitorUpdateParam.RemainingDays, "remaining-days", "", 0, "set SSL-Certificate remaining days")
	fs.BoolVarP(&simpleMonitorUpdateParam.NotifyEmail, "notify-email", "", false, "enable e-mail notification")
	fs.StringVarP(&simpleMonitorUpdateParam.EmailType, "email-type", "", "", "set e-mail type")
	fs.StringVarP(&simpleMonitorUpdateParam.SlackWebhook, "slack-webhook", "", "", "set slack-webhook URL")
	fs.IntVarP(&simpleMonitorUpdateParam.NotifyInterval, "notify-interval", "", 2, "set notify-interval(hours)")
	fs.StringSliceVarP(&simpleMonitorUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&simpleMonitorUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&simpleMonitorUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &simpleMonitorUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&simpleMonitorUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&simpleMonitorUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simpleMonitorUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simpleMonitorUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simpleMonitorUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simpleMonitorUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&simpleMonitorUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&simpleMonitorUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&simpleMonitorUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&simpleMonitorUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&simpleMonitorUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&simpleMonitorUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&simpleMonitorUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &simpleMonitorUpdateParam.Id), "id", "", "Set target ID")
	return cmd
}

func simpleMonitorDeleteCmd() *cobra.Command {
	simpleMonitorDeleteParam := params.NewDeleteSimpleMonitorParam()
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete SimpleMonitor",
		Long:    `Delete SimpleMonitor`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return simpleMonitorDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), simpleMonitorDeleteParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("delete local parameter: \n%s\n", debugMarshalIndent(simpleMonitorDeleteParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&simpleMonitorDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&simpleMonitorDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&simpleMonitorDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simpleMonitorDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simpleMonitorDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simpleMonitorDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simpleMonitorDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&simpleMonitorDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&simpleMonitorDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&simpleMonitorDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&simpleMonitorDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&simpleMonitorDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&simpleMonitorDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&simpleMonitorDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &simpleMonitorDeleteParam.Id), "id", "", "Set target ID")
	return cmd
}

func simpleMonitorHealthCmd() *cobra.Command {
	simpleMonitorHealthParam := params.NewHealthSimpleMonitorParam()
	cmd := &cobra.Command{
		Use: "health",

		Short: "Health SimpleMonitor",
		Long:  `Health SimpleMonitor`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return simpleMonitorHealthParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), simpleMonitorHealthParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("health local parameter: \n%s\n", debugMarshalIndent(simpleMonitorHealthParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&simpleMonitorHealthParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&simpleMonitorHealthParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simpleMonitorHealthParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simpleMonitorHealthParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simpleMonitorHealthParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simpleMonitorHealthParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&simpleMonitorHealthParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&simpleMonitorHealthParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&simpleMonitorHealthParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&simpleMonitorHealthParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&simpleMonitorHealthParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&simpleMonitorHealthParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&simpleMonitorHealthParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &simpleMonitorHealthParam.Id), "id", "", "Set target ID")
	return cmd
}

func init() {
	parent := simpleMonitorCmd()
	parent.AddCommand(simpleMonitorListCmd())
	parent.AddCommand(simpleMonitorCreateCmd())
	parent.AddCommand(simpleMonitorReadCmd())
	parent.AddCommand(simpleMonitorUpdateCmd())
	parent.AddCommand(simpleMonitorDeleteCmd())
	parent.AddCommand(simpleMonitorHealthCmd())
	rootCmd.AddCommand(parent)
}
