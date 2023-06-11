package main

import (
	"hm2/sessions_stat_service/graph"
	"hm2/sessions_stat_service/graph/model"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Games: make([]model.Game, 0), Index: map[string]int{}, Mutex: sync.Mutex{}}}))

	http.Handle("/", playground.Handler("GraphQL", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
