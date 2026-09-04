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

package cli

import (
	"fmt"
	"io"
	"log"

	"github.com/sacloud/sacloud-sdk-go/api/iaas"
	"github.com/sacloud/sacloud-sdk-go/api/iaas/helper/api"
	"github.com/sacloud/sacloud-sdk-go/api/webaccel"
	"github.com/sacloud/sacloud-sdk-go/common/saclient"
	"github.com/sacloud/usacloud/pkg/config"
)

type apiClient struct {
	option *config.Config

	iaasClient     iaas.APICaller
	webaccelClient *webaccel.Client
	commonClient   saclient.HttpRequestDoer
}

func newAPIClient(o *config.Config) *apiClient {
	c := &apiClient{
		option: o,
	}

	if o.FakeMode {
		// libsacloud fakeドライバはlogパッケージにシステムログを出すがusacloudからは利用しないため出力を抑制する
		log.SetOutput(io.Discard)
	}

	clientOption := &api.ClientOptions{
		AccessToken:          o.AccessToken,
		AccessTokenSecret:    o.AccessTokenSecret,
		AcceptLanguage:       o.AcceptLanguage,
		HttpRequestTimeout:   o.HTTPRequestTimeout,
		HttpRequestRateLimit: o.HTTPRequestRateLimit,
		RetryMax:             o.RetryMax,
		RetryWaitMax:         o.RetryWaitMax,
		RetryWaitMin:         o.RetryWaitMin,
		UserAgent:            UserAgent,
		Trace:                o.EnableHTTPTrace(),
	}

	c.iaasClient = api.NewCallerWithOptions(&api.CallerOptions{
		Options:       clientOption,
		APIRootURL:    o.APIRootURL,
		DefaultZone:   o.DefaultZone,
		TraceAPI:      o.EnableAPITrace(),
		FakeMode:      o.FakeMode,
		FakeStorePath: o.FakeStorePath,
	})

	c.webaccelClient = &webaccel.Client{
		Options: clientOption,
	}
	c.commonClient = saclient.NewFactory(clientOption)
	return c
}

func (c *apiClient) client(platformName string) interface{} {
	switch platformName {
	case "phy":
		panic("not yet implemented")
	case "objectstorage":
		panic("not yet implemented")
	case "iaas":
		return c.iaasClient
	case "webaccel":
		return c.webaccelClient
	case "":
		return c.commonClient
	}

	panic(fmt.Sprintf("unsupported platform: %s", platformName))
}
