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

package category

import (
	"github.com/sacloud/usacloud/pkg/category"
)

var (
	ResourceCategoryConfig = category.Category{
		Key:         "config",
		DisplayName: "Configuration",
		Order:       10,
	}

	ResourceCategoryAuth = category.Category{
		Key:         "auth",
		DisplayName: "Authentication",
		Order:       20,
	}

	ResourceCategoryComputing = category.Category{
		Key:         "computing",
		DisplayName: "Computing",
		Order:       30,
	}

	ResourceCategoryStorage = category.Category{
		Key:         "storage",
		DisplayName: "Storage",
		Order:       40,
	}

	ResourceCategoryNetworking = category.Category{
		Key:         "networking",
		DisplayName: "Networking",
		Order:       50,
	}

	ResourceCategoryNetworkingSub = category.Category{
		Key:         "networking-sub",
		DisplayName: "Networking(SubResources)",
		Order:       55,
	}

	ResourceCategoryAppliance = category.Category{
		Key:         "appliance",
		DisplayName: "Appliance",
		Order:       60,
	}

	ResourceCategorySecureMobile = category.Category{
		Key:         "securemobile",
		DisplayName: "SecureMobile",
		Order:       70,
	}
	ResourceCategoryCommonServiceItem = category.Category{
		Key:         "commonserviceitem",
		DisplayName: "Common service items",
		Order:       80,
	}

	ResourceCategoryCommonItem = category.Category{
		Key:         "commonitem",
		DisplayName: "Common items",
		Order:       90,
	}

	ResourceCategoryBilling = category.Category{
		Key:         "billing",
		DisplayName: "Billing",
		Order:       100,
	}

	ResourceCategoryLab = category.Category{
		Key:         "lab",
		DisplayName: "Lab",
		Order:       110,
	}

	ResourceCategoryWebAccel = category.Category{
		Key:         "webaccel",
		DisplayName: "WebAccelerator",
		Order:       120,
	}

	ResourceCategoryMisc = category.Category{
		Key:         "misc",
		DisplayName: "Other services",
		Order:       130,
	}

	ResourceCategoryZone = category.Category{
		Key:         "zone",
		DisplayName: "Region/Zone information",
		Order:       140,
	}

	ResourceCategoryInformation = category.Category{
		Key:         "information",
		DisplayName: "Service/Product information",
		Order:       150,
	}

	ResourceCategoryOther = category.Category{
		Key:         "other",
		DisplayName: "Other commands",
		Order:       160,
	}

	ResourceCategories = []category.Category{
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
