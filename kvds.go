// main executable.
package main

import (
	"os"

	"github.com/noranetworks/kvds/internal/core"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	s, ok := core.New(os.Args[1:])
	if !ok {
		os.Exit(1)
	}

	//srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Core: s}}))

	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)

	//log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	//log.Fatal(http.ListenAndServe(":"+port, nil))
	s.Wait()
}
