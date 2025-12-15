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

package e2e

import (
	"fmt"
	"net/http"
	"runtime"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/helper/api"
	"github.com/sacloud/packages-go/envvar"
	"github.com/sacloud/usacloud/pkg/version"
)

var SacloudAPICaller = api.NewCallerWithOptions(&api.CallerOptions{
	Options: &client.Options{
		AccessToken:       envvar.StringFromEnvMulti([]string{"SAKURA_ACCESS_TOKEN", "SAKURACLOUD_ACCESS_TOKEN"}, ""),
		AccessTokenSecret: envvar.StringFromEnvMulti([]string{"SAKURA_ACCESS_TOKEN_SECRET", "SAKURACLOUD_ACCESS_TOKEN_SECRET"}, ""),
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
		Trace:                envvar.StringFromEnvMulti([]string{"SAKURA_TRACE", "SAKURACLOUD_TRACE"}, "") != "",
	},
	TraceAPI: envvar.StringFromEnvMulti([]string{"SAKURA_TRACE", "SAKURACLOUD_TRACE"}, "") != "",
})
