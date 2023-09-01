package main

import (
	"net/http"

	"github.com/yaderv/medusario/internal/data"
)

func (app *application) createLocationHandler(w http.ResponseWriter, r *http.Request) {
	loc := &data.Location{}
	app.models.Locations.Insert(loc)
	app.logger.Println("db insert")
}
