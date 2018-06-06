package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/mitubaEX/graphQL_sample/application/graphql_util"
	"io/ioutil"
	"net/http"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {

	// schema と もらってきたqueryを入れて，実行
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
	// /graphql のリクエストを受け取れるようにする
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// bodyの読み取り処理
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		// query 実行
		result := executeQuery(fmt.Sprintf("%s", body), graphql_util.Schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
