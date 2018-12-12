package main

import (
	"fmt"
	"net/http"
)

func PostTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST tasks seems to be working")
}

func GetTasks(w http.ResponseWriter, r *http.Request) {

}
