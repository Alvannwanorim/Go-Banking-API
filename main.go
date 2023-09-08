package main

import (
	"log"

	_ "github.com/lib/pq"
)

func main() {
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
