package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/louistsaitszho/loft/database"

	"github.com/99designs/gqlgen/handler"
	_ "github.com/lib/pq"
	"github.com/louistsaitszho/loft"
)

const defaultPort = "8080"
const compatibleDatabaseSchemaVersion = 0

func main() {
	// Setup DB
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatalln("No DATABASE_URL from env")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Cannot connect to DB: %s", err.Error())
	}

	schemaVersion, err := database.GetSchemaVersion(db)
	if err != nil {
		switch err.(type) {
		case *database.ErrorNotFound:
			log.Println("Can't find schema version, there is probably no db here in the first place. Will initialize it now...")
			initErr := database.InitializeDatabase(db)
			if initErr != nil {
				log.Fatalf("Failed to init database: %v", initErr)
			}
		case *database.ErrorCorrupted:
			panic("not implemented")
		default:
			log.Panicf("not implemented: '%v'\n", err)
		}
	} else {
		switch {
		case schemaVersion < compatibleDatabaseSchemaVersion:
			database.PerformDatabaseMigration(schemaVersion, compatibleDatabaseSchemaVersion)
		case schemaVersion == compatibleDatabaseSchemaVersion:
			log.Println("Code & Database schema match 👍")
		case schemaVersion > compatibleDatabaseSchemaVersion:
			log.Fatalf("Database v%d is ahead of the code!? It should not go higher then %d", schemaVersion, compatibleDatabaseSchemaVersion)
		}
	}

	// Setup GraphQL
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	resolver := loft.NewResolver(db)
	http.Handle("/query", handler.GraphQL(loft.NewExecutableSchema(loft.Config{Resolvers: resolver})))

	// Setup network
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		domain = "http://localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("PORT not found from env, will fallback to %s", defaultPort)
		port = defaultPort
	}

	log.Printf("connect to %s:%s/ for GraphQL playground", domain, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
