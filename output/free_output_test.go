package output

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dummyOption struct{}

func (o dummyOption) GetOutputType() string { return "" }
func (o dummyOption) GetColumn() []string   { return []string{} }
func (o dummyOption) GetFormat() string     { return "test ID:{{.ID}}" }
func (o dummyOption) GetFormatFile() string { return "" }
func (o dummyOption) GetQuiet() bool        { return false }
func (o dummyOption) GetQuery() string      { return "" }

func TestFreeOutput_Print(t *testing.T) {
	buf := bytes.NewBufferString("")
	o := NewFreeOutput(buf, os.Stderr, dummyOption{})

	type dummy struct {
		ID int64
	}

	values := []interface{}{
		&dummy{ID: 1},
		&dummy{ID: 2},
	}

	err := o.Print(values...)

	assert.NoError(t, err)
	assert.Equal(t, testFreeOutputText, buf.String())

}

var testFreeOutputText = `test ID:1
test ID:2
`
