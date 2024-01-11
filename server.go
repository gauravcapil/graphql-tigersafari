package main

import (
	"log"
	"net/http"
	"os"

	"gaurav.kapil/tigerhall/dbutils"
	"gaurav.kapil/tigerhall/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	err := dbutils.Inititialize()
	if err != nil {
		log.Fatalf("DB connection failed to initialize: %s", err.Error())
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static", fs)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	mux.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/playgrounds for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
