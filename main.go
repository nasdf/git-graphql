package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nasdf/git-graphql/graph"
	"github.com/nasdf/git-graphql/graph/generated"
)

const bindAddr = ":8080"

func main() {
	res, err := graph.NewResolver()
	if err != nil {
		log.Fatal(err)
	}

	cfg := generated.Config{Resolvers: res}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost%s/ for GraphQL playground", bindAddr)
	log.Fatal(http.ListenAndServe(bindAddr, nil))
}
