package main

import (
	"fmt"
	"go_graphql_bench/projects/gqlgen/graph"
	"go_graphql_bench/projects/gqlgen/graph/generated"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
)

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/graphql", srv)

	fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
