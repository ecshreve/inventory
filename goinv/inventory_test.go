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

	mockLocation := goinv.StorageLocation("MockLocation")

	// Create Item
	item := goinv.Item{
		ID:       1,
		Name:     "Test Item",
		Category: goinv.Misc,
		Qty:      5,
		Location: "Test Location",
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
	expectedItem := "Test Item (MISC) x5 -- Test Location"
	if fmt.Sprintf("%v", items[0]) != expectedItem {
		t.Fatalf("Expected item to be '%s', got '%s'", expectedItem, items[0])
	}

	// Update Item
	newItem := goinv.Item{
		ID:       1,
		Name:     "Updated Item",
		Category: goinv.Device,
		Qty:      10,
		Location: mockLocation,
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

	items := []goinv.Item{
		{
			ID:       1,
			Name:     "Test Item 1",
			Category: goinv.Misc,
			Qty:      5,
			Location: "Test Location",
		},
		{
			ID:       2,
			Name:     "Test Item 2",
			Category: goinv.Device,
			Qty:      10,
			Location: "Test Location",
		},
	}

	// Create Items
	for _, item := range items {
		if err := mockInv.CreateItem(item); err != nil {
			t.Fatalf("Failed to create item: %v", err)
		}
	}

	deviceItems, err := mockInv.GetItemsByCategory(string(goinv.Device))
	if err != nil {
		t.Fatalf("Failed to get items by category: %v", err)
	}

	if len(deviceItems) != 1 {
		t.Fatalf("Expected 1 item, got %d", len(deviceItems))
	}
}

func TestIntegration_GetItemsByLocation(t *testing.T) {
	mockInv := setupMockInventory()

	items := []goinv.Item{
		{
			ID:       1,
			Name:     "Test Item 1",
			Category: goinv.Misc,
			Qty:      5,
			Location: "Test Location",
		},
		{
			ID:       2,
			Name:     "Test Item 2",
			Category: goinv.Device,
			Qty:      10,
			Location: "Test Location",
		},
	}

	// Create Items
	for _, item := range items {
		if err := mockInv.CreateItem(item); err != nil {
			t.Fatalf("Failed to create item: %v", err)
		}
	}

	locationItems, err := mockInv.GetItemsByLocation("Test Location")
	if err != nil {
		t.Fatalf("Failed to get items by location: %v", err)
	}

	if len(locationItems) != 2 {
		t.Fatalf("Expected 2 items, got %d", len(locationItems))
	}
}
