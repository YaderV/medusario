package data

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type Location struct {
	ID          int64  `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Address     string `json:"address" validate:"required"`
}

var validate *validator.Validate

func (l Location) Validate() error {
	validate = validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(l)
}

type LocationModel struct {
	DB *sql.DB
}

func (l LocationModel) Insert(loc *Location) error {
	return nil
}

type MockLocationModel struct{}

func (m MockLocationModel) Insert(loc *Location) error {
	return nil
}
