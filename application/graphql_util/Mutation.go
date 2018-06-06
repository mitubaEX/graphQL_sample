package graphql_util

import (
	"github.com/mitubaEX/graphQL_sample/application/graphql_util/fields"
	"github.com/graphql-go/graphql"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser":  fields.CreateUserField,
		"createEvent": fields.CreateEventField,
	},
})
