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

package category

var CommandCategories = []Category{
	{
		Key:         "basic",
		DisplayName: "Basic Commands",
		Order:       100,
	},
	{
		Key:         "connect",
		DisplayName: "Connect Commands",
		Order:       150,
	},
	{
		Key:         "operation",
		DisplayName: "Operation Commands",
		Order:       200,
	},
	{
		Key:         "subnet",
		DisplayName: "Subnet Operation Commands",
		Order:       210,
	},
	{
		Key:         "ipv6",
		DisplayName: "IPv6 Operation Commands",
		Order:       211,
	},
	{
		Key:         "certificate",
		DisplayName: "Certificate Management Commands",
		Order:       220,
	},
	{
		Key:         "cache",
		DisplayName: "Cache Management Commands",
		Order:       221,
	},
	{
		Key:         "power",
		DisplayName: "Power Management Commands",
		Order:       250,
	},
	{
		Key:         "backup",
		DisplayName: "Backup Commands",
		Order:       300,
	},
	{
		Key:         "monitor",
		DisplayName: "Monitoring Commands",
		Order:       500,
	},
	{
		Key:         "other",
		DisplayName: "Other Commands",
		Order:       10000,
	},
}
