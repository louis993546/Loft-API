package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) <= 0 {
		log.Println("No PORT from env, will try to use :8080 instead")
		port = "8080" //default port, not garuntee to be avilable
	}

	dbConnectionString := os.Getenv("DATABASE_URL")
	_, err := sql.Open("postgres", dbConnectionString)
	//TODO: should I defer db.close?
	if err != nil {
		log.Fatal(err)
	}

	//setup
	r := mux.NewRouter()
	r.Use(SetJSONAPIContentTypeMiddleware, LogRequestMiddleware)

	//routing
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/echo", AllEcho).Methods("GET", "PUT", "POST", "DELETE", "PATCH")
	r.HandleFunc("/lofts", PostLofts).Methods("POST")
	r.HandleFunc("/notes", PostNotes).Methods("POST")
	r.HandleFunc("/tasks", PostTasks).Methods("POST")
	r.HandleFunc("/events", PostEvents).Methods("POST")

	r.HandleFunc("/members", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "POST members seems to be working")
	}).Methods("POST")

	//start listening
	fmt.Printf("localhost:%s should be up\n", port)
	error := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if error != nil {
		fmt.Println("Something went wrong")
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Serve up a swagger UI")
}
