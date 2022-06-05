package main

import (
	"log"
	"os"

	"github.com/rwxd/zaster/internal"
	"github.com/rwxd/zaster/tui"
)

func openDatabase() *internal.JSONDatabase {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	filePath := homedir + "/zaster-db.json"
	db := internal.NewJSONDatabase(filePath)
	return &db
}

func main() {
	db := openDatabase()
	tui.StartTea(db)
}
