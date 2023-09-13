package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yaderv/medusario/internal/data"
)

func TestLocations(t *testing.T) {
	assert := require.New(t)
	app := &application{
		config: config{env: "testing"},
		logger: log.New(io.Discard, "", 0),
		models: data.NewMockModels(),
	}
	srv := httptest.NewServer(app.routes())
	defer srv.Close()

	t.Run("healthcheck", func(t *testing.T) {
		expected := map[string]string{
			"enviroment": "testing",
			"status":     "available",
			"version":    "1.0.0",
		}
		var output map[string]string

		res, err := get(srv, "/v1/healthcheck", &output)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(http.StatusOK, res.StatusCode)
		assert.Equal(expected, output)
	})

	t.Run("create location handler", func(t *testing.T) {
		jsonData := `{
			"title": "Portales Nte",
			"description": "Nice appartment",
			"address": "Test Address"
		}`
		expected := map[string]data.Location{
			"data": {
				Title:       "Portales Nte",
				Description: "Nice appartment",
				Address:     "Test Address",
			},
		}
		var output map[string]data.Location
		res, err := post(srv, "/v1/locations", jsonData, &output)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(http.StatusCreated, res.StatusCode)
		assert.Equal(expected, output)
	})

	t.Run("create location handler field validation", func(t *testing.T) {
		jsonData := `{
			"description": "Nice appartment",
			"address": "Test Address"
		}`
		expected := map[string]map[string]string{
			"error": {"Title": "required"},
		}
		var output map[string]map[string]string
		res, err := post(srv, "/v1/locations", jsonData, &output)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(http.StatusBadRequest, res.StatusCode)
		assert.Equal(expected, output)
	})

	t.Run("create location handler json empty", func(t *testing.T) {
		expected := map[string]string{
			"error": "body must not be empty",
		}
		var output map[string]string
		res, err := post(srv, "/v1/locations", "", &output)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(http.StatusBadRequest, res.StatusCode)
		assert.Equal(expected, output)
	})

	t.Run("create location handler wrong json format", func(t *testing.T) {
		jsonData := `{
			"title": "Portales Nte,
			"description": "Nice appartment",
			"address": "Test Address"
		}`
		expected := map[string]string{
			"error": "body contains badly-formed JSON (at character 29)",
		}
		var output map[string]string
		res, err := post(srv, "/v1/locations", jsonData, &output)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(http.StatusBadRequest, res.StatusCode)
		assert.Equal(expected, output)
	})

	t.Run("create location handler unexpected json field", func(t *testing.T) {
		jsonData := `{
			"title": "Portales Nte",
			"description": "Nice appartment",
			"address": "Test Address",
			"color": "blue"
		}`
		expected := map[string]string{
			"error": "body contains unknown key \"color\"",
		}
		var output map[string]string
		res, err := post(srv, "/v1/locations", jsonData, &output)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(http.StatusBadRequest, res.StatusCode)
		assert.Equal(expected, output)
	})

	t.Run("show location list handler", func(t *testing.T) {
		expected := map[string][]data.Location{
			"data": {
				{ID: 1, Title: "Narvarte", Description: "Nice house, good rommies", Address: "Test Address"},
				{ID: 2, Title: "Las Colinas", Description: "Expensive house in Mga", Address: "Test Address"},
			},
		}
		var output map[string][]data.Location
		res, err := get(srv, "/v1/locations", &output)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(http.StatusOK, res.StatusCode)
		assert.Equal(expected, output)
	})
}

func get(srv *httptest.Server, url string, output any) (*http.Response, error) {
	res, err := srv.Client().Get(srv.URL + url)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(output)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func post(srv *httptest.Server, url, jsonData string, output any) (*http.Response, error) {
	res, err := srv.Client().Post(srv.URL+url, "application/json", strings.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&output)
	if err != nil {
		return nil, err
	}
	return res, nil
}
