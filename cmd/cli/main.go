package main

import (
	"fmt"
	"go-app/server"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := server.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	server.NewCLI(store, os.Stdin).PlayPoker()
}
