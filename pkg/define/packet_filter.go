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

package define

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func PacketFilterResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             packetFilterListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: packetFilterListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:          schema.CommandCreate,
			Params:        packetFilterCreateParam(),
			IncludeFields: packetFilterDetailIncludes(),
			ExcludeFields: packetFilterDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        packetFilterReadParam(),
			IncludeFields: packetFilterDetailIncludes(),
			ExcludeFields: packetFilterDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        packetFilterUpdateParam(),
			IncludeFields: packetFilterDetailIncludes(),
			ExcludeFields: packetFilterDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        packetFilterDeleteParam(),
			IncludeFields: packetFilterDetailIncludes(),
			ExcludeFields: packetFilterDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         50,
		},
		"rule-info": {
			Type:               schema.CommandManipulateMulti,
			Aliases:            []string{"rules", "rule-list"},
			Params:             packetFilterRuleListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: packetFilterRuleListColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
			NoSelector:         true,
			Category:           "rule",
			Order:              10,
		},
		"rule-add": {
			Type:               schema.CommandManipulateSingle,
			Params:             packetFilterRuleAddParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: packetFilterRuleListColumns(),
			UseCustomCommand:   true,
			NoSelector:         true,
			Category:           "rule",
			Order:              20,
		},
		"rule-update": {
			Type:               schema.CommandManipulateSingle,
			Params:             packetFilterRuleUpdateParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: packetFilterRuleListColumns(),
			UseCustomCommand:   true,
			NoSelector:         true,
			Category:           "rule",
			Order:              30,
		},
		"rule-delete": {
			Type:               schema.CommandManipulateSingle,
			Params:             packetFilterRuleDeleteParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: packetFilterRuleListColumns(),
			UseCustomCommand:   true,
			ConfirmMessage:     "delete rule",
			NoSelector:         true,
			Category:           "rule",
			Order:              40,
		},
		"interface-connect": {
			Type:             schema.CommandManipulateSingle,
			Params:           packetFilterInterfaceConnectParam(),
			UseCustomCommand: true,
			NoOutput:         true,
			NoSelector:       true,
			Category:         "interface",
			Order:            10,
		},
		"interface-disconnect": {
			Type:             schema.CommandManipulateSingle,
			Params:           packetFilterInterfaceDisconnectParam(),
			UseCustomCommand: true,
			NoOutput:         true,
			NoSelector:       true,
			Category:         "interface",
			Order:            20,
		},
	}

	return &schema.Resource{
		Commands:          commands,
		ResourceCategory:  CategoryNetworking,
		CommandCategories: packetFilterCommandCategories,
	}
}

var packetFilterCommandCategories = []schema.Category{
	{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       10,
	},
	{
		Key:         "rule",
		DisplayName: "Filter-Rule Management",
		Order:       20,
	},
	{
		Key:         "interface",
		DisplayName: "Connection Management",
		Order:       30,
	},
}

func packetFilterListParam() map[string]*schema.Parameter {
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

func packetFilterCreateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"name":        paramRequiredName,
		"description": paramDescription,
	}
}

func packetFilterReadParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func packetFilterUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"name":        paramName,
		"description": paramDescription,
	}
}

func packetFilterDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func packetFilterRuleListParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

var allowPacketFilterProtocol = []string{"tcp", "udp", "icmp", "fragment", "ip"}

func packetFilterRuleAddParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "index to insert rule into",
			DefaultValue: 1,
			Category:     "rule",
			Order:        1,
		},
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target protocol[tcp/udp/icmp/fragment/ip]",
			ValidateFunc: validateInStrValues(allowPacketFilterProtocol...),
			Category:     "rule",
			Order:        10,
		},
		"source-network": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source network[A.A.A.A] or [A.A.A.A/N (N=1..31)] or [A.A.A.A/M.M.M.M]",
			ValidateFunc: validatePacketFilterSourceNetwork(),
			Category:     "rule",
			Order:        20,
		},
		"source-port": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]",
			ValidateFunc: validatePacketFilterPort(),
			Category:     "rule",
			Order:        30,
		},
		"destination-port": {
			Type:         schema.TypeString,
			Aliases:      []string{"dest-port"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set destination port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]",
			ValidateFunc: validatePacketFilterPort(),
			Category:     "rule",
			Order:        40,
		},
		"action": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set action[allow/deny]",
			ValidateFunc: validateInStrValues("allow", "deny"),
			Category:     "rule",
			Order:        50,
		},
		"description": paramDescription,
	}
}

func packetFilterRuleUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target rule",
			Required:    true,
			Category:    "rule",
			Order:       1,
		},
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target protocol[tcp/udp/icmp/fragment/ip]",
			ValidateFunc: validateInStrValues(allowPacketFilterProtocol...),
			Category:     "rule",
			Order:        10,
		},
		"source-network": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source network[A.A.A.A] or [A.A.A.A/N (N=1..31)] or [A.A.A.A/M.M.M.M]",
			ValidateFunc: validatePacketFilterSourceNetwork(),
			Category:     "rule",
			Order:        20,
		},
		"source-port": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]",
			ValidateFunc: validatePacketFilterPort(),
			Category:     "rule",
			Order:        30,
		},
		"destination-port": {
			Type:         schema.TypeString,
			Aliases:      []string{"dest-port"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set destination port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]",
			ValidateFunc: validatePacketFilterPort(),
			Category:     "rule",
			Order:        40,
		},
		"action": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set action[allow/deny]",
			ValidateFunc: validateInStrValues("allow", "deny"),
			Category:     "rule",
			Order:        50,
		},
		"description": paramDescription,
	}
}

func packetFilterRuleDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target rule",
			Required:    true,
			Category:    "rule",
			Order:       1,
		},
	}
}

func packetFilterInterfaceConnectParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"interface-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set interface ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			Category:     "interface",
			Order:        1,
		},
	}
}

func packetFilterInterfaceDisconnectParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"interface-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set interface ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			Category:     "interface",
			Order:        1,
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

			var strNum1 string
			strNum2 := "0"
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
