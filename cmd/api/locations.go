package main

import (
	"net/http"

	"github.com/yaderv/medusario/internal/data"
)

func (app *application) createLocationHandler(w http.ResponseWriter, r *http.Request) {
	var loc data.Location
	err := app.readJSON(r, loc)
	if err != nil {
		app.errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	//app.models.Locations.Insert(&loc)
	err = app.writeJSON(w, http.StatusOK, loc, nil)
	if err != nil {
		app.errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
}
