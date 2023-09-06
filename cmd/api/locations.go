package main

import (
	"net/http"

	"github.com/yaderv/medusario/internal/data"
)

func (app *application) createLocationHandler(w http.ResponseWriter, r *http.Request) {
	var loc data.Location
	err := app.readJSON(w, r, &loc)
	if err != nil {
		app.badRequestResponse(w, err.Error())
		return
	}
	err = loc.Validate()
	if err != nil {
		app.failedValidationResponse(w, err)
		return
	}
	//app.models.Locations.Insert(&loc)
	err = app.writeJSON(w, http.StatusOK, loc, nil)
	if err != nil {
		app.badRequestResponse(w, err.Error())
		return
	}
}
