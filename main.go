package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// ErrorResponseObject is a representation of a generic error object in this API, following JSONAPI specification
type ErrorResponseObject struct {
	Status string `json:"status"`
	Title  string `json:"title"`
	Code   string `json:"code"`
}

// ErrorResponse is what JSONAPI wants to see when application wants to return error
type ErrorResponse struct {
	Errors []ErrorResponseObject `json:"errors"`
}

// TimeResponse TODO: this is pretty self-explainatory so what should i write?
type TimeResponse struct {
	Time   string `json:"time"`
	Format string `json:"format"`
}

func responseError404LoftNotFound() []ErrorResponseObject {
	return []ErrorResponseObject{
		ErrorResponseObject{
			Status: "404",
			Title:  "Cannot find the loft you are specifying",
			Code:   "404_LOFT_NOT_FOUND",
		},
	}
}

func responseError404NoteNotFound() []ErrorResponseObject {
	return []ErrorResponseObject{
		ErrorResponseObject{
			Status: "404",
			Title:  "Cannot find the note you are specifying",
			Code:   "404_NOTE_NOT_FOUND",
		},
	}
}

func responseError404TaskNotFound() []ErrorResponseObject {
	return []ErrorResponseObject{
		ErrorResponseObject{
			Status: "404",
			Title:  "Cannot find the task you are specifying",
			Code:   "404_TASK_NOT_FOUND",
		},
	}
}

func responseError404EventNotFound() []ErrorResponseObject {
	return []ErrorResponseObject{
		ErrorResponseObject{
			Status: "404",
			Title:  "Cannot find the event you are specifying",
			Code:   "404_EVENT_NOT_FOUND",
		},
	}
}

func responseError404MemberNotFound() []ErrorResponseObject {
	return []ErrorResponseObject{
		ErrorResponseObject{
			Status: "404",
			Title:  "Cannot find the member you are specifying",
			Code:   "404_Member_NOT_FOUND",
		},
	}
}

func main() {
	port := 8080

	fmt.Printf("localhost:%d should be up\n", port)

	//setup
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	//routing
	r.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		timeNow := TimeResponse{time.Now().UTC().Format(time.RFC3339), "RFC3339"}
		json.NewEncoder(w).Encode(timeNow)
	}).Methods("GET", "PUT", "POST", "DELETE", "PATCH")

	r.HandleFunc("/lofts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "POST lofts seems to be working")
	}).Methods("POST")

	r.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "POST notes seems to be working")
	}).Methods("POST")

	r.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "POST tasks seems to be working")
	}).Methods("POST")

	r.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "POST events seems to be working")
	}).Methods("POST")

	r.HandleFunc("/members", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "POST members seems to be working")
	}).Methods("POST")

	//start listening
	error := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if error != nil {
		fmt.Println("Something went wrong")
	}
}

// commonMiddleware adds a bunch of commonly use stuff into the router
// - Content-Type: application/vnd.api+json
// - logger
// TODO: refactor this to take multiple middleware to avoid adding way too many stuff in here
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//logging
		// log.Println(r.URL)
		log.Printf("req = '%s'\n", r.URL.String())

		//content type
		w.Header().Add("Content-Type", "application/vnd.api+json")

		//ends
		next.ServeHTTP(w, r)
	})
}
