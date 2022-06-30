// Copyright 2017-2022 The sacloud/usacloud Authors
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

package e2e

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/helper/api"
	"github.com/sacloud/usacloud/pkg/version"
)

var SacloudAPICaller = api.NewCallerWithOptions(&api.CallerOptions{
	Options: &client.Options{
		AccessToken:       os.Getenv("SAKURACLOUD_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET"),
		UserAgent: fmt.Sprintf(
			"sacloud/usacloud@v%s:e2e-test (%s/%s; +https://github.com/sacloud/usacloud) %s",
			version.Version,
			runtime.GOOS,
			runtime.GOARCH,
			iaas.DefaultUserAgent,
		),
		HttpClient:           &http.Client{},
		HttpRequestTimeout:   300,
		HttpRequestRateLimit: 10,
		RetryMax:             10,
		Trace:                os.Getenv("SAKURACLOUD_TRACE") != "",
	},
	TraceAPI: os.Getenv("SAKURACLOUD_TRACE") != "",
})
