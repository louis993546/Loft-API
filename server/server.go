package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	_ "github.com/lib/pq"
	"github.com/louistsaitszho/loft"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("PORT not found from env, will fallback to %s", defaultPort)
		port = defaultPort
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Panic("No DATABASE_URL from env")
	}

	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(loft.NewExecutableSchema(loft.Config{Resolvers: &loft.Resolver{}})))

	domain := os.Getenv("DOMAIN")
	if domain == "" {
		domain = "http://localhost"
	}

	log.Printf("connect to %s:%s/ for GraphQL playground", domain, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
