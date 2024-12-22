package config

import "go.uber.org/fx"

type Config struct {
}

type params struct {
	fx.In
}

type result struct {
	fx.Out

	Config *Config
}

func NewConfig(params params) result {
	config := Config{}

	return result{Config: &config}
}
