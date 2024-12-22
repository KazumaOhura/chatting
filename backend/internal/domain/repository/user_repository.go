package repository

import "boons-drone.com/app/internal/domain/model"

type UserRepository interface {
	FindById(id int) (*model.User, error)
}
