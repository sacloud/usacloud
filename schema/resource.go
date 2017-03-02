package schema

import "sort"

type Resource struct {
	Aliases             []string
	Usage               string
	Commands            map[string]*Command
	AltResource         string // 空の場合はResourceのキーをCamelizeしてsacloud.XXXを対象とする。
	ListResultFieldName string
	CommandCategories   []Category
	ResourceCategory    Category
}

func (r *Resource) CommandCategory(key string) *Category {
	if key == "" {
		return DefaultCommandCategory
	}

	for _, cat := range r.CommandCategories {
		if cat.Key == key {
			return &cat
		}
	}

	return nil
}

func (c *Resource) SortedCommands() SortableCommands {

	params := SortableCommands{}
	for k, v := range c.Commands {
		params = append(params, SortableCommand{
			CommandKey: k,
			Command:    v,
			Category:   c.CommandCategory(v.Category),
		})
	}

	sort.Sort(params)
	return params
}

type SortableCommand struct {
	Category   *Category
	Command    *Command
	CommandKey string
}

type SortableCommands []SortableCommand

func (s SortableCommands) Len() int {
	return len(s)
}

func (s SortableCommands) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortableCommands) Less(i, j int) bool {

	if s[i].Category.Order == s[j].Category.Order {

		if s[i].Command.Order == s[j].Command.Order {
			return s[i].CommandKey < s[j].CommandKey
		} else {
			return s[i].Command.Order < s[j].Command.Order
		}

	} else {
		return s[i].Category.Order < s[j].Category.Order
	}
}
