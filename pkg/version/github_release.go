// Copyright 2017-2021 The Usacloud Authors
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
	"time"

	"github.com/hashicorp/go-version"
)

const cacheExpireDuration = time.Hour

// GitHubRelease GitHubのリリース情報
//
// refs: https://docs.github.com/en/free-pro-team@latest/rest/reference/repos#get-the-latest-release
type GitHubRelease struct {
	TagName string `json:"tag_name"`
	URL     string `json:"html_url"`
}

func (r *GitHubRelease) GreaterThanCurrent() (bool, error) {
	ghv, err := version.NewSemver(r.TagName)
	if err != nil {
		return false, err
	}

	current, err := version.NewSemver(Version)
	if err != nil {
		return false, err
	}

	return current.LessThan(ghv), nil
}
