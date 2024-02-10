package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"gaurav.kapil/graphql-tigersafari/auth"
	"gaurav.kapil/graphql-tigersafari/dbutils"
	"gaurav.kapil/graphql-tigersafari/emailserver"
	"gaurav.kapil/graphql-tigersafari/graph"
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

	// Start email server
	emailserver.Initialize()
	emailserver.StartEmailServer()
	<-emailserver.IsReady

	router := chi.NewRouter()

	// Adding auth middleware to inject the context of the header to be used furthers
	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	// Playground for checking API calls
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	// photo server for sharing evidence
	router.Handle("/photos/*", http.StripPrefix("/photos/", http.FileServer(http.Dir(dbutils.GetPhotoDir()))))

	log.Printf("connect to http://localhost:%s/playgrounds for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
