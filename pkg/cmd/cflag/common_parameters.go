// Copyright 2017-2020 The Usacloud Authors
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

package cflag

import "github.com/sacloud/libsacloud/v2/sacloud/types"

type NameParameter struct {
	Name string `cli:",category=common,order=10" validate:"required"`
}

type NameUpdateParameter struct {
	Name *string `cli:",category=common,order=10" validate:"omitempty,min=1"`
}

type DescParameter struct {
	Description string `cli:",category=common,order=20" validate:"omitempty,description"`
}

type DescUpdateParameter struct {
	Description *string `cli:",category=common,order=20" validate:"omitempty,description"`
}

type TagsParameter struct {
	Tags []string `cli:",category=common,order=30" validate:"omitempty,tags"`
}

type TagsUpdateParameter struct {
	Tags *[]string `cli:",category=common,order=30" validate:"omitempty,tags"`
}

type IconIDParameter struct {
	IconID types.ID `cli:",category=common,order=40"`
}

type IconIDUpdateParameter struct {
	IconID *types.ID `cli:",category=common,order=40"`
}
