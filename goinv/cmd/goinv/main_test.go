// main_test.go

package main

import (
	"goinv"
	"testing"
)

func TestCreateItem(t *testing.T) {
	mockInv := goinv.NewMockInventory()
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
		t.Errorf("Failed to create item: %v", err)
	}

	items, err := mockInv.GetItems()
	if err != nil {
		t.Errorf("Failed to get items: %v", err)
	}

	if len(items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items))
	}
}

func TestUpdateItem(t *testing.T) {
	mockInv := goinv.NewMockInventory()
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
		t.Errorf("Failed to create item: %v", err)
	}

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

	if err := mockInv.UpdateItem(1, newItem); err != nil {
		t.Errorf("Failed to update item: %v", err)
	}

	items, err := mockInv.GetItems()
	if err != nil {
		t.Errorf("Failed to get items: %v", err)
	}

	if len(items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items))
	}

	if items[0].Name != "Updated Item" {
		t.Errorf("Expected item name to be 'Updated Item', got '%s'", items[0].Name)
	}
}

func TestDeleteItem(t *testing.T) {
	mockInv := goinv.NewMockInventory()
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
		t.Errorf("Failed to create item: %v", err)
	}

	if err := mockInv.DeleteItem(1); err != nil {
		t.Errorf("Failed to delete item: %v", err)
	}

	items, err := mockInv.GetItems()
	if err != nil {
		t.Errorf("Failed to get items: %v", err)
	}

	if len(items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(items))
	}
}

func TestGetItemsByCategory(t *testing.T) {
	mockInv := goinv.NewMockInventory()
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
		t.Errorf("Failed to create item: %v", err)
	}

	if err := mockInv.CreateItem(item2); err != nil {
		t.Errorf("Failed to create item: %v", err)
	}

	items, err := mockInv.GetItemsByCategory(string(goinv.Device))
	if err != nil {
		t.Errorf("Failed to get items by category: %v", err)
	}

	if len(items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items))
	}
}

func TestPopulate(t *testing.T) {
	mockInv := goinv.NewMockInventory()
	if err := mockInv.Populate(); err != nil {
		t.Errorf("Failed to populate inventory: %v", err)
	}

	items, err := mockInv.GetItems()
	if err != nil {
		t.Errorf("Failed to get items: %v", err)
	}

	if len(items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(items))
	}

	locations, err := mockInv.GetStorageLocations()
	if err != nil {
		t.Errorf("Failed to get storage locations: %v", err)
	}

	if len(locations) != 3 {
		t.Errorf("Expected 3 locations, got %d", len(locations))
	}
}
