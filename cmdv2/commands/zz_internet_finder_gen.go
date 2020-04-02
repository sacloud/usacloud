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

func findInternetReadTargets(ctx Context, param *params.ReadInternetParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Internet

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Internet {
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
					for _, v := range res.Internet {
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

func findInternetUpdateTargets(ctx Context, param *params.UpdateInternetParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Internet

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Internet {
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
					for _, v := range res.Internet {
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

func findInternetDeleteTargets(ctx Context, param *params.DeleteInternetParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Internet

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Internet {
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
					for _, v := range res.Internet {
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

func findInternetUpdateBandwidthTargets(ctx Context, param *params.UpdateBandwidthInternetParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Internet

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Internet {
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
					for _, v := range res.Internet {
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

func findInternetSubnetInfoTargets(ctx Context, param *params.SubnetInfoInternetParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Internet

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Internet {
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
					for _, v := range res.Internet {
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

func findInternetSubnetAddTargets(ctx Context, param *params.SubnetAddInternetParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Internet

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Internet {
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
					for _, v := range res.Internet {
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

func findInternetSubnetDeleteTargets(ctx Context, param *params.SubnetDeleteInternetParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Internet

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Internet {
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
					for _, v := range res.Internet {
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

func findInternetSubnetUpdateTargets(ctx Context, param *params.SubnetUpdateInternetParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Internet

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Internet {
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
					for _, v := range res.Internet {
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

func findInternetIPv6InfoTargets(ctx Context, param *params.IPv6InfoInternetParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Internet

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Internet {
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
					for _, v := range res.Internet {
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

func findInternetIPv6EnableTargets(ctx Context, param *params.IPv6EnableInternetParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Internet

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Internet {
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
					for _, v := range res.Internet {
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

func findInternetIPv6DisableTargets(ctx Context, param *params.IPv6DisableInternetParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Internet

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Internet {
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
					for _, v := range res.Internet {
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

func findInternetMonitorTargets(ctx Context, param *params.MonitorInternetParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().Internet

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.Internet {
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
					for _, v := range res.Internet {
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
