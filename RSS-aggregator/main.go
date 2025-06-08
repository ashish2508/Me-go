package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable is not set")
	}
	
	router := chi.NewRouter()
	
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:          300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Starting server on port: %v", portString)
	
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)
}
