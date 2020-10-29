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
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func WebAccelResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "select"},
			Params:             webAccelListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: webAccelListColumns(),
			UseCustomCommand:   true,
			Category:           "basics",
			Order:              10,
		},
		"read": {
			Type:             schema.CommandRead,
			Params:           webAccelReadParam(),
			IncludeFields:    webAccelDetailIncludes(),
			ExcludeFields:    webAccelDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            20,
		},
		"certificate-info": {
			Type:             schema.CommandRead,
			Aliases:          []string{"cert-info"},
			Params:           webAccelCertInfoParam(),
			IncludeFields:    webAccelCertIncludes(),
			ExcludeFields:    webAccelCertExcludes(),
			UseCustomCommand: true,
			Category:         "certificate",
			Order:            10,
		},
		"certificate-new": {
			Type:             schema.CommandManipulateSingle,
			Aliases:          []string{"cert-new", "cert-create", "certificate-create"},
			Params:           webAccelCertCreateParam(),
			IncludeFields:    webAccelCertIncludes(),
			ExcludeFields:    webAccelCertExcludes(),
			UseCustomCommand: true,
			Category:         "certificate",
			Order:            15,
		},
		"certificate-update": {
			Type:             schema.CommandManipulateSingle,
			Aliases:          []string{"cert-update"},
			Params:           webAccelCertUpdateParam(),
			IncludeFields:    webAccelCertIncludes(),
			ExcludeFields:    webAccelCertExcludes(),
			UseCustomCommand: true,
			Category:         "certificate",
			Order:            20,
		},
		"delete-cache": {
			Type:               schema.CommandCustom,
			Aliases:            []string{"purge"},
			Params:             webAccelDeleteCacheParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: webAccelDeleteCacheColumns(),
			ArgsUsage:          "[URLs]...",
			UseCustomCommand:   true,
			Category:           "cache",
			Order:              10,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		ResourceCategory:    CategoryOther,
		CommandCategories:   webAccelCommandCategories,
		ListResultFieldName: "WebAccelSites",
		IsGlobal:            true,
	}
}

var webAccelCommandCategories = []schema.Category{
	{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       10,
	},
	{
		Key:         "certificate",
		DisplayName: "Certificate Management",
		Order:       20,
	},
	{
		Key:         "cache",
		DisplayName: "Cache Management",
		Order:       30,
	},
	{
		Key:         "other",
		DisplayName: "Other",
		Order:       1000,
	},
}

func webAccelListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{Name: "Origin"},
		{
			Name:    "SentData(LastWeek/GiB)",
			Sources: []string{"GibSentInLastWeek"},
		},
		{
			Name: "Domain",
			FormatFunc: func(values map[string]string) string {
				switch values["DomainType"] {
				case "own_domain":
					return values["Domain"]
				case "subdomain":
					return values["Subdomain"]
				}
				return ""
			},
		},
	}
}

func webAccelDeleteCacheColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Result"},
		{Name: "Status"},
		{Name: "URL"},
	}
}

func webAccelDetailIncludes() []string {
	return []string{}
}

func webAccelDetailExcludes() []string {
	return []string{}
}

func webAccelCertIncludes() []string {
	return []string{}
}

func webAccelCertExcludes() []string {
	return []string{
		"Current.CertificateChain",
		"Old.0.CertificateChain",
		"Old.1.CertificateChain",
		"Old.2.CertificateChain",
		"Old.3.CertificateChain",
		"Old.4.CertificateChain",
		"Current.SHA256Fingerprint",
		"Old.0.SHA256Fingerprint",
		"Old.1.SHA256Fingerprint",
		"Old.2.SHA256Fingerprint",
		"Old.3.SHA256Fingerprint",
		"Old.4.SHA256Fingerprint",
		"Certificate.Current.CertificateChain",
		"Certificate.Old.0.CertificateChain",
		"Certificate.Old.1.CertificateChain",
		"Certificate.Old.2.CertificateChain",
		"Certificate.Old.3.CertificateChain",
		"Certificate.Old.4.CertificateChain",
		"Certificate.Current.SHA256Fingerprint",
		"Certificate.Old.0.SHA256Fingerprint",
		"Certificate.Old.1.SHA256Fingerprint",
		"Certificate.Old.2.SHA256Fingerprint",
		"Certificate.Old.3.SHA256Fingerprint",
		"Certificate.Old.4.SHA256Fingerprint",
	}
}

func webAccelListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func webAccelReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func webAccelCertInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func webAccelCertCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cert": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set certificate(from file)",
			ValidateFunc: validateFileExists(),
			Category:     "cert",
			Order:        10,
		},
		"key": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set private key(from file)",
			ValidateFunc: validateFileExists(),
			Category:     "cert",
			Order:        20,
		},
		"cert-content": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Description:   "set certificate(from text)",
			ConflictsWith: []string{"cert"},
			Category:      "cert",
			Order:         30,
		},
		"key-content": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Description:   "set private key(from text)",
			ConflictsWith: []string{"key"},
			Category:      "cert",
			Order:         40,
		},
	}
}

func webAccelCertUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cert": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set certificate(from file)",
			ValidateFunc: validateFileExists(),
			Category:     "cert",
			Order:        10,
		},
		"key": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set private key(from file)",
			ValidateFunc: validateFileExists(),
			Category:     "cert",
			Order:        20,
		},
		"cert-content": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Description:   "set certificate(from text)",
			ConflictsWith: []string{"cert"},
			Category:      "cert",
			Order:         30,
		},
		"key-content": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Description:   "set private key(from text)",
			ConflictsWith: []string{"key"},
			Category:      "cert",
			Order:         40,
		},
	}
}

func webAccelDeleteCacheParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}
