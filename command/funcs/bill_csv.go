package funcs

import (
	"fmt"
	"os"
	"strings"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func BillCsv(ctx command.Context, params *params.CsvBillParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetBillAPI()

	// auth-status
	auth, err := client.GetAuthStatusAPI().Read()
	if err != nil {
		return fmt.Errorf("BillCsv is failed: %s", err)
	}
	accountID := auth.Account.ID
	if accountID == 0 {
		return fmt.Errorf("BillCsv is failed: %s", "invalid account id")
	}
	memberCD := auth.Member.Code

	if !strings.Contains(auth.ExternalPermission, "bill") {
		return fmt.Errorf("Don't have permission to view bills")
	}

	// validate param
	billID := params.BillId
	if !ctx.IsSet("bill-id") {
		bills, err := finder.ByContract(accountID)
		if err != nil {
			return fmt.Errorf("BillCsv is failed: %s", err)
		}
		if len(bills.Bills) == 0 {
			return fmt.Errorf("BillCsv is failed: Empty result")
		}
		billID = bills.Bills[0].BillID
	}

	// call Find()
	res, err := finder.GetDetailCSV(memberCD, billID)
	if err != nil {
		return fmt.Errorf("BillCsv is failed: %s", err)
	}

	var out = command.GlobalOption.Out
	if params.BillOutput != "" {
		file, err := os.Create(params.BillOutput)
		if err != nil {
			return fmt.Errorf("BillCsv is failed: %s", err)
		}
		out = file
		defer file.Close()
	}

	// write to out

	if !params.NoHeader {
		_, err := fmt.Fprintf(out, "\"%s\"\n", strings.Join(res.HeaderRow, "\",\""))
		if err != nil {
			return fmt.Errorf("BillCsv is failed: %s", err)
		}
	}

	for _, line := range res.BodyRows {
		_, err := fmt.Fprintf(out, "\"%s\"\n", strings.Join(line, "\",\""))
		if err != nil {
			return fmt.Errorf("BillCsv is failed: %s", err)
		}
	}

	return nil
}
