package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record Not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies MovieModel
	Users  UserModel
}

//returns the Models struct containing the initialized MovieModel
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
		Users:  UserModel{DB: db},
	}
}
