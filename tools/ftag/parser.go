package ftag

import (
	"fmt"
	"strings"
)

// const defaultTagName = "cli"

// type ParserConfig struct {
// 	TagName string
// }

type Parser struct {
	//	Config *ParserConfig
}

func Parse(t string) (Tag, error) {
	parser := &Parser{}
	return parser.Parse(t)
}

const (
	aliasesKey   = "aliases"
	shortHandKey = "short"
	descKey      = "desc"
)

func (p *Parser) Parse(t string) (Tag, error) {
	tag := Tag{}
	t = strings.TrimSpace(t)
	if t == "" {
		return tag, nil
	}

	tokens := strings.Split(t, `,`) // 1つ以上の要素を含むスライスを返す
	name := strings.TrimSpace(tokens[0])
	if name != "" {
		tag.Name = name
	}
	if len(tokens) > 1 {
		for _, token := range tokens[1:] {
			token = strings.TrimSpace(token)
			kv := strings.Split(token, `=`)
			if len(kv) != 2 {
				return tag, fmt.Errorf("got invalid tag value: %q", token)
			}

			key := strings.TrimSpace(kv[0])
			val := strings.TrimSpace(kv[1])

			switch key {
			case aliasesKey:
				names := strings.Split(val, ` `)
				for _, n := range names {
					if n != "" {
						tag.Aliases = append(tag.Aliases, n)
					}
				}
			case shortHandKey:
				if len(val) != 1 {
					return tag, fmt.Errorf("got invalid tag value: key 'short' must have only 1 character: %q", token)
				}
				tag.Shorthand = kv[1]
			case descKey:
				tag.Description = kv[1]
			default:
				return tag, fmt.Errorf("got invalid tag key: %q", token)
			}
		}
	}
	return tag, nil
}
