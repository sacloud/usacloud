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

var (
	proxyLBListParam                 = params.NewListProxyLBParam()
	proxyLBCreateParam               = params.NewCreateProxyLBParam()
	proxyLBReadParam                 = params.NewReadProxyLBParam()
	proxyLBUpdateParam               = params.NewUpdateProxyLBParam()
	proxyLBDeleteParam               = params.NewDeleteProxyLBParam()
	proxyLBPlanChangeParam           = params.NewPlanChangeProxyLBParam()
	proxyLBBindPortInfoParam         = params.NewBindPortInfoProxyLBParam()
	proxyLBBindPortAddParam          = params.NewBindPortAddProxyLBParam()
	proxyLBBindPortUpdateParam       = params.NewBindPortUpdateProxyLBParam()
	proxyLBBindPortDeleteParam       = params.NewBindPortDeleteProxyLBParam()
	proxyLBResponseHeaderInfoParam   = params.NewResponseHeaderInfoProxyLBParam()
	proxyLBResponseHeaderAddParam    = params.NewResponseHeaderAddProxyLBParam()
	proxyLBResponseHeaderUpdateParam = params.NewResponseHeaderUpdateProxyLBParam()
	proxyLBResponseHeaderDeleteParam = params.NewResponseHeaderDeleteProxyLBParam()
	proxyLBACMEInfoParam             = params.NewACMEInfoProxyLBParam()
	proxyLBACMESettingParam          = params.NewACMESettingProxyLBParam()
	proxyLBACMERenewParam            = params.NewACMERenewProxyLBParam()
	proxyLBServerInfoParam           = params.NewServerInfoProxyLBParam()
	proxyLBServerAddParam            = params.NewServerAddProxyLBParam()
	proxyLBServerUpdateParam         = params.NewServerUpdateProxyLBParam()
	proxyLBServerDeleteParam         = params.NewServerDeleteProxyLBParam()
	proxyLBCertificateInfoParam      = params.NewCertificateInfoProxyLBParam()
	proxyLBCertificateAddParam       = params.NewCertificateAddProxyLBParam()
	proxyLBCertificateUpdateParam    = params.NewCertificateUpdateProxyLBParam()
	proxyLBCertificateDeleteParam    = params.NewCertificateDeleteProxyLBParam()
	proxyLBMonitorParam              = params.NewMonitorProxyLBParam()
)

// proxyLBCmd represents the command to manage SAKURA Cloud ProxyLB
var proxyLBCmd = &cobra.Command{
	Use:   "proxyLB",
	Short: "A manage commands of ProxyLB",
	Long:  `A manage commands of ProxyLB`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var proxyLBListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find", "selector"},
	Short:   "List ProxyLB",
	Long:    `List ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(proxyLBListParam))
		return err
	},
}

func proxyLBListCmdInit() {
	fs := proxyLBListCmd.Flags()
	fs.IntVarP(&proxyLBListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&proxyLBListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&proxyLBListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringSliceVarP(&proxyLBListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.StringSliceVarP(&proxyLBListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &proxyLBListParam.Id), "id", "", "set filter by id(s)")
}

var proxyLBCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create ProxyLB",
	Long:  `Create ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("create parameter: \n%s\n", debugMarshalIndent(proxyLBCreateParam))
		return err
	},
}

