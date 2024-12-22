package main

import (
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Invoke(
			func(server *http.Server, migrate *migrate.Migrate) {},
		),
	)

	app.Run()
}
