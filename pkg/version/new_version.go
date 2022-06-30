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

//go:build !wasm
// +build !wasm

package version

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/sacloud/api-client-go/profile"
)

func NewVersionReleased() (*ReleaseInfo, error) {
	configDir, err := profile.ConfigDir()
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return nil, err
	}
	var releaseInfo *ReleaseInfo
	cachePath := filepath.Join(configDir, ".release_info_cache")
	if _, err := os.Stat(cachePath); err == nil {
		data, err := os.ReadFile(cachePath)
		if err != nil {
			return nil, err
		}
		var cached ReleaseInfo
		if err := json.Unmarshal(data, &cached); err != nil {
			return nil, err
		}
		if !cached.Expired() {
			releaseInfo = &cached
		}
	}

	if releaseInfo == nil {
		// 数秒待ってもダメなら諦める
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		fetched, err := fetchReleaseInfo(ctx)
		if err != nil {
			if ctx.Err() != nil && ctx.Err() == context.DeadlineExceeded {
				// GitHub APIがタイムアウトしたら現在のバージョンでダミー値を作成してキャッシュする
				fetched = &ReleaseInfo{
					Release: GitHubRelease{
						TagName: Version,
						URL:     "",
					},
					FetchedAt: time.Now().Add(-55 * time.Minute), // ダミー値の有効期限は5分
				}
			} else {
				return nil, err
			}
		}
		// write cache file
		data, err := json.Marshal(fetched)
		if err != nil {
			return nil, err
		}
		if err := os.WriteFile(cachePath, data, 0600); err != nil {
			return nil, err
		}
		releaseInfo = fetched
	}

	return releaseInfo, nil
}