func proxyLBCreateCmdInit() {
	fs := proxyLBCreateCmd.Flags()
	fs.StringVarP(&proxyLBCreateParam.Path, "path", "", "/", "set path of http/https healthcheck request")
	fs.IntVarP(&proxyLBCreateParam.DelayLoop, "delay-loop", "", 10, "set delay-loop of healthcheck")
	fs.StringVarP(&proxyLBCreateParam.SorryServerIpaddress, "sorry-server-ipaddress", "", "", "set sorry-server ip address")
	fs.IntVarP(&proxyLBCreateParam.SorryServerPort, "sorry-server-port", "", 0, "set sorry-server ports")
	fs.IntVarP(&proxyLBCreateParam.Timeout, "timeout", "", 10, "set timeout")
	fs.StringVarP(&proxyLBCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&proxyLBCreateParam.Protocol, "protocol", "", "tcp", "set healthcheck protocol[http/tcp]")
	fs.StringVarP(&proxyLBCreateParam.HostHeader, "host-header", "", "", "set host header of http/https healthcheck request")
	fs.StringVarP(&proxyLBCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&proxyLBCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &proxyLBCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.IntVarP(&proxyLBCreateParam.Plan, "plan", "", 1000, "set plan")
	fs.BoolVarP(&proxyLBCreateParam.StickySession, "sticky-session", "", false, "enable sticky-session")
}

var proxyLBReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read ProxyLB",
	Long:  `Read ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(proxyLBReadParam))
		return err
	},
}

func proxyLBReadCmdInit() {
}

var proxyLBUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update ProxyLB",
	Long:  `Update ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("update parameter: \n%s\n", debugMarshalIndent(proxyLBUpdateParam))
		return err
	},
}

func proxyLBUpdateCmdInit() {
	fs := proxyLBUpdateCmd.Flags()
	fs.StringVarP(&proxyLBUpdateParam.HostHeader, "host-header", "", "", "set host header of http/https healthcheck request")
	fs.IntVarP(&proxyLBUpdateParam.DelayLoop, "delay-loop", "", 0, "set delay-loop of healthcheck")
	fs.BoolVarP(&proxyLBUpdateParam.StickySession, "sticky-session", "", false, "enable sticky-session")
	fs.IntVarP(&proxyLBUpdateParam.Timeout, "timeout", "", 10, "set timeout")
	fs.StringVarP(&proxyLBUpdateParam.Name, "name", "", "", "set resource display name")
	fs.VarP(newIDValue(0, &proxyLBUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.StringVarP(&proxyLBUpdateParam.Protocol, "protocol", "", "", "set healthcheck protocol[http/tcp]")
	fs.StringVarP(&proxyLBUpdateParam.Path, "path", "", "", "set path of http/https healthcheck request")
	fs.StringVarP(&proxyLBUpdateParam.SorryServerIpaddress, "sorry-server-ipaddress", "", "", "set sorry-server ip address")
	fs.IntVarP(&proxyLBUpdateParam.SorryServerPort, "sorry-server-port", "", 0, "set sorry-server ports")
	fs.StringVarP(&proxyLBUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&proxyLBUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
}

var proxyLBDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete ProxyLB",
	Long:    `Delete ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(proxyLBDeleteParam))
		return err
	},
}

func proxyLBDeleteCmdInit() {
}

var proxyLBPlanChangeCmd = &cobra.Command{
	Use: "plan-change",

	Short: "Change ProxyLB plan",
	Long:  `Change ProxyLB plan`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBPlanChangeParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("plan-change parameter: \n%s\n", debugMarshalIndent(proxyLBPlanChangeParam))
		return err
	},
}

func proxyLBPlanChangeCmdInit() {
	fs := proxyLBPlanChangeCmd.Flags()
	fs.IntVarP(&proxyLBPlanChangeParam.Plan, "plan", "", 0, "set plan")
}

var proxyLBBindPortInfoCmd = &cobra.Command{
	Use:     "bind-port-info",
	Aliases: []string{"bind-port-list"},
	Short:   "BindPortInfo ProxyLB",
	Long:    `BindPortInfo ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBBindPortInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("bind-port-info parameter: \n%s\n", debugMarshalIndent(proxyLBBindPortInfoParam))
		return err
	},
}

func proxyLBBindPortInfoCmdInit() {
}

var proxyLBBindPortAddCmd = &cobra.Command{
	Use: "bind-port-add",

	Short: "BindPortAdd ProxyLB",
	Long:  `BindPortAdd ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBBindPortAddParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("bind-port-add parameter: \n%s\n", debugMarshalIndent(proxyLBBindPortAddParam))
		return err
	},
}

func proxyLBBindPortAddCmdInit() {
	fs := proxyLBBindPortAddCmd.Flags()
	fs.StringVarP(&proxyLBBindPortAddParam.Mode, "mode", "", "", "set bind mode[http/https/tcp]")
	fs.IntVarP(&proxyLBBindPortAddParam.Port, "port", "", 0, "set port number")
	fs.BoolVarP(&proxyLBBindPortAddParam.RedirectToHttps, "redirect-to-https", "", false, "enable to redirect to https")
	fs.BoolVarP(&proxyLBBindPortAddParam.SupportHttp2, "support-http-2", "", false, "enable http/2")
}

var proxyLBBindPortUpdateCmd = &cobra.Command{
	Use: "bind-port-update",

	Short: "BindPortUpdate ProxyLB",
	Long:  `BindPortUpdate ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBBindPortUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("bind-port-update parameter: \n%s\n", debugMarshalIndent(proxyLBBindPortUpdateParam))
		return err
	},
}

