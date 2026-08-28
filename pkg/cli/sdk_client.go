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
	"os"
	"strconv"
	"strings"
	"sync"

	legacysaclient "github.com/sacloud/saclient-go"
	sdksaclient "github.com/sacloud/sacloud-sdk-go/common/saclient"
	"github.com/sacloud/usacloud/pkg/config"
)

type sdkClientHolder struct {
	once         sync.Once
	option       *config.Config
	legacyClient legacysaclient.ClientAPI
	client       sdksaclient.ClientAPI
	err          error
}

func newSDKClientHolder(option *config.Config, legacyClient legacysaclient.ClientAPI) *sdkClientHolder {
	return &sdkClientHolder{
		option:       option,
		legacyClient: legacyClient,
	}
}

func (h *sdkClientHolder) Client() (sdksaclient.ClientAPI, error) {
	h.once.Do(func() {
		h.client, h.err = h.newClient()
	})
	return h.client, h.err
}

func (h *sdkClientHolder) newClient() (sdksaclient.ClientAPI, error) {
	env := append([]string(nil), os.Environ()...)
	env = setEnvValue(env, "SAKURA_ACCESS_TOKEN", h.option.AccessToken)
	env = setEnvValue(env, "SAKURA_ACCESS_TOKEN_SECRET", h.option.AccessTokenSecret)
	env = setEnvValue(env, "SAKURA_PRIVATE_KEY_PATH", h.option.PrivateKeyPEMPath)
	env = setEnvValue(env, "SAKURA_SERVICE_PRINCIPAL_ID", h.option.ServicePrincipalID)
	env = setEnvValue(env, "SAKURA_SERVICE_PRINCIPAL_KEY_KID", h.option.ServicePrincipalKeyID)
	env = setEnvValue(env, "SAKURA_TOKEN_ENDPOINT", h.option.TokenEndpoint)
	env = setEnvValue(env, "SAKURA_ACCEPT_LANGUAGE", h.option.AcceptLanguage)
	env = setEnvValue(env, "SAKURA_DEFAULT_ZONE", h.option.DefaultZone)
	env = setEnvIntValue(env, "SAKURA_RETRY_MAX", h.option.RetryMax)
	env = setEnvIntValue(env, "SAKURA_RETRY_WAIT_MAX", h.option.RetryWaitMax)
	env = setEnvIntValue(env, "SAKURA_RETRY_WAIT_MIN", h.option.RetryWaitMin)
	env = setEnvIntValue(env, "SAKURA_API_REQUEST_TIMEOUT", h.option.HTTPRequestTimeout)
	env = setEnvIntValue(env, "SAKURA_RATE_LIMIT", h.option.HTTPRequestRateLimit)

	endpoints, err := h.legacyClient.EndpointConfig()
	if err != nil {
		return nil, err
	}
	env = setEnvValue(env, "SAKURA_ZONE", endpoints.Zone)
	if len(endpoints.Zones) > 0 {
		env = setEnvValue(env, "SAKURA_ZONES", strings.Join(endpoints.Zones, ","))
	}
	for service, endpoint := range endpoints.Endpoints {
		key := strings.ToUpper(strings.NewReplacer("-", "_", ".", "_").Replace(service))
		env = setEnvValue(env, "SAKURA_ENDPOINTS_"+key, endpoint)
	}

	var base sdksaclient.Client
	if err := base.SetEnviron(env); err != nil {
		return nil, err
	}
	client, err := base.DupWith(sdksaclient.WithUserAgent(UserAgent))
	if err != nil {
		return nil, err
	}
	if err := client.Populate(); err != nil {
		return nil, err
	}
	return client, nil
}

func setEnvIntValue(env []string, key string, value int) []string {
	if value == 0 {
		return env
	}
	return setEnvValue(env, key, strconv.Itoa(value))
}

func setEnvValue(env []string, key, value string) []string {
	if value == "" {
		return env
	}
	prefix := key + "="
	result := env[:0]
	for _, item := range env {
		if !strings.HasPrefix(item, prefix) {
			result = append(result, item)
		}
	}
	return append(result, prefix+value)
}
