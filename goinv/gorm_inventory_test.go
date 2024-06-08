package goinv

import (
	"os"
	"testing"
)

// setup function to initialize the test environment
func setup(t *testing.T) *GormInventory {
	os.Setenv("ENV", "test")
	if err := os.Remove("test.db"); err != nil {
		t.Fatalf("Failed to remove test.db: %v", err)
	}

	inv, err := NewGormInventory()
	if err != nil {
		t.Fatalf("Failed to initialize inventory: %v", err)
	}
	// Optionally, clean the database or set initial state
	inv.db.Exec("DELETE FROM items")

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
		if err := inv.CreateItem(item); err != nil {
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
		if err := inv.UpdateItem(1, newItem); err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("TestDeleteItem", func(t *testing.T) {
		inv := setup(t)
		item := Item{Name: "TestItem3", Category: "TestCategory3"}
		inv.CreateItem(item)
		if err := inv.DeleteItem(1); err != nil {
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

	t.Run("TestGetItemsByLocation", func(t *testing.T) {
		inv := setup(t)
		item := Item{Name: "TestItem5", Category: "TestCategory5", Location: "SpecificLocation"}
		inv.CreateItem(item)
		items, err := inv.GetItemsByLocation("SpecificLocation")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(items) == 0 {
			t.Fatalf("Expected to get items, got empty list")
		}
	})

	t.Run("TestGetItemsWithFilter", func(t *testing.T) {
		inv := setup(t)

		filterCategory := "TestCategory6"
		filterLocation := "SpecificLocation"

		item := Item{Name: "TestItem6", Category: ItemCategory(filterCategory), Location: StorageLocation(filterLocation)}
		inv.CreateItem(item)
		filter := ItemFilter{Category: filterCategory, Location: filterLocation}
		items, err := inv.GetItemsWithFilter(filter)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(items) == 0 {
			t.Fatalf("Expected to get items, got empty list")
		}
	})
}
