package output

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
