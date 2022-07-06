package main

import (
	"fmt"
	"log"
	"net/http"
	v1 "userservice/userservice/api/v1"
	"userservice/userservice/db"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Some error occured. Err: %s", err)
		return
	}
	fmt.Println("Running....")
	StartServer()
}

func StartServer() {
	router := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000", // will map live IP
		},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{
			"Origin", "Authorization", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Header", "Accept",
			"Content-Type", "X-CSRF-Token",
		},
		ExposedHeaders: []string{
			"Content-Length", "Access-Control-Allow-Origin", "Origin",
		},
		AllowCredentials: true,
		MaxAge:           300,
	})
	// cross & loger middleware
	router.Use(cors.Handler)
	router.Use(
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
	)
	//database server connection'
	db.DatabaseServerStart()
	// Initialize the version 3 routes of the public API
	router.Route("/", v1.Routes)
	log.Fatal(http.ListenAndServe(":"+"9000", router))

}