func proxyLBBindPortUpdateCmdInit() {
	fs := proxyLBBindPortUpdateCmd.Flags()
	fs.IntVarP(&proxyLBBindPortUpdateParam.Index, "index", "", 0, "index of target server")
	fs.StringVarP(&proxyLBBindPortUpdateParam.Mode, "mode", "", "", "set bind mode[http/https/tcp]")
	fs.IntVarP(&proxyLBBindPortUpdateParam.Port, "port", "", 0, "set port number")
	fs.BoolVarP(&proxyLBBindPortUpdateParam.RedirectToHttps, "redirect-to-https", "", false, "enable to redirect to https")
	fs.BoolVarP(&proxyLBBindPortUpdateParam.SupportHttp2, "support-http-2", "", false, "enable http/2")
}

var proxyLBBindPortDeleteCmd = &cobra.Command{
	Use: "bind-port-delete",

	Short: "BindPortDelete ProxyLB",
	Long:  `BindPortDelete ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBBindPortDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("bind-port-delete parameter: \n%s\n", debugMarshalIndent(proxyLBBindPortDeleteParam))
		return err
	},
}

func proxyLBBindPortDeleteCmdInit() {
	fs := proxyLBBindPortDeleteCmd.Flags()
	fs.IntVarP(&proxyLBBindPortDeleteParam.Index, "index", "", 0, "index of target bind-port")
}

var proxyLBResponseHeaderInfoCmd = &cobra.Command{
	Use:     "response-header-info",
	Aliases: []string{"response-header-list"},
	Short:   "ResponseHeaderInfo ProxyLB",
	Long:    `ResponseHeaderInfo ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBResponseHeaderInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("response-header-info parameter: \n%s\n", debugMarshalIndent(proxyLBResponseHeaderInfoParam))
		return err
	},
}

func proxyLBResponseHeaderInfoCmdInit() {
	fs := proxyLBResponseHeaderInfoCmd.Flags()
	fs.IntVarP(&proxyLBResponseHeaderInfoParam.PortIndex, "port-index", "", 0, "index of target bind-port")
}

var proxyLBResponseHeaderAddCmd = &cobra.Command{
	Use: "response-header-add",

	Short: "ResponseHeaderAdd ProxyLB",
	Long:  `ResponseHeaderAdd ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBResponseHeaderAddParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("response-header-add parameter: \n%s\n", debugMarshalIndent(proxyLBResponseHeaderAddParam))
		return err
	},
}

func proxyLBResponseHeaderAddCmdInit() {
	fs := proxyLBResponseHeaderAddCmd.Flags()
	fs.IntVarP(&proxyLBResponseHeaderAddParam.PortIndex, "port-index", "", 0, "index of target bind-port")
	fs.StringVarP(&proxyLBResponseHeaderAddParam.Header, "header", "", "", "set Header")
	fs.StringVarP(&proxyLBResponseHeaderAddParam.Value, "value", "", "", "set Value")
}

var proxyLBResponseHeaderUpdateCmd = &cobra.Command{
	Use: "response-header-update",

	Short: "ResponseHeaderUpdate ProxyLB",
	Long:  `ResponseHeaderUpdate ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBResponseHeaderUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("response-header-update parameter: \n%s\n", debugMarshalIndent(proxyLBResponseHeaderUpdateParam))
		return err
	},
}

