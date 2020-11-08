package internal

var Schema = `
	schema {
	  query: Query
	  mutation: Mutation
	}
	type Mutation {
	  create(name: String!, info: String!, price: Int!): Product
	  delete(id: Int!): Product
	  update(id: Int!, name: String!, info: String!, price: Int!): Product
	}
	type Product {
	  id: Int!
	  info: String!
	  name: String!
	  price: Int!
	}
	type Query {
	  list: [Product!]!
	  product(id: Int!): Product
	}
`
