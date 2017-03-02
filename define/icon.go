package define

import (
	"encoding/base64"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
	"io/ioutil"
)

func IconResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"l", "ls", "find"},
			Params:             iconListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: iconListColumns(),
		},
		"create": {
			Type:          schema.CommandCreate,
			Aliases:       []string{"c"},
			Params:        iconCreateParam(),
			IncludeFields: iconDetailIncludes(),
			ExcludeFields: iconDetailExcludes(),
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        iconReadParam(),
			IncludeFields: iconDetailIncludes(),
			ExcludeFields: iconDetailExcludes(),
		},
		"update": {
			Type:          schema.CommandUpdate,
			Aliases:       []string{"u"},
			Params:        iconUpdateParam(),
			IncludeFields: iconDetailIncludes(),
			ExcludeFields: iconDetailExcludes(),
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"d", "rm"},
			Params:        iconDeleteParam(),
			IncludeFields: iconDetailIncludes(),
			ExcludeFields: iconDetailExcludes(),
		},
	}

	return &schema.Resource{
		Commands:         commands,
		ResourceCategory: CategoryCommonItem,
	}
}

func iconListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramScopeCond)
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
			Description:   "set icon image",
			Required:      true,
			ValidateFunc:  validateFileExists(),
			CustomHandler: iconSetImageContentUseBase64,
		},
	}
}

func iconReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func iconUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id":   paramID,
		"name": paramName,
		"tags": paramTags,
	}
}

func iconDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
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