func proxyLBResponseHeaderUpdateCmdInit() {
	fs := proxyLBResponseHeaderUpdateCmd.Flags()
	fs.IntVarP(&proxyLBResponseHeaderUpdateParam.Index, "index", "", 0, "index of target server")
	fs.IntVarP(&proxyLBResponseHeaderUpdateParam.PortIndex, "port-index", "", 0, "index of target bind-port")
	fs.StringVarP(&proxyLBResponseHeaderUpdateParam.Header, "header", "", "", "set Header")
	fs.StringVarP(&proxyLBResponseHeaderUpdateParam.Value, "value", "", "", "set Value")
}

var proxyLBResponseHeaderDeleteCmd = &cobra.Command{
	Use: "response-header-delete",

	Short: "ResponseHeaderDelete ProxyLB",
	Long:  `ResponseHeaderDelete ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBResponseHeaderDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("response-header-delete parameter: \n%s\n", debugMarshalIndent(proxyLBResponseHeaderDeleteParam))
		return err
	},
}

func proxyLBResponseHeaderDeleteCmdInit() {
	fs := proxyLBResponseHeaderDeleteCmd.Flags()
	fs.IntVarP(&proxyLBResponseHeaderDeleteParam.Index, "index", "", 0, "index of target bind-port")
	fs.IntVarP(&proxyLBResponseHeaderDeleteParam.PortIndex, "port-index", "", 0, "index of target bind-port")
}

var proxyLBACMEInfoCmd = &cobra.Command{
	Use: "acme-info",

	Short: "ACMEInfo ProxyLB",
	Long:  `ACMEInfo ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBACMEInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("acme-info parameter: \n%s\n", debugMarshalIndent(proxyLBACMEInfoParam))
		return err
	},
}

func proxyLBACMEInfoCmdInit() {
}

var proxyLBACMESettingCmd = &cobra.Command{
	Use: "acme-setting",

	Short: "ACMESetting ProxyLB",
	Long:  `ACMESetting ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBACMESettingParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("acme-setting parameter: \n%s\n", debugMarshalIndent(proxyLBACMESettingParam))
		return err
	},
}

func proxyLBACMESettingCmdInit() {
	fs := proxyLBACMESettingCmd.Flags()
	fs.StringVarP(&proxyLBACMESettingParam.CommonName, "common-name", "", "", "set common name")
	fs.BoolVarP(&proxyLBACMESettingParam.Disable, "disable", "", false, "the flag of disable Let's Encrypt")
	fs.BoolVarP(&proxyLBACMESettingParam.AcceptTos, "accept-tos", "", false, "the flag of accept Let's Encrypt's terms of services: https://letsencrypt.org/repository/")
}

var proxyLBACMERenewCmd = &cobra.Command{
	Use: "acme-renew",

	Short: "ACMERenew ProxyLB",
	Long:  `ACMERenew ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBACMERenewParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("acme-renew parameter: \n%s\n", debugMarshalIndent(proxyLBACMERenewParam))
		return err
	},
}

func proxyLBACMERenewCmdInit() {
}

var proxyLBServerInfoCmd = &cobra.Command{
	Use:     "server-info",
	Aliases: []string{"server-list"},
	Short:   "ServerInfo ProxyLB",
	Long:    `ServerInfo ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBServerInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("server-info parameter: \n%s\n", debugMarshalIndent(proxyLBServerInfoParam))
		return err
	},
}

func proxyLBServerInfoCmdInit() {
}

var proxyLBServerAddCmd = &cobra.Command{
	Use: "server-add",

	Short: "ServerAdd ProxyLB",
	Long:  `ServerAdd ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBServerAddParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("server-add parameter: \n%s\n", debugMarshalIndent(proxyLBServerAddParam))
		return err
	},
}

func proxyLBServerAddCmdInit() {
	fs := proxyLBServerAddCmd.Flags()
	fs.StringVarP(&proxyLBServerAddParam.Ipaddress, "ipaddress", "", "", "set target ipaddress")
	fs.BoolVarP(&proxyLBServerAddParam.Disabled, "disabled", "", false, "set disabled")
	fs.IntVarP(&proxyLBServerAddParam.Port, "port", "", 0, "set server ports")
}

