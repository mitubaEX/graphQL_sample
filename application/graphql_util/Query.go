package graphql_util

import (
	"github.com/mitubaEX/graphQL_sample/application/graphql_util/fields"
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user":      fields.UserField,
		"userList":  fields.UserListField,
		"event":     fields.EventField,
		"eventList": fields.EventListField,
	},
})
