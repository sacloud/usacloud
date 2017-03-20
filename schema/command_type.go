package schema

// CommandType is an enum of the type that can be represented by a schema.
type CommandType int //go:generate stringer -type=CommandType :: manual

const (
	CommandInvalid CommandType = iota
	CommandList
	CommandCreate
	CommandRead
	CommandUpdate
	CommandDelete
	CommandManipulateMulti
	CommandManipulateSingle
	CommandManipulateIDOnly
	CommandCustom
)

func (c CommandType) IsRequiredIDType() bool {
	switch c {
	case CommandRead, CommandUpdate, CommandDelete, CommandManipulateMulti, CommandManipulateSingle, CommandManipulateIDOnly:
		return true
	default:
		return false
	}
}

func (c CommandType) IsNeedSingleIDType() bool {
	return c == CommandManipulateSingle
}

func (c CommandType) IsNeedIDOnlyType() bool {
	return c == CommandManipulateIDOnly
}

func (c CommandType) IsNeedConfirmType() bool {
	switch c {
	case CommandCreate, CommandUpdate, CommandDelete, CommandManipulateMulti, CommandManipulateSingle, CommandManipulateIDOnly, CommandCustom:
		return true
	default:
		return false
	}
}
