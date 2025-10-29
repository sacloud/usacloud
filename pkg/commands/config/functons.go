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
	"os"
	"strings"

	saht "github.com/sacloud/saclient-go"
	"github.com/spf13/cobra"
)

func profileCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	op := saht.NewProfileOp(os.Environ()) // no context, create a new op

	names, err := op.List()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	if toComplete == "" {
		return names, cobra.ShellCompDirectiveNoFileComp
	}
	var results []string
	for _, n := range names {
		if strings.HasPrefix(n, toComplete) {
			results = append(results, n)
		}
	}
	return results, cobra.ShellCompDirectiveNoFileComp
}
