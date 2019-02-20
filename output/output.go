package output

type Output interface {
	Print(...interface{}) error
}

type Formatter interface {
	GetIncludeFields() []string
	GetExcludeFields() []string
	GetColumnDefs() []ColumnDef
	GetTableType() TableType
	Option
}
type Option interface {
	GetOutputType() string
	GetColumn() []string
	GetFormat() string
	GetFormatFile() string
	GetQuiet() bool
	GetQuery() string
	GetQueryFile() string
}

type TableType int //go:generate stringer -type=OutputTableType :: manual
const (
	TableDetail TableType = iota
	TableSimple
)

type tableWriter interface {
	append(map[string]string)
	render()
}

type ColumnDef struct {
	Name         string
	Sources      []string
	Format       string
	ValueMapping []map[string]string
	FormatFunc   func(values map[string]string) string
}

func (d *ColumnDef) GetSources() []string {
	if len(d.Sources) == 0 {
		return []string{d.Name}
	}
	return d.Sources
}

func (d *ColumnDef) GetFormat() string {
	if d.Format == "" {
		return "%s"
	}
	return d.Format
}
