package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (app *application) errorResponse(w http.ResponseWriter, status int, errors any) {
	env := envelope{"error": errors}
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) badRequestResponse(w http.ResponseWriter, msg string) {
	app.errorResponse(w, http.StatusBadRequest, msg)
}

func (app *application) failedValidationResponse(w http.ResponseWriter, err error) {
	errorMap := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		// NOTE: this should be improved and return the json tag name instead
		errorMap[err.Field()] = err.Tag()
	}
	app.errorResponse(w, http.StatusBadRequest, errorMap)
}

func (app *application) internalServerErrorResponse(w http.ResponseWriter, msg string) {
	app.errorResponse(w, http.StatusInternalServerError, msg)
}
