package main

import (
	"os"

	"goinv"

	"github.com/charmbracelet/log"
)

func main() {
	log.Info("Starting goinv")
	os.Setenv("ENV", "prod")

	inventory, err := goinv.NewGormInventory()
	if err != nil {
		log.Fatal("Failed to initialize inventory:", err)
	}

	srv := goinv.NewServer(inventory)

	log.Info("Listening on :8080")
	srv.Run(":8080")
}
