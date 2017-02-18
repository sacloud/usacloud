package output

type Output interface {
	Print(...interface{}) error
}

type OutputFormater interface {
	GetIncludeFields() []string
	GetExcludeFields() []string
	GetColumnDefs() []ColumnDef
	GetTableType() OutputTableType
}

type OutputTableType int //go:generate stringer -type=OutputTableType :: manual
const (
	TableDetail OutputTableType = iota
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
