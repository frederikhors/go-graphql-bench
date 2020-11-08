package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"go_graphql_bench/projects/graphql-go/internal"
	"net/http"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    internal.QueryType,
		Mutation: internal.MutationType,
	},
)

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		var p postData
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			w.WriteHeader(400)
			return
		}
		result := graphql.Do(graphql.Params{
			Context:        r.Context(),
			Schema:         schema,
			RequestString:  p.Query,
			VariableValues: p.Variables,
			OperationName:  p.Operation,
		})
		if err := json.NewEncoder(w).Encode(result); err != nil {
			fmt.Printf("could not write result to response: %s", err)
		}
	})

	fmt.Println("Server is running on port 3000")
	_ = http.ListenAndServe(":3000", nil)
}
