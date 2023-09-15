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
	app.models.Locations.Insert(&loc)
	err = app.writeJSON(w, http.StatusCreated, envelope{"data": loc}, nil)
	if err != nil {
		app.badRequestResponse(w, err.Error())
		return
	}
}

func (app *application) showLocationListHandler(w http.ResponseWriter, r *http.Request) {
	locs, err := app.models.Locations.SelectAll()
	if err != nil {
		app.internalServerErrorResponse(w, err.Error())
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"data": locs}, nil)
}
