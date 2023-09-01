package data

import "database/sql"

type Models struct {
	Locations interface {
		Insert(loc *Location) error
	}
}

func NewModels(db *sql.DB) *Models {
	return &Models{
		Locations: LocationModel{DB: db},
	}
}
