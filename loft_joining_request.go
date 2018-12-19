package main

import (
	"fmt"
	"net/http"
)

// PostLoftJoiningRequests creates a request for a user to join a loft
func PostLoftJoiningRequests(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST loft joining requests seems to be working")
}

// GetLoftJoiningRequests returns the list of requests waiting for approval by any of that loft's member
func GetLoftJoiningRequests(w http.ResponseWriter, r *http.Request) {

}