var proxyLBServerUpdateCmd = &cobra.Command{
	Use: "server-update",

	Short: "ServerUpdate ProxyLB",
	Long:  `ServerUpdate ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBServerUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("server-update parameter: \n%s\n", debugMarshalIndent(proxyLBServerUpdateParam))
		return err
	},
}

func proxyLBServerUpdateCmdInit() {
	fs := proxyLBServerUpdateCmd.Flags()
	fs.IntVarP(&proxyLBServerUpdateParam.Index, "index", "", 0, "index of target server")
	fs.StringVarP(&proxyLBServerUpdateParam.Ipaddress, "ipaddress", "", "", "set target ipaddress")
	fs.BoolVarP(&proxyLBServerUpdateParam.Disabled, "disabled", "", false, "set disabled")
	fs.IntVarP(&proxyLBServerUpdateParam.Port, "port", "", 0, "set server ports")
}

var proxyLBServerDeleteCmd = &cobra.Command{
	Use: "server-delete",

	Short: "ServerDelete ProxyLB",
	Long:  `ServerDelete ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBServerDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("server-delete parameter: \n%s\n", debugMarshalIndent(proxyLBServerDeleteParam))
		return err
	},
}

func proxyLBServerDeleteCmdInit() {
	fs := proxyLBServerDeleteCmd.Flags()
	fs.IntVarP(&proxyLBServerDeleteParam.Index, "index", "", 0, "index of target server")
}

var proxyLBCertificateInfoCmd = &cobra.Command{
	Use:     "certificate-info",
	Aliases: []string{"certificate-list", "cert-list", "cert-info"},
	Short:   "CertificateInfo ProxyLB",
	Long:    `CertificateInfo ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBCertificateInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("certificate-info parameter: \n%s\n", debugMarshalIndent(proxyLBCertificateInfoParam))
		return err
	},
}

func proxyLBCertificateInfoCmdInit() {
}

var proxyLBCertificateAddCmd = &cobra.Command{
	Use:     "certificate-add",
	Aliases: []string{"cert-add"},
	Short:   "CertificateAdd ProxyLB",
	Long:    `CertificateAdd ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBCertificateAddParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("certificate-add parameter: \n%s\n", debugMarshalIndent(proxyLBCertificateAddParam))
		return err
	},
}

func proxyLBCertificateAddCmdInit() {
	fs := proxyLBCertificateAddCmd.Flags()
	fs.StringVarP(&proxyLBCertificateAddParam.ServerCertificate, "server-certificate", "", "", "")
	fs.StringVarP(&proxyLBCertificateAddParam.IntermediateCertificate, "intermediate-certificate", "", "", "")
	fs.StringVarP(&proxyLBCertificateAddParam.PrivateKey, "private-key", "", "", "")
}

var proxyLBCertificateUpdateCmd = &cobra.Command{
	Use:     "certificate-update",
	Aliases: []string{"cert-update"},
	Short:   "CertificateUpdate ProxyLB",
	Long:    `CertificateUpdate ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBCertificateUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("certificate-update parameter: \n%s\n", debugMarshalIndent(proxyLBCertificateUpdateParam))
		return err
	},
}

func proxyLBCertificateUpdateCmdInit() {
	fs := proxyLBCertificateUpdateCmd.Flags()
	fs.StringVarP(&proxyLBCertificateUpdateParam.PrivateKey, "private-key", "", "", "")
	fs.StringVarP(&proxyLBCertificateUpdateParam.ServerCertificate, "server-certificate", "", "", "")
	fs.StringVarP(&proxyLBCertificateUpdateParam.IntermediateCertificate, "intermediate-certificate", "", "", "")
}

var proxyLBCertificateDeleteCmd = &cobra.Command{
	Use:     "certificate-delete",
	Aliases: []string{"cert-delete"},
	Short:   "CertificateDelete ProxyLB",
	Long:    `CertificateDelete ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBCertificateDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("certificate-delete parameter: \n%s\n", debugMarshalIndent(proxyLBCertificateDeleteParam))
		return err
	},
}

