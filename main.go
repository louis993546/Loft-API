package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// TimeResponse contains time in string + format that the time string is in.
type TimeResponse struct {
	Time   string `json:"time"`
	Format string `json:"format"`
}

func main() {
	port := os.Getenv("PORT")

	fmt.Printf("localhost:%s should be up\n", port)

	//setup
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	//routing
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Serve up a swagger UI")
	})
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
	error := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
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
