package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQueryType",
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "world", nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
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
		result := executeQuery(r.URL.Query()["query"][0], schema)
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/rest", func(w http.ResponseWriter, r *http.Request) {
		respondWithJSON(w, 200, r.URL.Query()["query"][0])
	})

	fmt.Println("Benchmark app listening on port 8080")

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}

// Payload payload type
type Payload interface{}

var (
	mutex sync.RWMutex
)

// respondWithJSON respond with json
func respondWithJSON(w http.ResponseWriter, code int, payload Payload) {
	buf := new(bytes.Buffer)
	defer buf.Reset()
	defer mutex.Unlock()

	mutex.Lock()
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
	w.Write(buf.Bytes())
	w.Header().Set("Content-Type", "application/json")
}
