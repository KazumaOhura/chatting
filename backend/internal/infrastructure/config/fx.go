package config

import (
	"go.uber.org/fx"
)

var configModule = fx.Provide(
	NewConfig,
)

var Module = fx.Module(
	"config",
	configModule,
)
