package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gitlab.com/kw3a/go-rss/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

var dbg = flag.Bool("debug", false, "Enable debug mode")

func main() {
	flag.Parse()
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT isn't found in the environment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL isn't found in the environment")
	}
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database")
	}
	queries := database.New(conn)
	db := queries
	apiCfg := apiConfig{
		DB: db,
	}
	go startScrapping(db, 10, time.Minute)
	r := chi.NewRouter()

	v1 := chi.NewRouter()
	v1.Post("/users", apiCfg.handlerCreateUser)
	v1.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1.Get("/readiness", handlerReadiness)
	v1.Get("/error", handlerError)
	v1.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1.Get("/feeds", apiCfg.handlerGetFeeds)
	v1.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetUserFeedFollows))
	v1.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))
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
