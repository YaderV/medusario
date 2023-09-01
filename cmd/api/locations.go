package main

import (
	"encoding/json"
	"net/http"

	"github.com/yaderv/medusario/internal/data"
)

func (app *application) createLocationHandler(w http.ResponseWriter, r *http.Request) {
	var loc data.Location
	err := json.NewDecoder(r.Body).Decode(&loc)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}
	app.models.Locations.Insert(&loc)
}
