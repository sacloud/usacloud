package schema

// HandlerType is an enum of the type that can be represented by a schema.
type HandlerType int //go:generate stringer -type=HandlerType :: manual

const (
	HandlerPathThrough HandlerType = iota
	HandlerPathThroughEach
	HandlerSort
	HandlerAndParams
	HandlerOrParams
	HandlerCustomFunc
)

// IsWhenListOnly return true when HandlerType is able to use with CommandList
func (h HandlerType) IsWhenListOnly() bool {
	switch h {
	case HandlerSort, HandlerAndParams, HandlerOrParams:
		return true
	default:
		return false
	}
}

// IsNeedSliceValue return true when HandlerType is need Slice ValueType
func (h HandlerType) IsNeedSliceValue() bool {
	switch h {
	case HandlerPathThroughEach, HandlerSort, HandlerAndParams, HandlerOrParams:
		return true
	default:
		return false
	}
}

// CanSetDestinationProp return true when HandlerType is able to have DestinationProp
func (h HandlerType) CanSetDestinationProp() bool {
	switch h {
	case HandlerPathThrough, HandlerPathThroughEach, HandlerAndParams, HandlerOrParams:
		return true
	default:
		return false
	}
}
