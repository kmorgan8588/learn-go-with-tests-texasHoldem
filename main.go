package main

import (
	"go-app/server"
	"log"
	"net/http"
)

func main() {
	store := server.NewInMemoryPlayerStore()
	server := server.NewPlayerServer(store)

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
