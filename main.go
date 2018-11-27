package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func responseError404LoftNotFound() []ErrorResponseObject {
	return []ErrorResponseObject{
		ErrorResponseObject{
			Status: "200",
			Title:  "Test",
			Code:   "404_NOTHING_NOT_FOUND",
		},
	}
}

func main() {
	fmt.Println("starting")

	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Get something")
		error := responseError404LoftNotFound()

		json.NewEncoder(w).Encode(error)
	})
	error := http.ListenAndServe(":8080", r)
	if error != nil {
		fmt.Println("Something went wrong")
	}

	fmt.Println("ending")
}

// commonMiddleware adds a bunch of commonly use stuff into the router
// - Content-Type: application/vnd.api+json
// - [TODO] logger
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//logging
		log.Println(r.URL)

		//content type
		w.Header().Add("Content-Type", "application/vnd.api+json")

		//ends
		next.ServeHTTP(w, r)
	})
}
