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

package test

import (
	"fmt"
	"os"
	"sync"

	"github.com/sacloud/libsacloud/v2"
	"github.com/sacloud/libsacloud/v2/helper/api"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/version"
)

var apiCaller sacloud.APICaller
var apiCallerInitOnce sync.Once

func APICaller() sacloud.APICaller {
	apiCallerInitOnce.Do(func() {
		apiCaller = api.NewCaller(&api.CallerOptions{
			AccessToken:          os.Getenv("SAKURACLOUD_ACCESS_TOKEN"),
			AccessTokenSecret:    os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET"),
			HTTPRequestRateLimit: 5,
			RetryMax:             10,
			RetryWaitMax:         30,
			RetryWaitMin:         1,
			UserAgent:            fmt.Sprintf("Usacloud(Test)/v%s (+https://github.com/sacloud/usacloud) libsacloud/%s", version.Version, libsacloud.Version),
			TraceAPI:             os.Getenv("SAKURACLOUD_TRACE") != "",
			TraceHTTP:            os.Getenv("SAKURACLOUD_TRACE") != "",
			FakeMode:             os.Getenv("FAKE_MODE") != "" || os.Getenv("TESTACC") != "",
		})
	})
	return apiCaller
}
