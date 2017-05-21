package schema

import "math"

type Category struct {
	Key         string
	DisplayName string
	Order       int
}

var DefaultResourceCategory = &Category{
	Key:         "default",
	DisplayName: "",
	Order:       math.MaxInt32,
}

var DefaultCommandCategory = &Category{
	Key:         "default",
	DisplayName: "",
	Order:       math.MaxInt32,
}

var FilterParamCategory = &Category{
	Key:         "filter",
	DisplayName: "Filter options",
	Order:       math.MaxInt32 - 60,
}
var LimitOffsetParamCategory = &Category{
	Key:         "limit-offset",
	DisplayName: "Limit/Offset options",
	Order:       math.MaxInt32 - 50,
}
var SortParamCategory = &Category{
	Key:         "sort",
	DisplayName: "Sort options",
	Order:       math.MaxInt32 - 40,
}

var CommonParamCategory = &Category{
	Key:         "common",
	DisplayName: "Common options",
	Order:       math.MaxInt32 - 30,
}

var InputParamCategory = &Category{
	Key:         "Input",
	DisplayName: "Input options",
	Order:       math.MaxInt32 - 20,
}

var OutputParamCategory = &Category{
	Key:         "output",
	DisplayName: "Output options",
	Order:       math.MaxInt32 - 10,
}

var DefaultParamCategory = &Category{
	Key:         "default",
	DisplayName: "Other options",
	Order:       math.MaxInt32,
}
