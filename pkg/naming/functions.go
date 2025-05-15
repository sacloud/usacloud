// Copyright 2017-2025 The sacloud/usacloud Authors
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

package naming

import (
	"fmt"
	"strings"

	"github.com/huandu/xstrings"
)

var normalizationWords = map[string]string{
	"Acme": "ACME",
	"Dns":  "DNS",
	"Gslb": "GSLB",

	"Ipv4":   "IPv4",
	"Ipv6":   "IPv6",
	"iPv4":   "ipv4",
	"iPv6":   "ipv6",
	"i_pv4":  "ipv4",
	"i_pv6":  "ipv6",
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
	"Nmi": "NMI",

	"Lb": "LB",

	"Sim": "SIM",
	"Ssh": "SSH",
	"Vpc": "VPC",
	"Vpn": "VPN",

	"L2tp":   "L2TP",
	"l_2tp":  "l2tp",
	"l-2tp":  "l2tp",
	"dns-11": "dns1", // xstrings.ToKebabCase対応
	"dns-22": "dns2",

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

func Normalize(name string) string {
	n := name
	for k, v := range normalizationWords {
		if strings.Contains(n, k) && !isIncludeInNormalizationIgnoreWords(n) {
			n = strings.ReplaceAll(n, k, v)
		}
	}
	return n
}

func ToSnakeCase(name string) string {
	return Normalize(xstrings.ToSnakeCase(name))
}

func ToKebabCase(name string) string {
	return Normalize(xstrings.ToKebabCase(name))
}

func ToCamelCase(name string) string {
	return Normalize(xstrings.ToPascalCase(xstrings.ToSnakeCase(name)))
}

func ToCamelCaseWithFirstLower(name string) string {
	return Normalize(xstrings.FirstRuneToLower(xstrings.ToCamelCase(xstrings.ToSnakeCase(name))))
}

func ToLower(name string) string {
	return strings.ToLower(Normalize(xstrings.ToCamelCase(xstrings.ToSnakeCase(name))))
}

func ToCLIFlag(name string) string {
	format := "--%s"
	if len(name) == 1 {
		format = "-%s"
	}
	// HACK nameに"tags[0]"のようなものがきた場合、tags[-0]となる。
	// さほど登場パターンはないため"[-"を置き換えれば実用に足るはず
	v := ToKebabCase(name)
	v = strings.ReplaceAll(v, "[-", "[")
	return fmt.Sprintf(format, v)
}

func FlattenStringList(list []string) string {
	if len(list) > 0 {
		return fmt.Sprintf(`"%s"`, strings.Join(list, `","`))
	}
	return ""
}
