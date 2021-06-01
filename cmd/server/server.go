package main

import (
	"log"
    "net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/rs/cors"

	auHandler "github.com/raganmartinez-hf/gqlgen-katana/handler/au"
	usHandler "github.com/raganmartinez-hf/gqlgen-katana/handler/us"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	// AU endpoints //
	auHandler.InitAUHandler(router)

	// US endpoints
	usHandler.InitUSHandler(router)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
