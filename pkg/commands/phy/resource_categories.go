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

package phy

import (
	"github.com/sacloud/usacloud/pkg/category"
)

var (
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

	ResourceCategoryAppliance = category.Category{
		Key:         "appliance",
		DisplayName: "Appliance",
		Order:       60,
	}

	ResourceCategoryMisc = category.Category{
		Key:         "misc",
		DisplayName: "Other services",
		Order:       130,
	}

	ResourceCategoryOther = category.Category{
		Key:         "other",
		DisplayName: "Other commands",
		Order:       160,
	}

	ResourceCategories = []category.Category{
		ResourceCategoryComputing,
		ResourceCategoryStorage,
		ResourceCategoryNetworking,
		ResourceCategoryAppliance,
		ResourceCategoryMisc,
		ResourceCategoryOther,
	}
)
