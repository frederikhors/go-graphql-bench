package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go_graphql_bench/domain"
	"go_graphql_bench/projects/gqlgen/graph/generated"
)

func (r *mutationResolver) Create(_ context.Context, name string, info string, price int32) (*domain.Product, error) {
	product := domain.Product{
		ID:    int32(len(products) + 1),
		Name:  name,
		Info:  info,
		Price: price,
	}
	products = append(products, &product)
	return &product, nil
}

func (r *mutationResolver) Delete(_ context.Context, id int32) (*domain.Product, error) {
	var product domain.Product
	for i, p := range products {
		if id == p.ID {
			product = *products[i]
			products = append(products[:i], products[i+1:]...)
		}
	}
	return &product, nil
}

func (r *mutationResolver) Update(_ context.Context, id int32, name string, info string, price int32) (*domain.Product, error) {
	for i, p := range products {
		if id == p.ID {
			products[i].Name = name
			products[i].Info = info
			products[i].Price = price
			return products[i], nil
		}
	}
	return nil, nil
}

func (r *queryResolver) List(_ context.Context) ([]*domain.Product, error) {
	return products, nil
}

func (r *queryResolver) Product(_ context.Context, id int32) (*domain.Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
