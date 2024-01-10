package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	fileserverHits int
	jwtSecret      string
	polkaKey       string
}

var dbg = flag.Bool("debug", false, "Enable debug mode")

func main() {
	flag.Parse()
	godotenv.Load()

	port := os.Getenv("PORT")

	if *dbg {
		err := os.Remove("database.json")
		if err != nil {
			fmt.Println("Database delete error:", err)
		}
	}

	r := chi.NewRouter()

	v1 := chi.NewRouter()
	v1.Get("/readiness", handlerReadiness)
	v1.Get("/error", handlerError)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Mount("/v1", v1)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
