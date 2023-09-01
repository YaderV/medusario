package data

import (
	"database/sql"
	"fmt"
)

type Location struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

type LocationModel struct {
	DB *sql.DB
}

func (l LocationModel) Insert(loc *Location) error {
	fmt.Println("Inserting location")
	return nil
}

type MockLocationModel struct{}

func (m MockLocationModel) Insert(loc *Location) error {
	return nil
}
