package completion

import (
	"fmt"
	"io"

	"github.com/sacloud/usacloud/command/profile"
)

func writeAllProfileName(w io.Writer) {
	profiles, err := profile.List()
	if err != nil {
		return
	}
	for _, p := range profiles {
		fmt.Fprintln(w, p)
	}
}
