//go:generate ../tools/bin/gen-input-models
//go:generate ../tools/bin/gen-cli-commands
//go:generate ../tools/bin/gen-command-funcs
package define

import "github.com/sacloud/usacloud/schema"

var Resources map[string]*schema.Resource = map[string]*schema.Resource{
	"Bridge":   BridgeResource(),
	"Icon":     IconResource(),
	"Internet": InternetResource(),
	"Switch":   SwitchResource(),
}
