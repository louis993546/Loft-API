package main

import (
	"fmt"
	"net/http"
)

func PostEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST events seems to be working")
}
