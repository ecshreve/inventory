package main

import (
	"context"
	"fmt"
	"goinv/ent"
	"goinv/ent/item"
	"goinv/ent/migrate"
	"goinv/ent/storagelocation"
	"log"
	"strconv"
	"strings"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

var SQLITE_DB = "file:file.db?mode=rwc&cache=shared&_fk=1"

func main() {
	// log.Fatal("SKIP the database")

	// Create ent.Client and run the schema migration.
	// client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open(dialect.SQLite, SQLITE_DB)
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Populate the storage locations.
	if err := populateLocations(client); err != nil {
		log.Fatal("populating ent client", err)
	}

	// Populate seedItems
	if err := populateSeedItems(client); err != nil {
		log.Fatal("populating ent client", err)
	}
}

func populateLocations(client *ent.Client) error {
	ctx := context.Background()

	// Create storage locations.
	allLocs := []*ent.StorageLocationCreate{}
	for _, loc := range AllLocations() {
		locMeta := strings.Split(loc, "_")
		s := storagelocation.Size(locMeta[0])
		c := storagelocation.Color(locMeta[2])

		allLocs = append(allLocs, client.StorageLocation.Create().SetName(loc).SetSize(s).SetColor(c))
	}
	client.StorageLocation.CreateBulk(allLocs...).SaveX(ctx)
	// log.Println("created storage locations:", createdLocs)

	return nil
}

func populateSeedItems(client *ent.Client) error {
	ctx := context.Background()

	// Parse csv data
	parsed := parseCSVData(SeedItemsCSV)
	fmt.Println(parsed)

	// Create items.
	allItems := []*ent.ItemCreate{}
	for _, itemStr := range parsed {
		if len(itemStr) != 4 {
			log.Default().Printf("invalid item: %v", itemStr)
			continue
		}
		cat := itemStr[1]
		name := itemStr[2]
		qty, err := strconv.Atoi(itemStr[3])
		if err != nil {
			log.Fatalf("invalid quantity: %s", itemStr[2])
		}

		loc := itemStr[0]

		// Get the storage location.
		locEnt := client.StorageLocation.Query().Where(storagelocation.NameEQ(loc)).OnlyX(ctx)
		if locEnt == nil {
			log.Fatalf("storage location not found: %s", loc)
		}

		allItems = append(allItems,
			client.Item.Create().
				SetCategory(item.Category(cat)).
				SetName(name).
				SetQuantity(qty).
				SetStorageLocation(locEnt),
		)
	}
	createdItems := client.Item.CreateBulk(allItems...).SaveX(ctx)
	log.Println("created items:", createdItems)

	return nil
}

func parseCSVData(data string) [][]string {
	// Split the data into lines.
	lines := strings.Split(data, "\n")
	// Split the lines into fields.
	fields := make([][]string, len(lines))
	for i, line := range lines {
		if i == 0 || len(line) == 0 {
			continue
		}
		fields[i] = strings.Split(line, ",")
	}
	return fields
}
