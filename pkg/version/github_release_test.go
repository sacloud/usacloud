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

package version

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGitHubRelease_GreaterThanCurrent(t *testing.T) {
	cases := []struct {
		current string
		github  string
		expect  bool
	}{
		{
			current: "v1.0.0",
			github:  "v1.0.1",
			expect:  true,
		},
		{
			current: "v1.0.0",
			github:  "v1.1.0",
			expect:  true,
		},
		{
			current: "v1.0.0",
			github:  "v2.0.0",
			expect:  true,
		},
		{
			current: "v1.0.0-dev",
			github:  "v1.0.1",
			expect:  true,
		},
		{
			current: "v1.0.0-beta.1",
			github:  "v1.0.1",
			expect:  true,
		},
		{
			current: "v1.0.0",
			github:  "v1.0.0",
			expect:  false,
		},
		{
			current: "v1.0.0",
			github:  "v0.0.1",
			expect:  false,
		},
		{
			current: "v1.0.0-dev",
			github:  "v0.0.1",
			expect:  false,
		},
	}
	for _, tc := range cases {
		Version = tc.current
		ghv := &GitHubRelease{TagName: tc.github}

		got, err := ghv.GreaterThanCurrent()
		require.NoError(t, err)
		require.Equal(t, tc.expect, got)
	}
}

func TestGitHubRelease_UnmarshalJSON(t *testing.T) {
	var releaseInfo GitHubRelease
	if err := json.Unmarshal([]byte(githubReleaseAPIResponse), &releaseInfo); err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "v1.0.0", releaseInfo.TagName)
	require.Equal(t, "https://github.com/octocat/Hello-World/releases/v1.0.0", releaseInfo.URL)
}

// from: https://docs.github.com/en/free-pro-team@latest/rest/reference/repos#get-the-latest-release
var githubReleaseAPIResponse = `{
  "url": "https://api.github.com/repos/octocat/Hello-World/releases/1",
  "html_url": "https://github.com/octocat/Hello-World/releases/v1.0.0",
  "assets_url": "https://api.github.com/repos/octocat/Hello-World/releases/1/assets",
  "upload_url": "https://uploads.github.com/repos/octocat/Hello-World/releases/1/assets{?name,label}",
  "tarball_url": "https://api.github.com/repos/octocat/Hello-World/tarball/v1.0.0",
  "zipball_url": "https://api.github.com/repos/octocat/Hello-World/zipball/v1.0.0",
  "id": 1,
  "node_id": "MDc6UmVsZWFzZTE=",
  "tag_name": "v1.0.0",
  "target_commitish": "master",
  "name": "v1.0.0",
  "body": "Description of the release",
  "draft": false,
  "prerelease": false,
  "created_at": "2013-02-27T19:35:32Z",
  "published_at": "2013-02-27T19:35:32Z",
  "assets": [
    {
      "url": "https://api.github.com/repos/octocat/Hello-World/releases/assets/1",
      "browser_download_url": "https://github.com/octocat/Hello-World/releases/download/v1.0.0/example.zip",
      "id": 1,
      "node_id": "MDEyOlJlbGVhc2VBc3NldDE=",
      "name": "example.zip",
      "label": "short description",
      "state": "uploaded",
      "content_type": "application/zip",
      "size": 1024,
      "download_count": 42,
      "created_at": "2013-02-27T19:35:32Z",
      "updated_at": "2013-02-27T19:35:32Z",
      "uploader": {
        "login": "octocat",
        "id": 1,
        "node_id": "MDQ6VXNlcjE=",
        "avatar_url": "https://github.com/images/error/octocat_happy.gif",
        "gravatar_id": "",
        "url": "https://api.github.com/users/octocat",
        "html_url": "https://github.com/octocat",
        "followers_url": "https://api.github.com/users/octocat/followers",
        "following_url": "https://api.github.com/users/octocat/following{/other_user}",
        "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
        "organizations_url": "https://api.github.com/users/octocat/orgs",
        "repos_url": "https://api.github.com/users/octocat/repos",
        "events_url": "https://api.github.com/users/octocat/events{/privacy}",
        "received_events_url": "https://api.github.com/users/octocat/received_events",
        "type": "User",
        "site_admin": false
      }
    }
  ]
}`
