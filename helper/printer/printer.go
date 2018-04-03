package printer

import (
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/sacloud/usacloud/command"
)

// Print is delegates to *color.Color or fmt depending on command.GlobalOption.NoColor flag
func Print(c *color.Color, a ...interface{}) {
	Fprint(command.GlobalOption.Out, c, a)
}

// Println is delegates to *color.Color or fmt depending on command.GlobalOption.NoColor flag
func Println(c *color.Color, a ...interface{}) {
	Fprintln(command.GlobalOption.Out, c, a)
}

// Printf is delegates to *color.Color or fmt depending on command.GlobalOption.NoColor flag
func Printf(c *color.Color, format string, a ...interface{}) {
	Fprintf(command.GlobalOption.Out, c, format, a)
}

// Fprint is delegates to *color.Color or fmt depending on command.GlobalOption.NoColor flag
func Fprint(w io.Writer, c *color.Color, a ...interface{}) {
	if command.GlobalOption.NoColor {
		fmt.Fprint(w, a...)
	} else {
		c.Fprint(w, a...)
	}
}

// Fprintln is delegates to *color.Color or fmt depending on command.GlobalOption.NoColor flag
func Fprintln(w io.Writer, c *color.Color, a ...interface{}) {
	if command.GlobalOption.NoColor {
		fmt.Fprintln(w, a...)
	} else {
		c.Fprintln(w, a...)
	}
}

// Fprintf is delegates to *color.Color or fmt depending on command.GlobalOption.NoColor flag
func Fprintf(w io.Writer, c *color.Color, format string, a ...interface{}) {
	if command.GlobalOption.NoColor {
		fmt.Fprintf(w, format, a...)
	} else {
		c.Fprintf(w, format, a...)
	}
}
