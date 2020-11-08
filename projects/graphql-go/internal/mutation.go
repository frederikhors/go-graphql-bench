package internal

import (
	"github.com/graphql-go/graphql"
	"go_graphql_bench/domain"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type: ProductType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"info": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				product := domain.Product{
					ID:    int32(len(domain.Products) + 1),
					Name:  params.Args["name"].(string),
					Info:  params.Args["info"].(string),
					Price: params.Args["price"].(int32),
				}
				domain.Products = append(domain.Products, product)
				return product, nil
			},
		},

		"update": &graphql.Field{
			Type: ProductType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"info": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				name, nameOk := params.Args["name"].(string)
				info, infoOk := params.Args["info"].(string)
				price, priceOk := params.Args["price"].(int32)
				for i, p := range domain.Products {
					if int32(id) == p.ID {
						if nameOk {
							domain.Products[i].Name = name
						}
						if infoOk {
							domain.Products[i].Info = info
						}
						if priceOk {
							domain.Products[i].Price = price
						}
						return domain.Products[i], nil
					}
				}
				return nil, nil
			},
		},

		"delete": &graphql.Field{
			Type: ProductType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				product := domain.Product{}
				for i, p := range domain.Products {
					if int32(id) == p.ID {
						product = domain.Products[i]
						domain.Products = append(domain.Products[:i], domain.Products[i+1:]...)
					}
				}
				return product, nil
			},
		},
	},
})
