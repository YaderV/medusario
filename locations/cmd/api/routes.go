package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mvrilo/go-redoc"
)

var doc = redoc.Redoc{
	Title:       "Location API",
	Description: "Location API provides endpoints to manage house or appartment with renting space.",
	SpecFile:    "./swagger.yaml",
	SpecPath:    "/v1/swagger.yaml",
	DocsPath:    "/v1/docs",
}

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.Handler(http.MethodGet, "/v1/docs", doc.Handler())
	router.HandlerFunc(http.MethodGet, "/v1/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./swagger.yaml")
	})
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/locations", app.createLocationHandler)
	router.HandlerFunc(http.MethodGet, "/v1/locations", app.showLocationListHandler)
	return router
}
