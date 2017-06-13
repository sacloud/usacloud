package completion

import (
	"fmt"
	"gopkg.in/urfave/cli.v2"
	"strconv"
)

func FlagNames(c *cli.Context, commandName string) {
	comm := c.App.Command(commandName)
	if comm == nil {
		return
	}
	for _, f := range comm.VisibleFlags() {
		for _, n := range f.Names() {
			format := "--%s\n"
			if len(n) == 1 {
				format = "-%s\n"
			}
			fmt.Printf(format, n)
		}
	}
}

func isSakuraID(id string) bool {
	_, err := strconv.ParseInt(id, 10, 64)
	return err == nil
}
