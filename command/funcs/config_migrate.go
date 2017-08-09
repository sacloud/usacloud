package funcs

import (
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/migration"
)

func ConfigMigrate(ctx command.Context, params *params.MigrateConfigParam) error {
	return migration.MigrateConfig()
}
