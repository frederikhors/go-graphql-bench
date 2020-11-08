package main

import (
	"fmt"
	"go_graphql_bench/projects/gophers/internal"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(internal.Schema, &internal.Resolver{}, opts...)

	http.Handle("/graphql", &relay.Handler{Schema: schema})

	fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
