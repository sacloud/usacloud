// Copyright 2017-2021 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tools

import (
	"fmt"

	"github.com/fatih/structs"
	"github.com/sacloud/usacloud/pkg/naming"
)

func (c *Command) CLIName() string {
	return naming.ToKebabCase(c.Name)
}

func (c *Command) CLIVariableFuncName() string {
	return fmt.Sprintf("%sCmd", naming.ToCamelCaseWithFirstLower(c.Name))
}

func (c *Command) ServiceRequestTypeName() string {
	name := c.Name
	if c.ServiceFuncAltName != "" {
		name = c.ServiceFuncAltName
	}
	return fmt.Sprintf("%sRequest", naming.ToCamelCase(name))
}

func (c *Command) ServiceFuncName() string {
	name := c.Name
	if c.ServiceFuncAltName != "" {
		name = c.ServiceFuncAltName
	}
	return fmt.Sprintf("%sWithContext", naming.ToCamelCase(name))
}

func (c *Command) CLICommandParameterTypeName() string {
	if c.Command.ParameterInitializer == nil {
		return ""
	}
	return structs.Name(c.Command.ParameterInitializer())
}

func (c *Command) InputParameterVariable() string {
	return fmt.Sprintf("%s%sParam", naming.ToCamelCaseWithFirstLower(c.Resource.Name), naming.ToCamelCase(c.Name))
}

func (c *Command) InputParameterTypeName() string {
	return fmt.Sprintf("%s%sParam", naming.ToCamelCase(c.Name), naming.ToCamelCase(c.Resource.Name))
}
