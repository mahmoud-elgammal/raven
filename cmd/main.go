package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mahmoud-elgammal/raven/graph"
	"github.com/mahmoud-elgammal/raven/pkg/database"
	"github.com/mahmoud-elgammal/raven/pkg/resolvers"
)

const (
	name        = "Raven"
	version     = "0.0.1"
	author      = "Mahmoud S. ElGammal"
	year        = 2022
	defaultPort = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	deriver, err := database.GetDriver("neo4j://localhost:7687", "neo4j", "password")

	if err != nil {
		log.Fatal(err)
	}

	defer database.CloseDriver(deriver)

	if err = database.ValidateDriver(deriver); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Neo4j driver is ready")
	}

	session, err := database.GetSession(deriver)

	if err != nil {
		log.Fatal(err)
	}

	defer database.CloseSession(session)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
