package schema

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchema_ValidateMinimum(t *testing.T) {

	var errs []error
	// empty schema
	s := &Schema{}
	errs = s.Validate("empty")
	assert.True(t, len(errs) > 0)

	// valid minimum schema
	s = &Schema{
		Type: TypeBool,
	}
	errs = s.Validate("minimum")
	assert.True(t, len(errs) == 0)
}

func TestSchema_ValidateEnvVars(t *testing.T) {

	var errs []error
	// duplicate EnvVars
	s := &Schema{
		Type:    TypeBool,
		EnvVars: []string{"a", "a", "b"},
	}
	errs = s.Validate("duplicate-env-vars")
	assert.True(t, len(errs) > 0)

	// valid EnvVars
	s = &Schema{
		Type:    TypeBool,
		EnvVars: []string{"a", "b"},
	}
	errs = s.Validate("valid-env-vars")
	assert.True(t, len(errs) == 0)
}

func TestSchema_ValidateDefaultValue(t *testing.T) {

	var errs []error
	// DefaultValue and Type ref
	s := &Schema{
		Type:         TypeBool,
		DefaultValue: 1, // int
	}
	errs = s.Validate("invalid-default-value1")
	assert.True(t, len(errs) > 0)

	s = &Schema{
		Type:         TypeBool,
		DefaultValue: "hoge", // string
	}
	errs = s.Validate("invalid-default-value2")
	assert.True(t, len(errs) > 0)

	s = &Schema{
		Type:         TypeBool,
		DefaultValue: []string{}, // string slice
	}
	errs = s.Validate("invalid-default-value3")
	assert.True(t, len(errs) > 0)

	s = &Schema{
		Type:         TypeBool,
		DefaultValue: true,
	}
	errs = s.Validate("valid-default-value")
	assert.True(t, len(errs) == 0)
}

func TestSchema_ValidateConflictsWithAndRequired(t *testing.T) {

	var errs []error

	// ConflictsWith
	s := &Schema{
		Type:          TypeBool,
		Required:      true,
		ConflictsWith: []string{"hoge"},
	}
	errs = s.Validate("invalid-required-with-conflictwith")
	assert.True(t, len(errs) > 0)

	// valid conflicts with
	s = &Schema{
		Type:          TypeBool,
		Required:      false,
		ConflictsWith: []string{"hoge"},
	}
	errs = s.Validate("valid-required-with-conflictwith")
	assert.True(t, len(errs) == 0)
}

func TestSchema_ValidateMaxMinItems(t *testing.T) {

	var errs []error
	// MaxItems/MinItems
	s := &Schema{
		Type:     TypeBool,
		MaxItems: 1, // with non slice type
	}
	errs = s.Validate("invalid-max-item")
	assert.True(t, len(errs) > 0)

	s = &Schema{
		Type:     TypeBool,
		MinItems: 1, // with non slice type
	}
	errs = s.Validate("invalid-min-item")
	assert.True(t, len(errs) > 0)

}

func TestSchema_ValidateDestination(t *testing.T) {
	resultMap := map[HandlerType]bool{
		HandlerPathThrough:     true,
		HandlerPathThroughEach: true,
		HandlerSort:            false,
		HandlerOrParams:        true,
		HandlerAndParams:       true,
		HandlerCustomFunc:      false,
	}
	s := &Schema{
		Type:            TypeStringList,
		DestinationProp: "test",
	}
	for handlerType, result := range resultMap {
		s.HandlerType = handlerType
		errs := s.Validate("test")
		assert.Equal(t, len(errs) == 0, result)
	}

}

func TestSchema_NeedSliceValueHandlers(t *testing.T) {
	resultMap := map[ValueType]bool{
		TypeString:     false,
		TypeInt:        false,
		TypeInt64:      false,
		TypeBool:       false,
		TypeFloat:      false,
		TypeStringList: true,
		TypeIntList:    true,
	}

	needSliceValueHandlers := []HandlerType{
		HandlerPathThroughEach,
		HandlerSort,
		HandlerOrParams,
		HandlerAndParams,
	}

	s := &Schema{}

	for _, handlerType := range needSliceValueHandlers {
		for valueType, result := range resultMap {
			s.Type = valueType
			s.HandlerType = handlerType

			errs := s.Validate("test")
			assert.Equal(t, len(errs) == 0, result)
		}
	}
}

func TestSchema_NeedCustomHadler(t *testing.T) {
	resultMap := map[HandlerType]bool{
		HandlerPathThrough:     true,
		HandlerPathThroughEach: true,
		HandlerSort:            true,
		HandlerAndParams:       true,
		HandlerOrParams:        true,
		HandlerCustomFunc:      false,
	}

	s := &Schema{
		Type: TypeStringList,
	}

	for handlerType, result := range resultMap {
		s.HandlerType = handlerType

		errs := s.Validate("test")
		assert.Equal(t, len(errs) == 0, result)
	}
}
