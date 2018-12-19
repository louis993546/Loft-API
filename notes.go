package main

import (
	"fmt"
	"net/http"
)

// PostNotes creates a note for the loft
func PostNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST notes seems to be working")
}

// GetNotes = GET /notes
// Get list of notes of a loft
func GetNotes(w http.ResponseWriter, r *http.Request) {

}
