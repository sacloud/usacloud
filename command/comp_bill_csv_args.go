package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"strings"
)

func BillCsvCompleteArgs(ctx Context, params *CsvBillParam, cur, prev, commandName string) {

	if !GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetBillAPI()

	// auth-status
	auth, err := client.GetAuthStatusAPI().Read()
	if err != nil {
		return
	}
	accountID := sacloud.NewResourceByStringID(auth.Account.ID).ID
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
