package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func DNSResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             dnsListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: dnsListColumns(),
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           dnsCreateParam(),
			IncludeFields:    dnsDetailIncludes(),
			ExcludeFields:    dnsDetailExcludes(),
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        dnsReadParam(),
			IncludeFields: dnsDetailIncludes(),
			ExcludeFields: dnsDetailExcludes(),
		},
		"update": {
			Type:             schema.CommandUpdate,
			Params:           dnsUpdateParam(),
			IncludeFields:    dnsDetailIncludes(),
			ExcludeFields:    dnsDetailExcludes(),
			UseCustomCommand: true,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        dnsDeleteParam(),
			IncludeFields: dnsDetailIncludes(),
			ExcludeFields: dnsDetailExcludes(),
		},
		"record-list": {
			Type:               schema.CommandManipulateSingle,
			Params:             dnsRecordListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: dnsRecordListColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
		},
		"record-add": {
			Type:               schema.CommandManipulateSingle,
			Params:             dnsRecordAddParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: dnsRecordListColumns(),
			UseCustomCommand:   true,
		},
		"record-update": {
			Type:               schema.CommandManipulateSingle,
			Params:             dnsRecordUpdateParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: dnsRecordListColumns(),
			UseCustomCommand:   true,
		},
		"record-delete": {
			Type:               schema.CommandManipulateSingle,
			Params:             dnsRecordDeleteParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: dnsRecordListColumns(),
			UseCustomCommand:   true,
			ConfirmMessage:     "delete record",
		},
	}

	return &schema.Resource{
		Commands:            commands,
		ResourceCategory:    CategoryCommonServiceItem,
		ListResultFieldName: "CommonServiceDNSItems",
	}
}

func dnsListParam() map[string]*schema.Schema {
	return CommonListParam
}

func dnsListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "NameServers",
			Sources: []string{"Status.NS.0", "Status.NS.1"},
			Format:  "%s / %s",
		},
	}
}

func dnsRecordListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "Type"},
		{Name: "Name"},
		{Name: "RData"},
	}
}

func dnsDetailIncludes() []string {
	return []string{}
}

func dnsDetailExcludes() []string {
	return []string{}
}

func dnsCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerPathThrough,
			Description:  "set DNS zone name",
			Required:     true,
			ValidateFunc: validateStrLen(2, 63),
		},
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func dnsReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func dnsUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func dnsDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func dnsRecordListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

var allowDNSTypes = []string{
	"a", "aaaa", "ns", "cname", "mx", "txt", "srv",
	"A", "AAAA", "NS", "CNAME", "MX", "TXT", "SRV",
}

func dnsRecordAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set name",
			Required:     true,
			ValidateFunc: validateStrLen(1, 63),
		},
		"type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set record type[A/AAAA/NS/CNAME/MX/TXT/SRV]",
			Required:     true,
			ValidateFunc: validateInStrValues(allowDNSTypes...),
			CompleteFunc: completeInStrValues(allowDNSTypes...),
		},
		"value": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set record data",
		},
		"ttl": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ttl",
			DefaultValue: 3600,
			ValidateFunc: validateIntRange(10, 3600000),
		},

		"mx-priority": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set MX priority",
			DefaultValue: 10,
			ValidateFunc: validateIntRange(1, 65535),
		},

		"srv-priority": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			DefaultValue: 0,
			ValidateFunc: validateIntRange(0, 65535),
		},
		"srv-weight": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			DefaultValue: 0,
			ValidateFunc: validateIntRange(0, 65535),
		},
		"srv-port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			DefaultValue: 0,
			ValidateFunc: validateIntRange(1, 65535),
		},
		"srv-target": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			ValidateFunc: validateStrLen(1, 254),
		},
	}
}
func dnsRecordUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target record",
			Required:    true,
		},
		"name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set name",
			ValidateFunc: validateStrLen(1, 63),
		},
		"type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set record type[A/AAAA/NS/CNAME/MX/TXT/SRV]",
			ValidateFunc: validateInStrValues(allowDNSTypes...),
			CompleteFunc: completeInStrValues(allowDNSTypes...),
		},
		"value": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set record data",
		},
		"ttl": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ttl",
			ValidateFunc: validateIntRange(10, 3600000),
		},

		"mx-priority": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set MX priority",
			ValidateFunc: validateIntRange(1, 65535),
		},

		"srv-priority": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			ValidateFunc: validateIntRange(0, 65535),
		},
		"srv-weight": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			ValidateFunc: validateIntRange(0, 65535),
		},
		"srv-port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			ValidateFunc: validateIntRange(1, 65535),
		},
		"srv-target": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			ValidateFunc: validateStrLen(1, 254),
		},
	}
}
func dnsRecordDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target record",
			Required:    true,
		},
	}
}
