package usecase

import (
	"boons-drone.com/app/internal/domain/model"
	"boons-drone.com/app/internal/usecase/request"
)

type UserUseCase interface {
	GetUsers(req request.GetUsersRequest) ([]model.User, error)
	GetUser(req request.GetUserRequest) (*model.User, error)
	CreateUser(req request.CreateUserRequest) (*model.User, error)
	UpdateUser(req request.UpdateUserRequest) (*model.User, error)
	DeleteUser(req request.DeleteUserRequest) (int, error)
}
