package funcs

import (
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func BillList(ctx command.Context, params *params.ListBillParam) error {
	client := ctx.GetAPIClient()
	finder := client.GetBillAPI()

	// validate
	if params.Month > 0 && params.Year == 0 {
		return fmt.Errorf("year is required when month is setted")
	}

	// auth-status
	auth, err := client.GetAuthStatusAPI().Read()
	if err != nil {
		return fmt.Errorf("BillList is failed: %s", err)
	}
	accountID := auth.Account.ID
	if accountID == 0 {
		return fmt.Errorf("BillList is failed: %s", "invalid account id")
	}

	if !strings.Contains(auth.ExternalPermission, "bill") {
		return fmt.Errorf("Don't have permission to view bills")
	}

	var findFunc func() (*api.BillResponse, error)
	// call API
	switch {
	case params.Month > 0:
		findFunc = func() (*api.BillResponse, error) {
			return finder.ByContractYearMonth(accountID, params.Year, params.Month)
		}
	case params.Year > 0:
		findFunc = func() (*api.BillResponse, error) {
			return finder.ByContractYear(accountID, params.Year)
		}
	default:
		findFunc = func() (*api.BillResponse, error) {
			return finder.ByContract(accountID)
		}
	}

	// call Find()
	res, err := findFunc()
	if err != nil {
		return fmt.Errorf("BillList is failed: %s", err)
	}

	list := []interface{}{}
	for i := range res.Bills {
		list = append(list, res.Bills[i])
	}

	return ctx.GetOutput().Print(list...)
}
