package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitubaEX/graphQL_sample/application/graphql_util"
	"github.com/graphql-go/graphql"
	"net/http"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.PostFormValue("query"), graphql_util.NewSchema())
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)
}
