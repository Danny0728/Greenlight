package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record Not found")
)

type Models struct {
	Movies MovieModel
}

//returns the Models struct containing the initialized MovieModel
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{
			DB: db,
		},
	}
}
