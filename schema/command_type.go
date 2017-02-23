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
	CommandManipulate // power-on/power-off
	CommandCustom
)

func (c CommandType) IsRequiredIDType() bool {
	switch c {
	case CommandRead, CommandUpdate, CommandDelete, CommandManipulate:
		return true
	default:
		return false
	}
}
