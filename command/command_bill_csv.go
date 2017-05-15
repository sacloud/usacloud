package command

import (
	"fmt"
	"os"
	"strings"
)

func BillCsv(ctx Context, params *CsvBillParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetBillAPI()

	// auth-status
	auth, err := client.GetAuthStatusAPI().Read()
	if err != nil {
		return fmt.Errorf("BillCsv is failed: %s", err)
	}
	memberCD := auth.Member.Code

	if !strings.Contains(auth.ExternalPermission, "bill") {
		return fmt.Errorf("Don't have permission to view bills")
	}

	// call Find()
	res, err := finder.GetDetailCSV(memberCD, params.Id)
	if err != nil {
		return fmt.Errorf("BillCsv is failed: %s", err)
	}

	var out = GlobalOption.Out
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
