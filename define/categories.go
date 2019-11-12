// Copyright 2017-2019 The Usacloud Authors
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

package define

import "github.com/sacloud/usacloud/schema"

var CategoryConfig = schema.Category{
	Key:         "config",
	DisplayName: "Configuration",
	Order:       3,
}
var CategoryAuth = schema.Category{
	Key:         "auth",
	DisplayName: "Authentication",
	Order:       5,
}
var CategoryComputing = schema.Category{
	Key:         "computing",
	DisplayName: "Computing",
	Order:       10,
}

var CategoryStorage = schema.Category{
	Key:         "storage",
	DisplayName: "Storage",
	Order:       20,
}

var CategoryNetworking = schema.Category{
	Key:         "networking",
	DisplayName: "Networking",
	Order:       30,
}

var CategoryAppliance = schema.Category{
	Key:         "appliance",
	DisplayName: "Appliance",
	Order:       40,
}
var CategoryCommonServiceItem = schema.Category{
	Key:         "commonserviceitem",
	DisplayName: "Common service items",
	Order:       50,
}

var CategoryCommonItem = schema.Category{
	Key:         "commonitem",
	DisplayName: "Common items",
	Order:       60,
}

var CategoryBilling = schema.Category{
	Key:         "billing",
	DisplayName: "Billing",
	Order:       70,
}

var CategoryCoupon = schema.Category{
	Key:         "coupon",
	DisplayName: "Coupon",
	Order:       75,
}

var CategoryOther = schema.Category{
	Key:         "saas",
	DisplayName: "Other services",
	Order:       80,
}

var CategoryInformation = schema.Category{
	Key:         "information",
	DisplayName: "Service/Product informations",
	Order:       90,
}

var CategorySummary = schema.Category{
	Key:         "summary",
	DisplayName: "Summary",
	Order:       100,
}
