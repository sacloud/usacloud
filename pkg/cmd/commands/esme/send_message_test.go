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
	"errors"
	"strings"
	"testing"

	"github.com/sacloud/usacloud/pkg/validate"
	"github.com/stretchr/testify/require"
)

func TestSendMessageParameter_Validate(t *testing.T) {
	validate.InitializeValidator([]string{"is1a"})

	cases := []struct {
		in  *sendMessageParameter
		err error
	}{
		// default
		{
			in: newSendMessageParameter(),
			err: errors.New(strings.Join([]string{
				"validation error:",
				"\t--destination: required",
				"\t--sender: required",
			}, "\n")),
		},
		// valid
		{
			in: &sendMessageParameter{
				Destination: "819012345678",
				Sender:      "example",
			},
			err: nil,
		},
		// with invalid domain-name
		{
			in: &sendMessageParameter{
				Destination: "819012345678",
				Sender:      "example",
				DomainName:  "example",
			},
			err: errors.New(strings.Join([]string{
				"validation error:",
				"\t--domain-name: fqdn",
			}, "\n")),
		},
		// with valid domain-name
		{
			in: &sendMessageParameter{
				Destination: "819012345678",
				Sender:      "example",
				DomainName:  "www.example.com",
			},
			err: nil,
		},
		// with valid domain-name
		// see: https://github.com/sacloud/usacloud/pull/824
		{
			in: &sendMessageParameter{
				Destination: "819012345678",
				Sender:      "example",
				DomainName:  "www.example.com.",
			},
			err: nil,
		},
	}

	for _, tc := range cases {
		err := validate.Exec(tc.in)
		require.Equal(t, tc.err, err)
	}
}
