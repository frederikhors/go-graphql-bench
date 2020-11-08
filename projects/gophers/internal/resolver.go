package internal

import "go_graphql_bench/domain"

type Resolver struct{}

func (r *Resolver) Product(args struct{ ID int32 }) *domain.Product {
	for _, product := range domain.Products {
		if product.ID == args.ID {
			return &product
		}
	}
	return nil
}

func (r *Resolver) List() []domain.Product {
	return domain.Products
}

func (r *Resolver) Create(args *struct {
	Name  string
	Info  string
	Price int32
}) *domain.Product {
	product := domain.Product{
		ID:    int32(len(domain.Products) + 1),
		Name:  args.Name,
		Info:  args.Info,
		Price: args.Price,
	}
	domain.Products = append(domain.Products, product)
	return &product
}

func (r *Resolver) Update(args *struct {
	ID    int32
	Name  string
	Info  string
	Price int32
}) *domain.Product {
	for i, p := range domain.Products {
		if args.ID == p.ID {
			domain.Products[i].Name = args.Name
			domain.Products[i].Info = args.Info
			domain.Products[i].Price = args.Price
			return &domain.Products[i]
		}
	}
	return nil
}

func (r *Resolver) Delete(args *struct{ ID int32 }) *domain.Product {
	var product domain.Product
	for i, p := range domain.Products {
		if args.ID == p.ID {
			product = domain.Products[i]
			domain.Products = append(domain.Products[:i], domain.Products[i+1:]...)
		}
	}
	return &product
}
