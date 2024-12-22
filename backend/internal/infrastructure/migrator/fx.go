package migrator

import "go.uber.org/fx"

var migrateModule = fx.Provide(
	fx.Annotate(
		NewMigrate,
		fx.OnStart(OnStart),
		fx.OnStop(OnStop),
	),
)

var Module = fx.Module(
	"migrate",
	migrateModule,
)
