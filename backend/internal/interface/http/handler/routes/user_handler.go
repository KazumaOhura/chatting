package routes

import (
	repo "boons-drone.com/app/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo repo.UserRepository
}

func (uh *UserHandler) handleGetUsers(c *gin.Context) {

}

func (uh *UserHandler) handleGetUser(c *gin.Context) {

}

func NewUserHandler(repo repo.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (uh *UserHandler) RegisterRoutes(engine *gin.Engine) {
	userGroup := engine.Group("/users")

	{
		userGroup.GET("/", uh.handleGetUsers)
		userGroup.GET("/:id", uh.handleGetUser)
	}
}
