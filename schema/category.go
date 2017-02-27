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

var DefaultParamCategory = &Category{
	Key:         "default",
	DisplayName: "Other options",
	Order:       math.MaxInt32,
}
