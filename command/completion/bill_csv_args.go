package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"strings"
)

func BillCsvCompleteArgs(ctx command.Context, params *params.CsvBillParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetBillAPI()

	// auth-status
	auth, err := client.GetAuthStatusAPI().Read()
	if err != nil {
		return
	}
	accountID := auth.Account.ID
	if accountID == 0 {
		return
	}

	if !strings.Contains(auth.ExternalPermission, "bill") {
		return
	}

	// call Find()
	res, err := finder.ByContract(accountID)
	if err != nil {
		return
	}

	for i := range res.Bills {
		fmt.Println(res.Bills[i].BillID)
	}

}
