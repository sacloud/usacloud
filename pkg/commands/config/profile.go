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

package config

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"slices"

	"github.com/sacloud/saclient-go"
	"github.com/sacloud/usacloud/pkg/config"
)

// v1.20.0 より前(api-client-go/profile)では、default プロファイルは
// 設定ファイルが存在しなくても常に存在するものとして扱われていた。
// saclient.ProfileOp は設定ファイルの実在を前提とするため、ここで互換性を保つ。

// currentOrDefault current プロファイル名を返す
//
// current ファイルが存在しない、または空の場合は default を返す
func currentOrDefault(op saclient.ProfileAPI) (string, error) {
	name, err := op.GetCurrentName()
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return config.DefaultProfileName, nil
	case err != nil:
		return "", err
	case name == "":
		return config.DefaultProfileName, nil
	default:
		return name, nil
	}
}

// listWithDefault プロファイル名の一覧を返す
//
// default は設定ファイルの有無にかかわらず常に先頭に含める
func listWithDefault(op saclient.ProfileAPI) ([]string, error) {
	names, err := op.List()
	if err != nil {
		return nil, err
	}
	names = slices.DeleteFunc(names, func(n string) bool { return n == config.DefaultProfileName })
	return append([]string{config.DefaultProfileName}, names...), nil
}

// useProfile 指定のプロファイルを current に設定する
func useProfile(op saclient.ProfileAPI, name string) error {
	if name == config.DefaultProfileName {
		return useDefault(op)
	}
	return op.SetCurrentName(name)
}

// useDefault default プロファイルを current に設定する
//
// default の設定ファイルが存在する場合は current に書き込む。
// 存在しない場合は current ファイルを削除し、未設定(=default)の状態に戻す。
// saclient は current が指すプロファイルが存在しないとエラーになるため、
// 設定ファイルの無い default を current に書き込んではならない。
func useDefault(op saclient.ProfileAPI) error {
	names, err := op.List()
	if err != nil {
		return err
	}
	if slices.Contains(names, config.DefaultProfileName) {
		return op.SetCurrentName(config.DefaultProfileName)
	}
	err = os.Remove(filepath.Join(op.Dir(), "current"))
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	return err
}
