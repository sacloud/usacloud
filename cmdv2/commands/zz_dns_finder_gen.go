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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-resource-finder'; DO NOT EDIT

package commands

import (
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/sacloud/usacloud/pkg/utils"
)

func findDNSRecordInfoTargets(ctx Context, param *params.RecordInfoDNSParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().DNS

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.CommonServiceDNSItems {
			if hasTags(&v, param.Selector) {
				ids = append(ids, v.GetID())
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := sacloud.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					apiClient.Reset()
					apiClient.SetFilterBy("Name", idOrName)
					res, err := apiClient.Find()
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.CommonServiceDNSItems {
						if len(param.Selector) == 0 || hasTags(&v, param.Selector) {
							ids = append(ids, v.GetID())
						}
					}
				}
			}

		}

	}

	ids = utils.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}
	if len(ids) != 1 {
		return ids, fmt.Errorf("could not run with multiple targets: %v", ids)
	}

	return ids, nil
}

func findDNSRecordBulkUpdateTargets(ctx Context, param *params.RecordBulkUpdateDNSParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().DNS

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.CommonServiceDNSItems {
			if hasTags(&v, param.Selector) {
				ids = append(ids, v.GetID())
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := sacloud.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					apiClient.Reset()
					apiClient.SetFilterBy("Name", idOrName)
					res, err := apiClient.Find()
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.CommonServiceDNSItems {
						if len(param.Selector) == 0 || hasTags(&v, param.Selector) {
							ids = append(ids, v.GetID())
						}
					}
				}
			}

		}

	}

	ids = utils.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}
	if len(ids) != 1 {
		return ids, fmt.Errorf("could not run with multiple targets: %v", ids)
	}

	return ids, nil
}

func findDNSRecordAddTargets(ctx Context, param *params.RecordAddDNSParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().DNS

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.CommonServiceDNSItems {
			if hasTags(&v, param.Selector) {
				ids = append(ids, v.GetID())
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := sacloud.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					apiClient.Reset()
					apiClient.SetFilterBy("Name", idOrName)
					res, err := apiClient.Find()
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.CommonServiceDNSItems {
						if len(param.Selector) == 0 || hasTags(&v, param.Selector) {
							ids = append(ids, v.GetID())
						}
					}
				}
			}

		}

	}

	ids = utils.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}
	if len(ids) != 1 {
		return ids, fmt.Errorf("could not run with multiple targets: %v", ids)
	}

	return ids, nil
}

func findDNSReadTargets(ctx Context, param *params.ReadDNSParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().DNS

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.CommonServiceDNSItems {
			if hasTags(&v, param.Selector) {
				ids = append(ids, v.GetID())
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := sacloud.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					apiClient.Reset()
					apiClient.SetFilterBy("Name", idOrName)
					res, err := apiClient.Find()
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.CommonServiceDNSItems {
						if len(param.Selector) == 0 || hasTags(&v, param.Selector) {
							ids = append(ids, v.GetID())
						}
					}
				}
			}

		}

	}

	ids = utils.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}
	if len(ids) != 1 {
		return ids, fmt.Errorf("could not run with multiple targets: %v", ids)
	}

	return ids, nil
}

func findDNSRecordUpdateTargets(ctx Context, param *params.RecordUpdateDNSParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().DNS

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.CommonServiceDNSItems {
			if hasTags(&v, param.Selector) {
				ids = append(ids, v.GetID())
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := sacloud.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					apiClient.Reset()
					apiClient.SetFilterBy("Name", idOrName)
					res, err := apiClient.Find()
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.CommonServiceDNSItems {
						if len(param.Selector) == 0 || hasTags(&v, param.Selector) {
							ids = append(ids, v.GetID())
						}
					}
				}
			}

		}

	}

	ids = utils.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}
	if len(ids) != 1 {
		return ids, fmt.Errorf("could not run with multiple targets: %v", ids)
	}

	return ids, nil
}

func findDNSRecordDeleteTargets(ctx Context, param *params.RecordDeleteDNSParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().DNS

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.CommonServiceDNSItems {
			if hasTags(&v, param.Selector) {
				ids = append(ids, v.GetID())
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := sacloud.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					apiClient.Reset()
					apiClient.SetFilterBy("Name", idOrName)
					res, err := apiClient.Find()
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.CommonServiceDNSItems {
						if len(param.Selector) == 0 || hasTags(&v, param.Selector) {
							ids = append(ids, v.GetID())
						}
					}
				}
			}

		}

	}

	ids = utils.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}
	if len(ids) != 1 {
		return ids, fmt.Errorf("could not run with multiple targets: %v", ids)
	}

	return ids, nil
}

func findDNSUpdateTargets(ctx Context, param *params.UpdateDNSParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().DNS

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.CommonServiceDNSItems {
			if hasTags(&v, param.Selector) {
				ids = append(ids, v.GetID())
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := sacloud.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					apiClient.Reset()
					apiClient.SetFilterBy("Name", idOrName)
					res, err := apiClient.Find()
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.CommonServiceDNSItems {
						if len(param.Selector) == 0 || hasTags(&v, param.Selector) {
							ids = append(ids, v.GetID())
						}
					}
				}
			}

		}

	}

	ids = utils.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}

	return ids, nil
}

func findDNSDeleteTargets(ctx Context, param *params.DeleteDNSParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().DNS

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.CommonServiceDNSItems {
			if hasTags(&v, param.Selector) {
				ids = append(ids, v.GetID())
			}
		}
		if len(ids) == 0 {
			return ids, fmt.Errorf("finding resource id is failed: not found with search param [tags=%s]", param.Selector)
		}
	} else {
		for _, arg := range args {
			for _, a := range strings.Split(arg, "\n") {
				idOrName := a
				if id := sacloud.StringID(idOrName); !id.IsEmpty() {
					ids = append(ids, id)
				} else {
					apiClient.Reset()
					apiClient.SetFilterBy("Name", idOrName)
					res, err := apiClient.Find()
					if err != nil {
						return ids, fmt.Errorf("finding resource id is failed: %s", err)
					}
					if res.Count == 0 {
						return ids, fmt.Errorf("finding resource id is failed: not found with search param [%q]", idOrName)
					}
					for _, v := range res.CommonServiceDNSItems {
						if len(param.Selector) == 0 || hasTags(&v, param.Selector) {
							ids = append(ids, v.GetID())
						}
					}
				}
			}

		}

	}

	ids = utils.UniqIDs(ids)
	if len(ids) == 0 {
		return ids, fmt.Errorf("finding resource is is failed: not found")
	}

	return ids, nil
}
