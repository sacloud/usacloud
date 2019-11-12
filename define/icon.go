// Copyright 2017-2019 The Usacloud Authors
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

package define

import (
	"encoding/base64"
	"io/ioutil"

	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func IconResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "selector"},
			Params:             iconListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: iconListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:          schema.CommandCreate,
			Params:        iconCreateParam(),
			IncludeFields: iconDetailIncludes(),
			ExcludeFields: iconDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        iconReadParam(),
			IncludeFields: iconDetailIncludes(),
			ExcludeFields: iconDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        iconUpdateParam(),
			IncludeFields: iconDetailIncludes(),
			ExcludeFields: iconDetailExcludes(),
			Category:      "basics",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        iconDeleteParam(),
			IncludeFields: iconDetailIncludes(),
			ExcludeFields: iconDetailExcludes(),
			Category:      "basics",
			Order:         50,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		ResourceCategory: CategoryCommonItem,
	}
}

func iconListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramScopeCond, paramTagsCond)
}

func iconListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{Name: "Scope"},
		{Name: "URL"},
	}
}

func iconDetailIncludes() []string {
	return []string{}
}

func iconDetailExcludes() []string {
	return []string{}
}

func iconCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": paramRequiredName,
		"tags": paramTags,
		"image": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerCustomFunc,
			Description:   "set file path for upload",
			Required:      true,
			ValidateFunc:  validateFileExists(),
			CustomHandler: iconSetImageContentUseBase64,
			Category:      "icon",
			Order:         10,
		},
	}
}

func iconReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func iconUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": paramName,
		"tags": paramTags,
	}
}

func iconDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func iconSetImageContentUseBase64(name string, s interface{}, d interface{}) {

	type imageSrc interface {
		GetImage() string
	}
	type imageDest interface {
		SetImage(img string)
	}

	src, ok1 := s.(imageSrc)
	dest, ok2 := d.(imageDest)
	if ok1 && ok2 {
		filePath := src.GetImage()
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		img := base64.StdEncoding.EncodeToString(b)
		dest.SetImage(img)
	}
}
