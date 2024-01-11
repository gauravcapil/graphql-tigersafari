package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"gaurav.kapil/tigerhall/auth"
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

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.Handle("/photos/*", http.StripPrefix("/photos/", http.FileServer(http.Dir(dbutils.GetPhotoDir()))))

	log.Printf("connect to http://localhost:%s/playgrounds for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
