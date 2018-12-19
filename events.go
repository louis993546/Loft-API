package main

import (
	"fmt"
	"net/http"
)

// PostEvents creates the event for the loft, and return the event back
func PostEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST events seems to be working")
}

// GetEvents gets a list of events for the loft with pagination
func GetEvents(w http.ResponseWriter, r *http.Request) {

}
