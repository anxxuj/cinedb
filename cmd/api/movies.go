package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/anxxuj/cinedb/internal/data"
)

// Create a new movie
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// Show the details of a specific movie
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Dummy data
	data := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Nigga Man",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
