package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/manjurulhoque/go-bookstore/pkg/routes"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	envPath, err := filepath.Abs("./../../")
	if err != nil {
		panic(err)
	}
	err = godotenv.Load(envPath + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
