package main

import "net/http"

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, msg string) {
	env := envelope{"error": msg}
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
