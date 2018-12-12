package main

import (
	"fmt"
	"net/http"
)

func PostLofts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST lofts seems to be working")
}
