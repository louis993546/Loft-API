package main

import (
	"fmt"
	"net/http"
)

// PostMembers creates a new member for the loft
func PostMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST members seems to be working")
}

// GetMembers returns a list of members for the loft
func GetMembers(w http.ResponseWriter, r *http.Request) {

}
