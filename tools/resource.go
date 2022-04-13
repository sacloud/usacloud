// Copyright 2017-2022 The Usacloud Authors
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
	"strings"

	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/naming"
)

type Resource struct {
	*core.Resource
	Commands []*Command
}

func NewResource(r *core.Resource) *Resource {
	resource := &Resource{Resource: r}

	for _, command := range r.Commands() {
		resource.Commands = append(resource.Commands, NewCommand(resource, command))
	}

	return resource
}

func NewResources(resources []*core.Resource) []*Resource {
	var results []*Resource
	for _, r := range resources {
		results = append(results, NewResource(r))
	}
	return results
}

func (r *Resource) Parent() *Resource {
	if r.Resource.Parent() != nil {
		return NewResource(r.Resource.Parent())
	}
	return nil
}

func (r *Resource) PackageDirName() string {
	prefix := ""
	if r.Parent() != nil {
		prefix = r.Parent().PackageDirName() + "/"
	}

	n := naming.ToLower(r.Name)
	// ハイフンやアンダーバーは除去する
	n = strings.ReplaceAll(n, "-", "")
	n = strings.ReplaceAll(n, "_", "")
	switch n {
	case "switch":
		n = "swytch"
	case "interface":
		n = "iface"
	}
	return prefix + n
}
