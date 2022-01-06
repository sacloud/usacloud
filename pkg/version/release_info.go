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

package version

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// GitHubReleaseInfoURL リリース情報取得用のURL - ビルド時に-ldflagsで差し替えることがあるためconstではなくvarにしておく
var GitHubReleaseInfoURL = "https://api.github.com/repos/sacloud/usacloud/releases/latest"

// ReleaseInfo リリース情報のキャッシュ
type ReleaseInfo struct {
	Release   GitHubRelease
	FetchedAt time.Time
}

func (r *ReleaseInfo) Expired() bool {
	return r.FetchedAt.Add(cacheExpireDuration).Before(time.Now())
}

func fetchReleaseInfo(ctx context.Context) (*ReleaseInfo, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, GitHubReleaseInfoURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var githubRelease GitHubRelease
	if err := json.Unmarshal(data, &githubRelease); err != nil {
		return nil, err
	}
	return &ReleaseInfo{
		Release:   githubRelease,
		FetchedAt: time.Now(),
	}, nil
}
