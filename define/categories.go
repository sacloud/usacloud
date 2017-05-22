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
