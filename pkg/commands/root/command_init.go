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

//go:build !js

package root

import (
	"flag"
	"log"
	"os"
	"slices"

	"github.com/sacloud/usacloud/pkg/config"
)

func init() {
	Command.Flags().SortFlags = false
	Command.PersistentFlags().SortFlags = false

	if err := config.TheClient.SetEnviron(slices.Clone(os.Environ())); err != nil {
		log.Printf("Failed to load environment variables: %s", err)
	}
	config.InitConfig(Command.PersistentFlags())

	// This AddGoFlagSet() silently ignores duplicated flags;
	// They need extra touches.  Done in config.LoadConfigValue().
	Command.PersistentFlags().AddGoFlagSet(config.TheClient.FlagSet(flag.ContinueOnError))
}
