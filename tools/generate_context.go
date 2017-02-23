package tools

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
	"go/build"
	"path/filepath"
)

type GenerateContext struct {
	R           string
	C           string
	P           string
	ResourceDef map[string]*schema.Resource
}

func NewGenerateContext() *GenerateContext {
	ctx := &GenerateContext{
		ResourceDef: define.Resources,
	}
	return ctx
}

func (c *GenerateContext) SetCurrentR(k string) {
	c.R = k
}

/*
   Get name functions
*/

func (c *GenerateContext) CamelR() string {
	return ToCamelCaseName(c.R)
}

func (c *GenerateContext) Camelr() string {
	return ToCamelWithFirstLower(c.R)
}

func (c *GenerateContext) DashR() string {
	return ToDashedName(c.R)
}

func (c *GenerateContext) SnakeR() string {
	return ToSnakeCaseName(c.R)
}

func (c *GenerateContext) CamelC() string {
	return ToCamelCaseName(c.C)
}

func (c *GenerateContext) Camelc() string {
	return ToCamelWithFirstLower(c.C)
}

func (c *GenerateContext) DashC() string {
	return ToDashedName(c.C)
}

func (c *GenerateContext) SnakeC() string {
	return ToSnakeCaseName(c.C)
}

func (c *GenerateContext) CamelP() string {
	return ToCamelCaseName(c.P)
}

func (c *GenerateContext) Camelp() string {
	return ToCamelWithFirstLower(c.P)
}

func (c *GenerateContext) DashP() string {
	return ToDashedName(c.P)
}

func (c *GenerateContext) SnakeP() string {
	return ToSnakeCaseName(c.P)
}

func (c *GenerateContext) Gopath() string {
	gopath := build.Default.GOPATH
	gopath = filepath.SplitList(gopath)[0]
	return gopath
}

/*
   Get current context schema
*/

func (c *GenerateContext) CurrentResource() *schema.Resource {
	return c.ResourceDef[c.R]
}

func (c *GenerateContext) CurrentCommand() *schema.Command {
	return c.CurrentResource().Commands[c.C]
}

func (c *GenerateContext) CurrentParam() *schema.Schema {
	return c.CurrentCommand().Params[c.P]
}

/*
   Get contextual name functions
*/

func (c *GenerateContext) InputModelFileName() string {
	return fmt.Sprintf("params_%s_gen.go", c.SnakeR())
}

func (c *GenerateContext) InputModelTypeName() string {
	return fmt.Sprintf("%s%sParam", c.CamelC(), c.CamelR())
}

func (c *GenerateContext) InputParamFieldName() string {
	return c.CamelP()
}

func (c *GenerateContext) InputParamFlagName() string {
	return c.DashP()
}

func (c *GenerateContext) InputParamSetterFuncName() string {
	n := c.CurrentParam().DestinationProp
	if n == "" {
		n = fmt.Sprintf("Set%s", c.CamelP())
	}
	return n
}

func (c *GenerateContext) InputParamDestinationName() string {
	n := c.CurrentParam().DestinationProp
	if n == "" {
		n = c.CamelP()
	}
	return n
}

func (c *GenerateContext) InputParamCLIFlagName() string {
	return ToCLIFlagName(c.P)
}

func (c *GenerateContext) InputParamVariableName() string {
	return fmt.Sprintf("%sParam", ToCamelWithFirstLower(c.C))
}

func (c *GenerateContext) CommandFuncName() string {
	return fmt.Sprintf("%s%s", c.CamelR(), c.CamelC())
}

func (c *GenerateContext) CommandFileName(useCustomCommand bool) string {
	if useCustomCommand {
		return fmt.Sprintf("command_%s_%s.go", c.SnakeR(), c.SnakeC())
	}
	return fmt.Sprintf("command_%s_%s_gen.go", c.SnakeR(), c.SnakeC())
}

func (c *GenerateContext) CLICommandsFileName() string {
	return fmt.Sprintf("cli_%s_gen.go", c.SnakeR())
}

func (c *GenerateContext) CommandResourceName() string {
	n := c.CurrentCommand().AltResource
	if n == "" {
		n = c.CurrentResource().AltResource
	}
	if n == "" {
		n = c.CamelR()
	}
	return n
}

func (c *GenerateContext) FindResultFieldName() string {
	return c.CurrentCommand().ListResultFieldName
}