func proxyLBCertificateDeleteCmdInit() {
}

var proxyLBMonitorCmd = &cobra.Command{
	Use: "monitor",

	Short: "Monitor ProxyLB",
	Long:  `Monitor ProxyLB`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := proxyLBMonitorParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor parameter: \n%s\n", debugMarshalIndent(proxyLBMonitorParam))
		return err
	},
}

func proxyLBMonitorCmdInit() {
	fs := proxyLBMonitorCmd.Flags()
	fs.StringVarP(&proxyLBMonitorParam.Start, "start", "", "", "set start-time")
	fs.StringVarP(&proxyLBMonitorParam.End, "end", "", "", "set end-time")
	fs.StringVarP(&proxyLBMonitorParam.KeyFormat, "key-format", "", "sakuracloud.proxylb.{{.ID}}", "set monitoring value key-format")
}

func init() {
	parent := proxyLBCmd

	proxyLBListCmdInit()
	parent.AddCommand(proxyLBListCmd)

	proxyLBCreateCmdInit()
	parent.AddCommand(proxyLBCreateCmd)

	proxyLBReadCmdInit()
	parent.AddCommand(proxyLBReadCmd)

	proxyLBUpdateCmdInit()
	parent.AddCommand(proxyLBUpdateCmd)

	proxyLBDeleteCmdInit()
	parent.AddCommand(proxyLBDeleteCmd)

	proxyLBPlanChangeCmdInit()
	parent.AddCommand(proxyLBPlanChangeCmd)

	proxyLBBindPortInfoCmdInit()
	parent.AddCommand(proxyLBBindPortInfoCmd)

	proxyLBBindPortAddCmdInit()
	parent.AddCommand(proxyLBBindPortAddCmd)

	proxyLBBindPortUpdateCmdInit()
	parent.AddCommand(proxyLBBindPortUpdateCmd)

	proxyLBBindPortDeleteCmdInit()
	parent.AddCommand(proxyLBBindPortDeleteCmd)

	proxyLBResponseHeaderInfoCmdInit()
	parent.AddCommand(proxyLBResponseHeaderInfoCmd)

	proxyLBResponseHeaderAddCmdInit()
	parent.AddCommand(proxyLBResponseHeaderAddCmd)

	proxyLBResponseHeaderUpdateCmdInit()
	parent.AddCommand(proxyLBResponseHeaderUpdateCmd)

	proxyLBResponseHeaderDeleteCmdInit()
	parent.AddCommand(proxyLBResponseHeaderDeleteCmd)

	proxyLBACMEInfoCmdInit()
	parent.AddCommand(proxyLBACMEInfoCmd)

	proxyLBACMESettingCmdInit()
	parent.AddCommand(proxyLBACMESettingCmd)

	proxyLBACMERenewCmdInit()
	parent.AddCommand(proxyLBACMERenewCmd)

	proxyLBServerInfoCmdInit()
	parent.AddCommand(proxyLBServerInfoCmd)

	proxyLBServerAddCmdInit()
	parent.AddCommand(proxyLBServerAddCmd)

	proxyLBServerUpdateCmdInit()
	parent.AddCommand(proxyLBServerUpdateCmd)

	proxyLBServerDeleteCmdInit()
	parent.AddCommand(proxyLBServerDeleteCmd)

	proxyLBCertificateInfoCmdInit()
	parent.AddCommand(proxyLBCertificateInfoCmd)

	proxyLBCertificateAddCmdInit()
	parent.AddCommand(proxyLBCertificateAddCmd)

	proxyLBCertificateUpdateCmdInit()
	parent.AddCommand(proxyLBCertificateUpdateCmd)

	proxyLBCertificateDeleteCmdInit()
	parent.AddCommand(proxyLBCertificateDeleteCmd)

	proxyLBMonitorCmdInit()
	parent.AddCommand(proxyLBMonitorCmd)

	rootCmd.AddCommand(parent)
}
