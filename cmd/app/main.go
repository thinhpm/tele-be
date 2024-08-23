package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"be/internal/config"
	"be/internal/handler"
	"be/internal/repository"
	"be/internal/service"

	"github.com/gorilla/mux"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Log the HTTP request details
		log.Printf("Started %s %s", r.Method, r.RequestURI)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the completion details
		duration := time.Since(startTime)
		log.Printf("Completed %s %s in %v", r.Method, r.RequestURI, duration)
	})
}

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Faill to load config file: %v", err)
	}

	repo, err := repository.NewMongoDBRepository(cfg.MongoDBURI, cfg.Database)

	if err != nil {
		log.Fatalln("Fail to connect db: %v", err)
	}

	defer repo.Disconnect()

	mongoService := service.NewMongoDBService(repo)
	mongoHandler := handler.NewMongoDBHandler(mongoService)

	r := mux.NewRouter()

	r.Use(LoggingMiddleware)

	r.HandleFunc("/api/v1/games", mongoHandler.GetAllGames).Methods("GET")
	r.HandleFunc("/api/v1/games/{game_id}", mongoHandler.GetGameByGameId).Methods("GET")

	http.Handle("/", r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
