package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command/profile"
	"io"
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
