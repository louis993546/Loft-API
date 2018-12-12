package main

import (
	"fmt"
	"net/http"
)

func PostNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST notes seems to be working")
}

func GetNotes(w http.ResponseWriter, r *http.Request) {

}
