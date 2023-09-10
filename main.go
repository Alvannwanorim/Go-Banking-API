package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := APIServer{listenAddr: ":3000", store: store}
	server.Run()
}
