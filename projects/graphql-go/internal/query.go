package internal

import (
	"github.com/graphql-go/graphql"
	"go_graphql_bench/domain"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"product": &graphql.Field{
				Type:        ProductType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						for _, product := range domain.Products {
							if int(product.ID) == id {
								return product, nil
							}
						}
					}
					return nil, nil
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(ProductType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return domain.Products, nil
				},
			},
		},
	})