package ftag

// Tag structにつけられたftagの値
//
// 例: タグが`example,aliases=foo bar,short=e,desc=foobar`の場合
// - Name: example
// - Aliases: foo, bar
// - Shorthand: e
// - Description: foobar
type Tag struct {
	Name        string
	Aliases     []string
	Shorthand   string
	Description string
}
