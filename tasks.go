package main

import (
	"fmt"
	"net/http"
)

// PostTasks creats a task for the loft
func PostTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST tasks seems to be working")
}

// GetTasks returns a list of tasks for the loft with pagination
func GetTasks(w http.ResponseWriter, r *http.Request) {

}
