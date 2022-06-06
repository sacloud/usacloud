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

package cli

import (
	"fmt"
	"io"
	"log"
	"net/http"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/helper/api"
	"github.com/sacloud/usacloud/pkg/config"
	"github.com/sacloud/webaccel-api-go"
)

type apiClient struct {
	option *config.Config

	iaasClient     iaas.APICaller
	webaccelClient *webaccel.Client
	commonClient   client.HttpRequestDoer
}

func newAPIClient(o *config.Config) *apiClient {
	c := &apiClient{
		option: o,
	}

	if o.FakeMode {
		// libsacloud fakeドライバはlogパッケージにシステムログを出すがusacloudからは利用しないため出力を抑制する
		log.SetOutput(io.Discard)
	}

	clientOption := &client.Options{
		AccessToken:          o.AccessToken,
		AccessTokenSecret:    o.AccessTokenSecret,
		AcceptLanguage:       o.AcceptLanguage,
		HttpClient:           http.DefaultClient,
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
	c.commonClient = client.NewFactory(clientOption).NewHttpRequestDoer()
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
