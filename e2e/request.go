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

package e2e

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// HttpGet 指定のURLにGETでリクエストを行い、ステータスコード200以外の場合はエラーを返す
func HttpGet(url string) error {
	res, err := http.Get(url) // nolint
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("got unexpected status code: %d", res.StatusCode)
	}
	return nil
}

// HttpRequestUntilSuccess ステータス200が返ってくるまで定期的に指定のURLにGETでリクエストを行う
func HttpRequestUntilSuccess(url string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	doneCh := make(chan error)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				doneCh <- ctx.Err()
				return
			case <-ticker.C:
				if err := HttpGet(url); err != nil {
					continue
				}
				doneCh <- nil
				return
			}
		}
	}()

	return <-doneCh
}
