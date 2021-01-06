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

package core

var (
	ResourceCategoryConfig = Category{
		Key:         "config",
		DisplayName: "Configuration",
		Order:       10,
	}

	ResourceCategoryAuth = Category{
		Key:         "auth",
		DisplayName: "Authentication",
		Order:       20,
	}

	ResourceCategoryComputing = Category{
		Key:         "computing",
		DisplayName: "Computing",
		Order:       30,
	}

	ResourceCategoryStorage = Category{
		Key:         "storage",
		DisplayName: "Storage",
		Order:       40,
	}

	ResourceCategoryNetworking = Category{
		Key:         "networking",
		DisplayName: "Networking",
		Order:       50,
	}

	ResourceCategoryNetworkingSub = Category{
		Key:         "networking-sub",
		DisplayName: "Networking(SubResources)",
		Order:       55,
	}

	ResourceCategoryAppliance = Category{
		Key:         "appliance",
		DisplayName: "Appliance",
		Order:       60,
	}

	ResourceCategorySecureMobile = Category{
		Key:         "securemobile",
		DisplayName: "SecureMobile",
		Order:       70,
	}
	ResourceCategoryCommonServiceItem = Category{
		Key:         "commonserviceitem",
		DisplayName: "Common service items",
		Order:       80,
	}

	ResourceCategoryCommonItem = Category{
		Key:         "commonitem",
		DisplayName: "Common items",
		Order:       90,
	}

	ResourceCategoryBilling = Category{
		Key:         "billing",
		DisplayName: "Billing",
		Order:       100,
	}

	ResourceCategoryLab = Category{
		Key:         "lab",
		DisplayName: "Lab",
		Order:       110,
	}

	ResourceCategoryWebAccel = Category{
		Key:         "webaccel",
		DisplayName: "WebAccelerator",
		Order:       120,
	}

	ResourceCategoryMisc = Category{
		Key:         "misc",
		DisplayName: "Other services",
		Order:       130,
	}

	ResourceCategoryZone = Category{
		Key:         "zone",
		DisplayName: "Region/Zone information",
		Order:       140,
	}

	ResourceCategoryInformation = Category{
		Key:         "information",
		DisplayName: "Service/Product information",
		Order:       150,
	}

	ResourceCategoryOther = Category{
		Key:         "other",
		DisplayName: "Other commands",
		Order:       160,
	}

	ResourceCategories = []Category{
		ResourceCategoryConfig,
		ResourceCategoryAuth,
		ResourceCategoryComputing,
		ResourceCategoryStorage,
		ResourceCategoryNetworking,
		ResourceCategoryNetworkingSub,
		ResourceCategoryAppliance,
		ResourceCategorySecureMobile,
		ResourceCategoryCommonServiceItem,
		ResourceCategoryCommonItem,
		ResourceCategoryBilling,
		ResourceCategoryLab,
		ResourceCategoryWebAccel,
		ResourceCategoryMisc,
		ResourceCategoryZone,
		ResourceCategoryInformation,
		ResourceCategoryOther,
	}
)
