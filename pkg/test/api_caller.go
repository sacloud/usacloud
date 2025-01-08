// Copyright 2017-2025 The sacloud/usacloud Authors
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

package test

import (
	"fmt"
	"os"
	"sync"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/helper/api"
	"github.com/sacloud/usacloud/pkg/version"
)

var apiCaller iaas.APICaller
var apiCallerInitOnce sync.Once

func APICaller() iaas.APICaller {
	apiCallerInitOnce.Do(func() {
		apiCaller = api.NewCallerWithOptions(&api.CallerOptions{
			Options: &client.Options{
				AccessToken:          os.Getenv("SAKURACLOUD_ACCESS_TOKEN"),
				AccessTokenSecret:    os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET"),
				HttpRequestRateLimit: 5,
				RetryMax:             10,
				RetryWaitMax:         30,
				RetryWaitMin:         1,
				Trace:                os.Getenv("SAKURACLOUD_TRACE") != "",
				UserAgent:            fmt.Sprintf("Usacloud(Test)/v%s (+https://github.com/sacloud/usacloud) %s", version.Version, iaas.DefaultUserAgent),
			},
			TraceAPI: os.Getenv("SAKURACLOUD_TRACE") != "",
			FakeMode: os.Getenv("FAKE_MODE") != "" || os.Getenv("TESTACC") != "",
		})
	})
	return apiCaller
}
