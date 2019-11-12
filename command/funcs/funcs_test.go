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

package funcs

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/sacloud/usacloud/command"
)

type dummyFlagContext struct{}

func (d *dummyFlagContext) IsSet(name string) bool {
	return true
}

var dummyContext command.Context

func TestMain(m *testing.M) {
	accessToken := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	if accessToken == "" || accessTokenSecret == "" {
		log.Println("Please Set ENV 'SAKURACLOUD_ACCESS_TOKEN' and 'SAKURACLOUD_ACCESS_TOKEN_SECRET'")
		os.Exit(0) // exit normal
	}
	zone := os.Getenv("SAKURACLOUD_ZONE")
	if zone == "" {
		zone = "tk1v"
	}

	command.GlobalOption.AccessToken = accessToken
	command.GlobalOption.AccessTokenSecret = accessTokenSecret
	command.GlobalOption.Zone = zone

	dummyContext = command.NewContext(&dummyFlagContext{}, []string{}, nil)
	dummyContext.GetAPIClient().UserAgent = fmt.Sprintf("usacloud-unit-test")

	ret := m.Run()
	os.Exit(ret)
}
