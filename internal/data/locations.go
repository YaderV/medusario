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
	stmt := `
		INSERT INTO locations (title, description, address) 
		VALUES ($1, $2, $3) 
		RETURNING id`
	args := []any{loc.Title, loc.Description, loc.Address}
	return l.DB.QueryRow(stmt, args...).Scan(&loc.ID)
}

func (l LocationModel) SelectAll() ([]Location, error) {
	locs := make([]Location, 0)
	stmt := `SELECT id, title, description, address FROM locations`
	rows, err := l.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var loc Location
		err := rows.Scan(&loc.ID, &loc.Title, &loc.Description, &loc.Address)
		if err != nil {
			return nil, err
		}
		locs = append(locs, loc)

	}
	return locs, nil
}

type MockLocationModel struct{}

func (m MockLocationModel) Insert(loc *Location) error {
	return nil
}

func (m MockLocationModel) SelectAll() ([]Location, error) {
	return nil, nil
}
