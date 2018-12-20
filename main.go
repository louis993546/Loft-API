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
	r.Use(LogRequestMiddleware)

	//routing
	r.HandleFunc("/", home).Methods("GET")

	r.HandleFunc("/echo", GetEcho).Methods("GET")

	//TODO: I need to think about the sign up process to decide what the endpoint should be
	// r.HandleFunc("/lofts", SignUp).Methods("POST")

	r.HandleFunc("/lofts/{loftID}/events", CreateEvent).Methods("POST")
	r.HandleFunc("/lofts/{loftID}/events", GetEvents).Methods("GET")

	r.HandleFunc("/lofts/{loftID}/task", CreateTask).Methods("POST")
	r.HandleFunc("/lofts/{loftID}/tasks", GetTasks).Methods("GET")

	r.HandleFunc("/lofts/{loftID}/notes", CreateNote).Methods("POST")
	r.HandleFunc("/lofts/{loftID}/notes", GetNotes).Methods("GET")

	r.HandleFunc("/lofts/{loftID}/members", GetMembers).Methods("GET")

	r.HandleFunc("/lofts/{loftID}/join_request", CreateJoinRequest).Methods("POST")
	//TODO: I need to think about the sign up process to decide what the "approve/deny join request" should be
	// r.HandleFunc("/lofts/{loftID}/join_request/{joinRequestID}/decision", SignUp).Methods("POST")

	//start listening
	fmt.Printf("localhost:%s should be up\n", port)
	error := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if error != nil {
		log.Fatal(error)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Serve up a swagger UI")
}
