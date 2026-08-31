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
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"testing"

	legacysaclient "github.com/sacloud/saclient-go"
	sdksaclient "github.com/sacloud/sacloud-sdk-go/common/saclient"
	"github.com/sacloud/usacloud/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestSDKClientHolder(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	privateKeyPath := filepath.Join(t.TempDir(), "private-key.pem")
	require.NoError(t, os.WriteFile(privateKeyPath, pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}), 0600))

	var legacyClient legacysaclient.Client
	require.NoError(t, legacyClient.SetEnviron([]string{
		"SAKURACLOUD_PROFILE_DIR=" + t.TempDir(),
		"SAKURA_ACCESS_TOKEN=legacy-token",
		"SAKURA_ACCESS_TOKEN_SECRET=legacy-secret",
		"SAKURA_ENDPOINTS_EVENTBUS=http://127.0.0.1:18085/",
	}))
	require.NoError(t, legacyClient.Populate())

	option := &config.Config{
		TokenEndpoint:         "https://example.com/oauth2/token",
		ServicePrincipalID:    "service-principal-id",     // #nosec G101 -- non-secret test fixture
		ServicePrincipalKeyID: "service-principal-key-id", // #nosec G101 -- non-secret test fixture
		PrivateKeyPEMPath:     privateKeyPath,
	}
	option.HTTPRequestRateLimit = 17
	holder := newSDKClientHolder(option, &legacyClient)
	client, err := holder.Client()
	require.NoError(t, err)
	sameClient, err := holder.Client()
	require.NoError(t, err)
	require.Same(t, client, sameClient)

	endpoints, err := client.EndpointConfig()
	require.NoError(t, err)
	require.Equal(t, "http://127.0.0.1:18085/", endpoints.Endpoints["eventbus"])

	settings := client.(*sdksaclient.Client).JSON()
	require.Equal(t, "https://example.com/oauth2/token", settings["TokenEndpoint"])
	require.Equal(t, "service-principal-id", settings["ServicePrincipalID"])
	require.Equal(t, "service-principal-key-id", settings["ServicePrincipalKeyID"])
	require.Equal(t, privateKeyPath, settings["PrivateKeyPEMPath"])
	require.Equal(t, int64(17), settings["APIRequestRateLimit"])
}

func TestSDKClientHolderUsesResolvedProfileCredentials(t *testing.T) {
	profileDir := t.TempDir()
	op, err := sdksaclient.NewProfileOp([]string{"SAKURA_PROFILE_DIR=" + profileDir})
	require.NoError(t, err)
	require.NoError(t, op.Create(&sdksaclient.Profile{
		Name: "staging",
		Attributes: map[string]any{
			"AccessToken":       "staging-token",
			"AccessTokenSecret": "staging-secret",
		},
	}))
	require.NoError(t, op.Create(&sdksaclient.Profile{
		Name: "production",
		Attributes: map[string]any{
			"ServicePrincipalID":    "production-service-principal",
			"ServicePrincipalKeyID": "production-key",
			"PrivateKeyPEMPath":     "/production/private-key.pem",
		},
	}))
	require.NoError(t, op.SetCurrentName("production"))

	t.Setenv("SAKURA_PROFILE_DIR", profileDir)
	t.Setenv("SAKURA_PROFILE", "production")

	var legacyClient legacysaclient.Client
	require.NoError(t, legacyClient.SetEnviron([]string{
		"SAKURA_PROFILE_DIR=" + profileDir,
		"SAKURA_PROFILE=staging",
		"SAKURA_ACCESS_TOKEN=staging-token",
		"SAKURA_ACCESS_TOKEN_SECRET=staging-secret",
	}))
	require.NoError(t, legacyClient.Populate())

	holder := newSDKClientHolder(&config.Config{Profile: "staging"}, &legacyClient)
	holder.option.AccessToken = "staging-token"
	holder.option.AccessTokenSecret = "staging-secret"

	client, err := holder.Client()
	require.NoError(t, err)

	settings := client.(*sdksaclient.Client).JSON()
	require.Equal(t, "staging-token", settings["AccessToken"])
	require.Equal(t, "staging-secret", settings["AccessTokenSecret"])
	require.Equal(t, "staging", settings["ProfileName"])
	require.NotContains(t, settings, "ServicePrincipalID")
	require.NotContains(t, settings, "ServicePrincipalKeyID")
	require.NotContains(t, settings, "PrivateKeyPEMPath")
}

func TestSetEnvValueReplacesExistingValue(t *testing.T) {
	actual := setEnvValue([]string{"A=old", "B=value", "A=older"}, "A", "new")
	require.Equal(t, []string{"B=value", "A=new"}, actual)
}
