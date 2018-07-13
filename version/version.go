package version

import (
	"fmt"
	"runtime"
)

var (
	// Version app version
	Version = "0.0.0" // set on build time
	// Revision git commit short commithash
	Revision = "xxxxxx" // set on build time
)

// FullVersion return sackerel full version text
func FullVersion() string {
	return fmt.Sprintf("%s %s/%s, build %s", Version, runtime.GOOS, runtime.GOARCH, Revision)
}
