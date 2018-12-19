package main

import (
	"fmt"
	"net/http"
)

// PostLofts creates a new loft and the first user
func PostLofts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST lofts seems to be working")
}
