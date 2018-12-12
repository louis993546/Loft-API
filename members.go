package main

import (
	"fmt"
	"net/http"
)

func PostMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST members seems to be working")
}
