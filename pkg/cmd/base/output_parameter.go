package base

type OutputParameter struct {
	OutputType string `cli:",short=o,aliases=out,category=output,desc=Output format: one of the following [table/json/yaml]"`
	Quiet      bool   `cli:",short=q,category=output,desc=Output IDs only"`
	Format     string `cli:",aliases=fmt,category=output,desc=Output format in Go templates"`
	FormatFile string `cli:",category=output,desc=Output format in Go templates(from file)"`
	Query      string `cli:",category=output,desc=JMESPath query"`
	QueryFile  string `cli:",category=output,desc=JMESPath query(from file)"`
}
