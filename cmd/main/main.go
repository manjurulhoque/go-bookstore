package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/manjurulhoque/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
