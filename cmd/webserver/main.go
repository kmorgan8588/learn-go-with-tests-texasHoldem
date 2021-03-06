package main

import (
	"go-app/server"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := server.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := server.NewPlayerServer(store)

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
