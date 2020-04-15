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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-resource-finder'; DO NOT EDIT

package cli

import (
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/pkg/utils"
)

func findLoadBalancerReadTargets(ctx Context, param *params.ReadLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerUpdateTargets(ctx Context, param *params.UpdateLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerDeleteTargets(ctx Context, param *params.DeleteLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerBootTargets(ctx Context, param *params.BootLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerShutdownTargets(ctx Context, param *params.ShutdownLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerShutdownForceTargets(ctx Context, param *params.ShutdownForceLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerResetTargets(ctx Context, param *params.ResetLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerWaitForBootTargets(ctx Context, param *params.WaitForBootLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerWaitForDownTargets(ctx Context, param *params.WaitForDownLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerVipInfoTargets(ctx Context, param *params.VipInfoLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerVipAddTargets(ctx Context, param *params.VipAddLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerVipUpdateTargets(ctx Context, param *params.VipUpdateLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerVipDeleteTargets(ctx Context, param *params.VipDeleteLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerServerInfoTargets(ctx Context, param *params.ServerInfoLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerServerAddTargets(ctx Context, param *params.ServerAddLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerServerUpdateTargets(ctx Context, param *params.ServerUpdateLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerServerDeleteTargets(ctx Context, param *params.ServerDeleteLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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

func findLoadBalancerMonitorTargets(ctx Context, param *params.MonitorLoadBalancerParam) ([]sacloud.ID, error) {
	var ids []sacloud.ID
	args := ctx.Args()
	apiClient := ctx.GetAPIClient().LoadBalancer

	if len(args) == 0 {
		if len(param.Selector) == 0 {
			return ids, fmt.Errorf("ID or Name argument or --selector option is required")
		}
		apiClient.Reset()
		res, err := apiClient.Find()
		if err != nil {
			return ids, fmt.Errorf("finding resource id is failed: %s", err)
		}
		for _, v := range res.LoadBalancers {
			if utils.HasTags(&v, param.Selector) {
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
					for _, v := range res.LoadBalancers {
						if len(param.Selector) == 0 || utils.HasTags(&v, param.Selector) {
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