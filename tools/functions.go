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

package tools

import (
	"fmt"
	"go/format"
	"log"
	"strings"

	"github.com/sacloud/libsacloud/sacloud"

	"github.com/huandu/xstrings"
)

func Sformat(buf []byte) []byte {
	src, err := format.Source(buf)
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		log.Printf("generated: \n%s", string(buf))
		return buf
	}
	return src
}

var normalizationWords = map[string]string{
	"Acme": "ACME",
	"Dns":  "DNS",
	"Gslb": "GSLB",

	"Ipv4":   "IPv4",
	"Ipv6":   "IPv6",
	"iPv4":   "ipv4",
	"iPv6":   "ipv6",
	"i_pv_4": "ipv4",
	"i_pv_6": "ipv6",
	"ipv_4":  "ipv4",
	"ipv_6":  "ipv6",
	"i-pv-4": "ipv4",
	"i-pv-6": "ipv6",
	"ipv-4":  "ipv4",
	"ipv-6":  "ipv6",

	"Iso": "ISO",
	"Cpu": "CPU",
	"Ftp": "FTP",
	"Nfs": "NFS",

	"Lb": "LB",

	"Sim": "SIM",
	"Ssh": "SSH",
	"Vpc": "VPC",
	"Vpn": "VPN",

	"L2tp":  "L2TP",
	"l_2tp": "l2tp",
	"l-2tp": "l2tp",

	"Ipsec": "IPsec",
}

var normalizationIgnoreWords = []string{"Simple", "simple"}

func isIncludeInNormalizationIgnoreWords(w string) bool {
	for _, v := range normalizationIgnoreWords {
		if strings.Contains(w, v) {
			return true
		}
	}
	return false
}

func NormalizeName(name string) string {
	n := name
	for k, v := range normalizationWords {
		if strings.Contains(n, k) && !isIncludeInNormalizationIgnoreWords(n) {
			n = strings.Replace(n, k, v, -1)
		}
	}
	return n
}

func ToSnakeCaseName(name string) string {
	return NormalizeName(xstrings.ToSnakeCase(name))
}

func ToDashedName(name string) string {
	return NormalizeName(xstrings.ToKebabCase(name))
}

func ToCamelCaseName(name string) string {
	return NormalizeName(xstrings.ToCamelCase(xstrings.ToSnakeCase(name)))
}

func ToCamelWithFirstLower(name string) string {
	return NormalizeName(xstrings.FirstRuneToLower(xstrings.ToCamelCase(xstrings.ToSnakeCase(name))))
}

func ToCLIFlagName(name string) string {
	format := "--%s"
	if len(name) == 1 {
		format = "-%s"
	}
	return fmt.Sprintf(format, ToDashedName(name))
}

func FlattenStringList(list []string) string {
	if len(list) > 0 {
		return fmt.Sprintf(`"%s"`, strings.Join(list, `","`))
	}
	return ""
}

func FlattenIntList(list []int) string {
	if len(list) > 0 {
		tmp := []string{}
		for _, s := range list {
			tmp = append(tmp, fmt.Sprintf("%d", s))
		}
		return strings.Join(tmp, ",")
	}
	return ""
}

func FlattenUintList(list []uint) string {
	if len(list) > 0 {
		tmp := []string{}
		for _, s := range list {
			tmp = append(tmp, fmt.Sprintf("%d", s))
		}
		return strings.Join(tmp, ",")
	}
	return ""
}
func FlattenInt64List(list []int64) string {
	if len(list) > 0 {
		tmp := []string{}
		for _, s := range list {
			tmp = append(tmp, fmt.Sprintf("%d", s))
		}
		return strings.Join(tmp, ",")
	}
	return ""
}

func FlattenUint64List(list []uint64) string {
	if len(list) > 0 {
		tmp := []string{}
		for _, s := range list {
			tmp = append(tmp, fmt.Sprintf("%d", s))
		}
		return strings.Join(tmp, ",")
	}
	return ""
}

func FlattenFloatList(list []float32) string {
	if len(list) > 0 {
		tmp := []string{}
		for _, s := range list {
			tmp = append(tmp, fmt.Sprintf("%f", s))
		}
		return strings.Join(tmp, ",")
	}
	return ""
}

func FlattenFloat64List(list []float64) string {
	if len(list) > 0 {
		tmp := []string{}
		for _, s := range list {
			tmp = append(tmp, fmt.Sprintf("%f", s))
		}
		return strings.Join(tmp, ",")
	}
	return ""
}

func FlattenIDList(list []sacloud.ID) string {
	if len(list) > 0 {
		tmp := []string{}
		for _, s := range list {
			tmp = append(tmp, fmt.Sprintf("sacloud.ID(%d)", s.Int64()))
		}
		return strings.Join(tmp, ",")
	}
	return ""

}
