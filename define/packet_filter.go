package define

import (
	"fmt"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
	"strconv"
	"strings"
)

func PacketFilterResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"l", "ls", "find"},
			Params:             packetFilterListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: packetFilterListColumns(),
		},
		"create": {
			Type:          schema.CommandCreate,
			Aliases:       []string{"c"},
			Params:        packetFilterCreateParam(),
			IncludeFields: packetFilterDetailIncludes(),
			ExcludeFields: packetFilterDetailExcludes(),
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        packetFilterReadParam(),
			IncludeFields: packetFilterDetailIncludes(),
			ExcludeFields: packetFilterDetailExcludes(),
		},
		"update": {
			Type:          schema.CommandUpdate,
			Aliases:       []string{"u"},
			Params:        packetFilterUpdateParam(),
			IncludeFields: packetFilterDetailIncludes(),
			ExcludeFields: packetFilterDetailExcludes(),
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"d", "rm"},
			Params:        packetFilterDeleteParam(),
			IncludeFields: packetFilterDetailIncludes(),
			ExcludeFields: packetFilterDetailExcludes(),
		},
		"rule-list": {
			Type:               schema.CommandManipulateMulti,
			Aliases:            []string{"rules"},
			Params:             packetFilterRuleListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: packetFilterRuleListColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
		},
		"rule-add": {
			Type:               schema.CommandManipulateSingle,
			Params:             packetFilterRuleAddParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: packetFilterRuleListColumns(),
			UseCustomCommand:   true,
		},
		"rule-update": {
			Type:               schema.CommandManipulateSingle,
			Params:             packetFilterRuleUpdateParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: packetFilterRuleListColumns(),
			UseCustomCommand:   true,
		},
		"rule-delete": {
			Type:               schema.CommandManipulateSingle,
			Params:             packetFilterRuleDeleteParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: packetFilterRuleListColumns(),
			UseCustomCommand:   true,
			ConfirmMessage:     "delete rule",
		},
		"interface-connect": {
			Type:             schema.CommandManipulateSingle,
			Params:           packetFilterInterfaceConnectParam(),
			UseCustomCommand: true,
			NoOutput:         true,
		},
		"interface-disconnect": {
			Type:             schema.CommandManipulateSingle,
			Params:           packetFilterInterfaceDisconnectParam(),
			UseCustomCommand: true,
			NoOutput:         true,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		ResourceCategory: CategoryNetworking,
	}
}

func packetFilterListParam() map[string]*schema.Schema {
	return CommonListParam
}

func packetFilterListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
	}
}

func packetFilterRuleListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "Protocol"},
		{
			Name:    "Source-Network",
			Sources: []string{"SourceNetwork"},
		},
		{
			Name:    "Source-Port",
			Sources: []string{"SourcePort"},
		},
		{
			Name:    "Destination-Port",
			Sources: []string{"DestinationPort"},
		},
		{Name: "Action"},
		{Name: "Description"},
	}
}

func packetFilterDetailIncludes() []string {
	return []string{}
}

func packetFilterDetailExcludes() []string {
	return []string{}
}

func packetFilterCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramRequiredName,
		"description": paramDescription,
	}
}

func packetFilterReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func packetFilterUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
	}
}

func packetFilterDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func packetFilterRuleListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

var allowPacketFilterProtocol = []string{"tcp", "udp", "icmp", "fragment", "ip"}

func packetFilterRuleAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "index to insert rule into",
			DefaultValue: 1,
		},
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target protocol[tcp/udp/icmp/fragment/ip]",
			ValidateFunc: validateInStrValues(allowPacketFilterProtocol...),
			CompleteFunc: completeInStrValues(allowPacketFilterProtocol...),
		},
		"source-network": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source network[A.A.A.A] or [A.A.A.A/N (N=1..31)] or [A.A.A.A/M.M.M.M]",
			ValidateFunc: validatePacketFilterSourceNetwork(),
		},
		"source-port": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]",
			ValidateFunc: validatePacketFilterPort(),
		},
		"destination-port": {
			Type:         schema.TypeString,
			Aliases:      []string{"dest-port"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set destination port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]",
			ValidateFunc: validatePacketFilterPort(),
		},
		"action": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set action[allow/deny]",
			ValidateFunc: validateInStrValues("allow", "deny"),
			CompleteFunc: completeInStrValues("allow", "deny"),
		},
		"description": paramDescription,
	}
}

func packetFilterRuleUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target rule",
			Required:    true,
		},
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target protocol[tcp/udp/icmp/fragment/ip]",
			ValidateFunc: validateInStrValues(allowPacketFilterProtocol...),
			CompleteFunc: completeInStrValues(allowPacketFilterProtocol...),
		},
		"source-network": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source network[A.A.A.A] or [A.A.A.A/N (N=1..31)] or [A.A.A.A/M.M.M.M]",
			ValidateFunc: validatePacketFilterSourceNetwork(),
		},
		"source-port": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]",
			ValidateFunc: validatePacketFilterPort(),
		},
		"destination-port": {
			Type:         schema.TypeString,
			Aliases:      []string{"dest-port"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set destination port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]",
			ValidateFunc: validatePacketFilterPort(),
		},
		"action": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set action[allow/deny]",
			ValidateFunc: validateInStrValues("allow", "deny"),
			CompleteFunc: completeInStrValues("allow", "deny"),
		},
		"description": paramDescription,
	}
}

func packetFilterRuleDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target rule",
			Required:    true,
		},
	}
}

func packetFilterInterfaceConnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"interface-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set interface ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeInterfaceID(),
		},
	}
}

func packetFilterInterfaceDisconnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"interface-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set interface ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeInterfaceID(),
		},
	}
}

func validatePacketFilterSourceNetwork() schema.ValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if value, ok := object.(string); ok {
			if value == "" {
				return res
			}

			tokens := strings.Split(value, "/")
			tokenLen := len(tokens)
			ipv4Validator := validateIPv4Address()
			validateError := fmt.Errorf("%q: Invalid format,need [A.A.A.A] or [A.A.A.A/N (N=1..31)] or [A.A.A.A/M.M.M.M]", fieldName)

			// has slash?
			switch tokenLen {
			case 1:
				// ipv4 address only
				res = append(res, ipv4Validator(fieldName, object)...)
			case 2:

				// first , validate ipv4 address
				errs := ipv4Validator(fieldName, tokens[0])
				if len(errs) == 0 {

					// next , second token is Number?
					num, err := strconv.Atoi(tokens[1])
					if err == nil {
						// token is number. value is in range?
						if !(1 <= num && num <= 31) {
							res = append(res, validateError)
						}
					} else {
						// second token is ipv4 addr?
						errs := ipv4Validator(fieldName, tokens[1])
						if len(errs) > 0 {
							res = append(res, validateError)
						}
					}

				} else {
					res = append(res, validateError)
				}

			default:
				res = append(res, validateError)
			}

		}

		return res
	}
}

func validatePacketFilterPort() schema.ValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if value, ok := object.(string); ok {
			if value == "" {
				return res
			}

			validateError := fmt.Errorf("%q: Invalid format,need [N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]", fieldName)
			isInRange := func(num int) bool {
				return 0 <= num && num <= 65535
			}

			strNum1, strNum2 := "0", "0"
			tokens := strings.Split(value, "-")
			if len(tokens) == 2 {
				strNum1, strNum2 = tokens[0], tokens[1]
			} else {
				tokens = strings.Split(value, "/")
				if len(tokens) == 2 {
					strNum1, strNum2 = tokens[0], tokens[1]
				} else {
					strNum1 = value
				}

			}

			num1, num2 := 0, 0
			var err error
			// number format
			num1, err = strconv.Atoi(strNum1)
			if err != nil {
				res = append(res, validateError)
			}
			num2, err = strconv.Atoi(strNum2)
			if err != nil {
				res = append(res, validateError)
			}

			if len(res) == 0 {
				for _, v := range []int{num1, num2} {
					if !isInRange(v) {
						res = append(res, validateError)
						break
					}
				}
			}

			if len(res) == 0 {
				if num2 > 0 && num1 > num2 {
					res = append(res, validateError)
				}
			}
		}

		return res
	}
}
