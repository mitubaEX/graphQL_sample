package graphql_util

import "github.com/graphql-go/graphql"

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

func NewSchema() graphql.Schema {
	return schema
}
