package database

import (
	"database/sql"

	"boons-drone.com/app/internal/domain/model"
)

type BoilUserRepository struct {
	db *sql.DB
}

func GetAll() ([]model.User, error) {

}
