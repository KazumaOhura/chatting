package handler

import (
	"net/http"

	"boons-drone.com/app/internal/interface/http/handler/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Handler interface {
	RegisterRoutes(engine *gin.Engine)
}

var userHandler = fx.Provide(
	fx.Annotate(
		routes.NewUserHandler,
		fx.As(new(http.Handler)),
	),
)

var Module = fx.Module(
	"handler",
	userHandler,
)
