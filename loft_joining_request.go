package main

import (
	"fmt"
	"net/http"
)

func PostLoftJoiningRequests(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST loft joining requests seems to be working")
}

func GetLoftJoiningRequests(w http.ResponseWriter, r *http.Request) {

}
