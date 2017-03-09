package schema

import (
	"github.com/sacloud/usacloud/output"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommand_Validate(t *testing.T) {
	var errs []error

	// empty
	c := &Command{}
	errs = c.Validate()
	assert.True(t, len(errs) > 0)

	// minimum
	c = &Command{
		Type: CommandCustom,
	}
	errs = c.Validate()
	assert.True(t, len(errs) == 0)

}

func TestCommand_ParamsHandlerType(t *testing.T) {
	var errs []error

	// parameters of valid only when CommandType is CommandList
	params := map[string]*Schema{
		"p1": {
			Type:        TypeStringList,
			HandlerType: HandlerSort, // valid only when Type is CommandList
		},
		"p2": {
			Type:        TypeStringList,
			HandlerType: HandlerAndParams, // valid only when Type is CommandList
		},
		"p3": {
			Type:        TypeStringList,
			HandlerType: HandlerOrParams, // valid only when Type is CommandList
		},
	}

	// command type isnot CommandList
	invalidCommand := &Command{
		Type:   CommandCreate,
		Params: params,
	}
	errs = invalidCommand.Validate()
	assert.True(t, len(errs) > 0)

	// command type is CommandList
	validCommand := &Command{
		Type:                CommandList,
		Params:              params,
		ListResultFieldName: "foo",
		TableType:           output.TableSimple,
		TableColumnDefines:  []output.ColumnDef{{Name: "Test"}},
	}
	errs = validCommand.Validate()
	assert.True(t, len(errs) == 0)
}

func TestCommand_TableType_ColumnDef(t *testing.T) {
	var errs []error

	c := &Command{
		Type:                CommandList,
		ListResultFieldName: "test",
	}
	errs = c.Validate()
	assert.True(t, len(errs) > 0)

	// commandList required TableColumnDefines and TableSimple
	c.TableColumnDefines = []output.ColumnDef{
		{Name: "Test"},
	}
	c.TableType = output.TableSimple

	errs = c.Validate()
	assert.True(t, len(errs) == 0)

	c.TableType = output.TableDetail
	errs = c.Validate()
	assert.True(t, len(errs) > 0)

}

func TestCommand_Categories_Params(t *testing.T) {

	var errs []error

	c := &Command{
		Type: CommandCustom,
		ParamCategories: []Category{
			{
				Key:         "test-category",
				DisplayName: "Test",
				Order:       1,
			},
		},
		Params: map[string]*Schema{
			"test": {
				Type:        TypeInt,
				HandlerType: HandlerNoop,
				Category:    "test-category",
				Order:       1,
			},
		},
	}

	errs = c.Validate()
	assert.True(t, len(errs) == 0)

	// not exists category
	c.Params["test"].Category = "not-exists-category"
	errs = c.Validate()
	assert.True(t, len(errs) > 0)
	c.Params["test"].Category = "test-category"

	// not exists param
	c.ParamCategories = append(c.ParamCategories, *c.ParamCategory("test-category"))
	c.ParamCategories[1].Key = "have-not-same-category-params"
	errs = c.Validate()
	assert.True(t, len(errs) > 0)

}
