package goinv_test

import (
	"fmt"
	"goinv"
	"os"
	"testing"
)

// helper function for common setup
func setupMockInventory() *goinv.MockInventory {
	os.Setenv("ENV", "test")
	return goinv.NewMockInventory()
}

func TestIntegration_CreateUpdateDeleteItem(t *testing.T) {
	mockInv := setupMockInventory()

	// Create Item
	item := goinv.Item{
		ID:       1,
		Name:     "Test Item",
		Category: goinv.Misc,
		Qty:      5,
		Location: goinv.StorageLocation{
			Description: "Test Location",
			Location:    "Test Room",
		},
	}

	if err := mockInv.CreateItem(item); err != nil {
		t.Fatalf("Failed to create item: %v", err)
	}

	items, err := mockInv.GetItems()
	if err != nil {
		t.Fatalf("Failed to get items: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("Expected 1 item, got %d", len(items))
	}
	if fmt.Sprintf("%v", items[0]) != "Test Item (Misc) x5 -- Test Location (Test Room)" {
		t.Fatalf("Expected item to be 'Test Item (Misc) x5 -- Test Location (Test Room)', got '%s'", items[0])
	}

	// Update Item
	newItem := goinv.Item{
		ID:       1,
		Name:     "Updated Item",
		Category: goinv.Device,
		Qty:      10,
		Location: goinv.StorageLocation{
			Description: "Updated Location",
			Location:    "Updated Room",
		},
	}

	// Attempt to update item with non-existent ID
	if err := mockInv.UpdateItem(2, newItem); err == nil {
		t.Fatalf("Expected error, got nil")
	}

	if err := mockInv.UpdateItem(1, newItem); err != nil {
		t.Fatalf("Failed to update item: %v", err)
	}

	updatedItems, err := mockInv.GetItems()
	if err != nil {
		t.Fatalf("Failed to get updated items: %v", err)
	}

	if len(updatedItems) != 1 {
		t.Fatalf("Expected 1 item, got %d", len(updatedItems))
	}

	if updatedItems[0].Name != "Updated Item" {
		t.Fatalf("Expected item name to be 'Updated Item', got '%s'", updatedItems[0].Name)
	}

	// Attempt to delete item with non-existent ID
	if err := mockInv.DeleteItem(2); err == nil {
		t.Fatalf("Expected error, got nil")
	}

	// Delete Item
	if err := mockInv.DeleteItem(1); err != nil {
		t.Fatalf("Failed to delete item: %v", err)
	}

	finalItems, err := mockInv.GetItems()
	if err != nil {
		t.Fatalf("Failed to get final items: %v", err)
	}

	if len(finalItems) != 0 {
		t.Fatalf("Expected 0 items, got %d", len(finalItems))
	}
}

func TestIntegration_GetItemsByCategory(t *testing.T) {
	mockInv := setupMockInventory()

	item1 := goinv.Item{
		ID:       1,
		Name:     "Test Item 1",
		Category: goinv.Misc,
		Qty:      5,
		Location: goinv.StorageLocation{
			Description: "Test Location",
			Location:    "Test Room",
		},
	}

	item2 := goinv.Item{
		ID:       2,
		Name:     "Test Item 2",
		Category: goinv.Device,
		Qty:      10,
		Location: goinv.StorageLocation{
			Description: "Test Location",
			Location:    "Test Room",
		},
	}

	if err := mockInv.CreateItem(item1); err != nil {
		t.Fatalf("Failed to create item1: %v", err)
	}

	if err := mockInv.CreateItem(item2); err != nil {
		t.Fatalf("Failed to create item2: %v", err)
	}

	items, err := mockInv.GetItemsByCategory(string(goinv.Device))
	if err != nil {
		t.Fatalf("Failed to get items by category: %v", err)
	}

	if len(items) != 1 {
		t.Fatalf("Expected 1 item, got %d", len(items))
	}
}

func TestIntegration_GetItemsByStorageLocation(t *testing.T) {
	mockInv := setupMockInventory()

	loc := goinv.StorageLocation{
		ID:          1,
		Description: "Test Location",
		Location:    "Test Room",
	}

	item1 := goinv.Item{
		ID:         1,
		Name:       "Test Item 1",
		Category:   goinv.Misc,
		Qty:        5,
		LocationID: loc.ID,
	}

	item2 := goinv.Item{
		ID:         2,
		Name:       "Test Item 2",
		Category:   goinv.Device,
		Qty:        10,
		LocationID: loc.ID,
	}

	if err := mockInv.CreateStorageLocation(loc); err != nil {
		t.Fatalf("Failed to create storage location: %v", err)
	}

	if err := mockInv.CreateItem(item1); err != nil {
		t.Fatalf("Failed to create item1: %v", err)
	}

	if err := mockInv.CreateItem(item2); err != nil {
		t.Fatalf("Failed to create item2: %v", err)
	}

	items, err := mockInv.GetItemsByStorageLocation(loc.ID)
	if err != nil {
		t.Fatalf("Failed to get items by storage location: %v", err)
	}

	if len(items) != 2 {
		t.Fatalf("Expected 2 items, got %d", len(items))
	}
}

func TestIntegration_Populate(t *testing.T) {
	mockInv := setupMockInventory()

	locs := []goinv.StorageLocation{
		{ID: 1, Description: "MockDescription", Location: "MockLocation"},
	}

	items := []goinv.Item{
		{
			ID:         1,
			Name:       "MockItem",
			Category:   goinv.Misc,
			Qty:        5,
			LocationID: locs[0].ID,
			Location:   locs[0],
		},
	}

	if err := mockInv.Populate(items, locs); err != nil {
		t.Fatalf("Failed to populate storage locations and items: %v", err)
	}

	locations, err := mockInv.GetStorageLocations()
	if err != nil {
		t.Fatalf("Failed to get storage locations: %v", err)
	}

	if len(locations) != 1 {
		t.Fatalf("Expected 1 location, got %d", len(locations))
	}

	populatedItems, err := mockInv.GetItems()
	if err != nil {
		t.Fatalf("Failed to get items: %v", err)
	}

	if len(populatedItems) != 1 {
		t.Fatalf("Expected 1 item, got %d", len(populatedItems))
	}

	if populatedItems[0].LocationID != locs[0].ID {
		t.Fatalf("Expected item location ID to be %d, got %d", locs[0].ID, populatedItems[0].LocationID)
	}
}
