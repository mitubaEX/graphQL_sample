package fields

import (
	"errors"
	"github.com/mitubaEX/graphQL_sample/application/graphql_util/types"
	"github.com/mitubaEX/graphQL_sample/domain/model/user"
	"github.com/mitubaEX/graphQL_sample/domain/service"
	"github.com/mitubaEX/graphQL_sample/infrastructure"
	"github.com/graphql-go/graphql"
)

// fetch single user
var UserField = &graphql.Field{
	Type:        types.UserType,
	Description: "Get single user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		userId, isOK := params.Args["id"].(string)
		if isOK {
			return service.FindUserById(userId)
		}

		return nil, errors.New("no userId")
	},
}

// fetch all user
var UserListField = &graphql.Field{
	Type:        graphql.NewList(types.UserType),
	Description: "List of users",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return infrastructure.NewUserRepository().UserList(), nil
	},
}

// create user
var CreateUserField = &graphql.Field{
	Type:        types.UserType,
	Description: "Create new user",
	Args: graphql.FieldConfigArgument{
		"userName": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"photoURL": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		userName, _ := params.Args["userName"].(string)
		description, _ := params.Args["description"].(string)
		photoURL, _ := params.Args["photoURL"].(string)
		email, _ := params.Args["email"].(string)

		newUser, err := user.NewUser(userName, description, photoURL, email)
		if err != nil {
			panic(err)
		}
		infrastructure.NewUserRepository().Store(newUser)
		return newUser, nil
	},
}
