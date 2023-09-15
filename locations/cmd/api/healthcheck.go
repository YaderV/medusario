package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	js := envelope{
		"status":     "available",
		"enviroment": app.config.env,
		"version":    version,
	}
	err := app.writeJSON(w, http.StatusOK, js, nil)
	if err != nil {

	}
}
