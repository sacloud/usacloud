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

package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/sacloud/libsacloud/sacloud"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/profile"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/helper/migration"
	"github.com/sacloud/usacloud/helper/printer"
)

func checkConfigVersion() error {
	return migration.CheckConfigVersion()
}

type FlagHandler interface {
	IsSet(name string) bool
	Set(name, value string) error
	String(name string) string
	StringSlice(name string) []string
}

func applyConfigFromFile(c FlagHandler) error {
	profileKey := "profile"
	profileName := c.String(profileKey)
	if profileName == "" {
		n, err := profile.GetCurrentName()
		if err != nil {
			return fmt.Errorf("Failed to load current profile: %s", err)
		}
		profileName = n
	}

	// load config file
	v, err := profile.LoadConfigFile(profileName)
	if err != nil {
		return err
	}

	if !c.IsSet("token") && v.AccessToken != "" {
		c.Set("token", v.AccessToken)
		command.GlobalOption.AccessToken = v.AccessToken
	}
	if !c.IsSet("secret") && v.AccessTokenSecret != "" {
		c.Set("secret", v.AccessTokenSecret)
		command.GlobalOption.AccessTokenSecret = v.AccessTokenSecret
	}
	if !c.IsSet("zone") && v.Zone != "" {
		c.Set("zone", v.Zone)
		command.GlobalOption.Zone = v.Zone
	}
	if !c.IsSet("default-output-type") && v.DefaultOutputType != "" {
		c.Set("default-output-type", v.DefaultOutputType)
		command.GlobalOption.DefaultOutputType = v.DefaultOutputType
	}

	if !c.IsSet("accept-language") && v.AcceptLanguage != "" {
		c.Set("accept-language", v.AcceptLanguage)
		command.GlobalOption.AcceptLanguage = v.AcceptLanguage
	}
	if !c.IsSet("retry-max") && v.RetryMax > 0 {
		c.Set("retry-max", fmt.Sprintf("%d", v.RetryMax))
		command.GlobalOption.RetryMax = v.RetryMax
	}
	if !c.IsSet("retry-interval") && v.RetryIntervalSec > 0 {
		c.Set("retry-interval", fmt.Sprintf("%d", v.RetryIntervalSec))
		command.GlobalOption.RetryIntervalSec = v.RetryIntervalSec
	}
	if !c.IsSet("api-request-timeout") && v.APIRequestTimeout > 0 {
		c.Set("api-request-timeout", fmt.Sprintf("%d", v.APIRequestTimeout))
		command.GlobalOption.APIRequestTimeout = v.APIRequestTimeout
	}
	if !c.IsSet("no-color") && v.NoColor {
		c.Set("no-color", "true")
		command.GlobalOption.NoColor = v.NoColor
	}

	// for string-slice
	zones := []string{}
	if c.IsSet("zones") {
		zones = c.StringSlice("zones")
	} else if len(v.Zones) > 0 {
		zones = v.Zones
	} else {
		if z, ok := os.LookupEnv("USACLOUD_ZONES"); ok {
			zones = strings.Split(z, ",")
		}
	}
	command.GlobalOption.Zones = zones

	if !c.IsSet("api-root-url") && v.APIRootURL != "" {
		c.Set("api-root-url", v.APIRootURL)
		command.GlobalOption.APIRootURL = v.APIRootURL
	}

	if len(command.GlobalOption.Zones) > 0 {
		define.AllowZones = command.GlobalOption.Zones
	}
	if command.GlobalOption.APIRootURL != "" {
		api.SakuraCloudAPIRoot = command.GlobalOption.APIRootURL
	}

	return nil
}

func toSakuraID(id string) (sacloud.ID, bool) {
	sid := sacloud.StringID(id)
	return sid, !sid.IsEmpty()
}

func toSakuraIDs(ids []int64) []sacloud.ID {
	var res []sacloud.ID
	for _, v := range ids {
		res = append(res, sacloud.ID(v))
	}
	return res
}

func hasTags(target interface{}, tags []string) bool {
	type tagHandler interface {
		HasTag(target string) bool
	}

	tagHolder, ok := target.(tagHandler)
	if !ok {
		return false
	}

	// 完全一致 + AND条件
	res := true
	for _, p := range tags {
		if !tagHolder.HasTag(p) {
			res = false
			break
		}
	}
	return res
}

func isTerminal() bool {
	is := func(fd uintptr) bool {
		return isatty.IsTerminal(fd) || isatty.IsCygwinTerminal(fd)
	}
	return is(os.Stdin.Fd()) && is(os.Stdout.Fd())
}

func printWarning(warn string) {
	if warn == "" {
		return
	}
	printer.Fprintf(command.GlobalOption.Err, color.New(color.FgYellow), "[WARN] %s\n", warn)
}
