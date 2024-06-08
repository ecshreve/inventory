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

	// Load the inventory from a CSV file
	items, err := goinv.ParseInventoryCSV("data/inventory.csv")
	if err != nil {
		log.Fatal("Failed to parse inventory CSV:", err)
	}

	for _, item := range items {
		if err := inventory.CreateItem(item); err != nil {
			log.Fatal("Failed to create item:", err)
		}
	}

	srv := goinv.NewServer(inventory)

	log.Info("Listening on :8080")
	srv.Run(":8080")
}
