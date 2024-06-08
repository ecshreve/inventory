package goinv

import (
	"os"
	"testing"
)

// setup function to initialize the test environment
func setup(t *testing.T) *GormInventory {
	os.Setenv("ENV", "test")
	inv, err := NewGormInventory()
	if err != nil {
		t.Fatalf("Failed to initialize inventory: %v", err)
	}
	// Optionally, clean the database or set initial state
	inv.db.Exec("DELETE FROM items")
	inv.db.Exec("DELETE FROM storage_locations")

	return inv
}

func TestInventory(t *testing.T) {
	t.Run("TestNewGormInventory", func(t *testing.T) {
		inv := setup(t)
		if inv.db == nil {
			t.Fatalf("Expected db to be initialized, got nil")
		}
	})

	t.Run("TestCreateItem", func(t *testing.T) {
		inv := setup(t)
		item := Item{Name: "TestItem", Category: "TestCategory"}
		err := inv.CreateItem(item)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("TestGetItems", func(t *testing.T) {
		inv := setup(t)
		item := Item{Name: "TestItem1", Category: "TestCategory1"}
		inv.CreateItem(item)
		items, err := inv.GetItems()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(items) == 0 {
			t.Fatalf("Expected to get items, got empty list")
		}
	})

	t.Run("TestUpdateItem", func(t *testing.T) {
		inv := setup(t)
		item := Item{Name: "TestItem2", Category: "TestCategory2"}
		inv.CreateItem(item)
		newItem := Item{Name: "UpdatedItem", Category: "UpdatedCategory"}
		err := inv.UpdateItem(1, newItem)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("TestDeleteItem", func(t *testing.T) {
		inv := setup(t)
		item := Item{Name: "TestItem3", Category: "TestCategory3"}
		inv.CreateItem(item)
		err := inv.DeleteItem(1)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("TestGetItemsByCategory", func(t *testing.T) {
		inv := setup(t)
		item := Item{Name: "TestItem4", Category: "SpecificCategory"}
		inv.CreateItem(item)
		items, err := inv.GetItemsByCategory("SpecificCategory")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(items) == 0 {
			t.Fatalf("Expected to get items, got empty list")
		}
	})

	t.Run("TestCreateStorageLocation", func(t *testing.T) {
		inv := setup(t)
		location := StorageLocation{Description: "TestLocationDescription1", Location: "TestLocation1"}
		err := inv.CreateStorageLocation(location)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("TestGetStorageLocations", func(t *testing.T) {
		inv := setup(t)
		location := StorageLocation{Description: "TestLocationDescription1", Location: "TestLocation1"}
		inv.CreateStorageLocation(location)
		locations, err := inv.GetStorageLocations()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(locations) == 0 {
			t.Fatalf("Expected to get locations, got empty list")
		}
	})

	t.Run("TestGetItemsByStorageLocation", func(t *testing.T) {
		inv := setup(t)
		loc := StorageLocation{Description: "TestLocationDescription1", Location: "TestLocation1"}
		inv.CreateStorageLocation(loc)
		item := Item{Name: "TestItem5", Category: "TestCategory5", Qty: 3, LocationID: loc.ID}
		inv.CreateItem(item)
		items, err := inv.GetItemsByStorageLocation(loc.ID)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(items) == 0 {
			t.Fatalf("Expected to get items, got empty list")
		}
	})

	t.Run("TestPopulate", func(t *testing.T) {
		inv := setup(t)

		locs := []StorageLocation{
			{ID: 1, Description: "TestLocationDescription1", Location: "TestLocation1"},
		}

		items := []Item{
			{Name: "TestItem1", Category: "TestCategory1", Qty: 3, LocationID: locs[0].ID},
		}

		err := inv.Populate(items, locs)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		allItems, _ := inv.GetItems()
		allLocs, _ := inv.GetStorageLocations()

		if len(allLocs) != len(locs) {
			t.Fatalf("Expected %d locations, got %d", len(locs), len(allLocs))
		}

		if len(allItems) != len(items) {
			t.Fatalf("Expected %d items, got %d", len(items), len(allItems))
		}
	})
}
