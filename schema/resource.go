package schema

type Resource struct {
	Aliases     []string
	Usage       string
	Commands    map[string]*Command
	AltResource string // 空の場合はResourceのキーをCamelizeしてsacloud.XXXを対象とする。
}
