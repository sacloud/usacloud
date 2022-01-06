// Copyright 2017-2022 The Usacloud Authors
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

package esme

import (
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/output"
)

var sendMessageCommand = &core.Command{
	Name:     "send-message",
	Aliases:  []string{"send"},
	Category: "operation",
	Order:    20,

	ColumnDefs: []output.ColumnDef{
		{Name: "MessageID"},
		{Name: "Status"},
		{Name: "OTP"},
	},

	SelectorType: core.SelectorTypeRequireSingle, // 2重送信などの事故を防ぐためにSingle

	ParameterInitializer: func() interface{} {
		return newSendMessageParameter()
	},
}

type sendMessageParameter struct {
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	Destination string `cli:",aliases=dest" validate:"required"`
	Sender      string `validate:"required"`
	DomainName  string `validate:"omitempty,fqdn"`
	OTP         string
}

func newSendMessageParameter() *sendMessageParameter {
	return &sendMessageParameter{}
}

func init() {
	Resource.AddCommand(sendMessageCommand)
}

func (p *sendMessageParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &sendMessageParameter{
		Destination: "81zzzzzzzzzz",
		Sender:      "example-sender",
		DomainName:  "www.example.com",
		OTP:         "your-otp",
	}
}
