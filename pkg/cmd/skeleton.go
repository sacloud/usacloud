// Copyright 2017-2020 The Usacloud Authors
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

package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/sacloud/usacloud/pkg/cli"
)

type skeletonHandler interface {
	SetGenerateSkeleton(bool)
	FillValueToSkeleton()
}

func generateSkeleton(ctx cli.Context, params interface{}) error {
	v, ok := params.(skeletonHandler)
	if !ok {
		return nil
	}

	v.SetGenerateSkeleton(false)
	v.FillValueToSkeleton()

	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return fmt.Errorf("marshaling to JSON is failed: %s", err)
	}
	_, err = fmt.Fprintln(ctx.IO().Out(), string(data))
	return err
}
