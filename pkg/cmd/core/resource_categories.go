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

package core

var (
	ResourceCategoryConfig = Category{
		Key:         "config",
		DisplayName: "Configuration",
		Order:       3,
	}
	ResourceCategoryAuth = Category{
		Key:         "auth",
		DisplayName: "Authentication",
		Order:       5,
	}
	ResourceCategoryComputing = Category{
		Key:         "computing",
		DisplayName: "Computing",
		Order:       10,
	}

	ResourceCategoryStorage = Category{
		Key:         "storage",
		DisplayName: "Storage",
		Order:       20,
	}

	ResourceCategoryNetworking = Category{
		Key:         "networking",
		DisplayName: "Networking",
		Order:       30,
	}

	ResourceCategoryAppliance = Category{
		Key:         "appliance",
		DisplayName: "Appliance",
		Order:       40,
	}

	ResourceCategorySecureMobile = Category{
		Key:         "securemobile",
		DisplayName: "SecureMobile",
		Order:       45,
	}
	ResourceCategoryCommonServiceItem = Category{
		Key:         "commonserviceitem",
		DisplayName: "Common service items",
		Order:       50,
	}

	ResourceCategoryCommonItem = Category{
		Key:         "commonitem",
		DisplayName: "Common items",
		Order:       60,
	}

	ResourceCategoryBilling = Category{
		Key:         "billing",
		DisplayName: "Billing",
		Order:       70,
	}

	ResourceCategoryCoupon = Category{
		Key:         "coupon",
		DisplayName: "Coupon",
		Order:       75,
	}

	ResourceCategoryLab = Category{
		Key:         "lab",
		DisplayName: "Lab",
		Order:       78,
	}

	ResourceCategoryOther = Category{
		Key:         "misc",
		DisplayName: "Other services",
		Order:       80,
	}

	ResourceCategoryInformation = Category{
		Key:         "information",
		DisplayName: "Service/Product informations",
		Order:       90,
	}

	ResourceCategorySummary = Category{
		Key:         "summary",
		DisplayName: "Summary",
		Order:       100,
	}

	ResourceCategories = []Category{
		ResourceCategoryConfig,
		ResourceCategoryAuth,
		ResourceCategoryComputing,
		ResourceCategoryStorage,
		ResourceCategoryNetworking,
		ResourceCategoryAppliance,
		ResourceCategoryCommonServiceItem,
		ResourceCategoryCommonItem,
		ResourceCategoryBilling,
		ResourceCategoryCoupon,
		ResourceCategoryOther,
		ResourceCategoryInformation,
		ResourceCategorySummary,
	}
)
