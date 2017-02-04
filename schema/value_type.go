package schema

// ValueType is an enum of the type that can be represented by a schema.
type ValueType int //go:generate stringer -type=ValueType :: manual

const (
	TypeInvalid ValueType = iota
	TypeBool
	TypeInt
	TypeInt64
	TypeFloat
	TypeString
	TypeIntList
	TypeStringList
)

// IsSliceType return true if type is TypeIntList or TypeStringList
func (v ValueType) IsSliceType() bool {
	switch v {
	case TypeIntList, TypeStringList:
		return true
	default:
		return false
	}
}
