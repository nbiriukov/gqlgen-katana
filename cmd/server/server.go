package main

import (
	"log"
	"net/http"
	"os"

	auHandler "github.com/raganmartinez-hf/gqlgen-katana/handler/au"
	usHandler "github.com/raganmartinez-hf/gqlgen-katana/handler/us"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// AU endpoints //
	auHandler.InitAUHandler()

	// US endpoints
	usHandler.InitUSHandler()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
