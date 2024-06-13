package goinv_test

import (
	"bytes"
	"encoding/json"
	"goinv"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	inventory := goinv.NewMockInventory()
	srv := goinv.NewServer(inventory)

	assert.NotNil(t, srv)
}

func TestGetItems(t *testing.T) {
	inventory := goinv.NewMockInventory()
	srv := goinv.NewServer(inventory)

	// Create an Item
	item := goinv.Item{
		Name:     "Test Item",
		Qty:      10,
		Category: goinv.Cable,
		Location: goinv.StorageLocation("A1"),
	}

	// Add the item to the inventory
	err := inventory.CreateItem(item)
	assert.NoError(t, err)

	// Create a new HTTP request to the "/items" endpoint
	req, err := http.NewRequest(http.MethodGet, "/items", nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Serve the request using the router
	srv.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var items []goinv.Item
	err = json.Unmarshal(w.Body.Bytes(), &items)
	assert.NoError(t, err)
	assert.NotEmpty(t, items)
	assert.Len(t, items, len(inventory.Items))
	assert.Equal(t, inventory.Items, items)
}

func TestCreateItem(t *testing.T) {
	inventory := goinv.NewMockInventory()
	srv := goinv.NewServer(inventory)

	// Create an Item
	item := goinv.Item{
		ID:       1,
		Name:     "Test Item",
		Qty:      10,
		Category: goinv.Cable,
		Location: goinv.StorageLocation("A1"),
	}

	// Marshal the item to JSON
	data, err := json.Marshal(item)
	assert.NoError(t, err)

	// Create a new HTTP request to the "/item" endpoint
	req, err := http.NewRequest(http.MethodPost, "/item", bytes.NewReader(data))
	assert.NoError(t, err)

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Serve the request using the router
	srv.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Check the response body
	var newItem goinv.Item
	err = json.Unmarshal(w.Body.Bytes(), &newItem)
	assert.NoError(t, err)
	assert.Equal(t, item, newItem)
}

func TestUpdateItem(t *testing.T) {
	inventory := goinv.NewMockInventory()
	srv := goinv.NewServer(inventory)

	// Create an Item
	item := goinv.Item{
		ID:       1,
		Name:     "Test Item",
		Qty:      10,
		Category: goinv.Cable,
		Location: goinv.StorageLocation("A1"),
	}

	// Add the item to the inventory
	err := inventory.CreateItem(item)
	assert.NoError(t, err)

	// Update the item
	item.Qty = 20

	// Marshal the item to JSON
	data, err := json.Marshal(item)
	assert.NoError(t, err)

	// Create a new HTTP request to the "/item/:id" endpoint
	req, err := http.NewRequest(http.MethodPut, "/item/1", bytes.NewReader(data))
	assert.NoError(t, err)

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Serve the request using the router
	srv.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var updatedItem goinv.Item
	err = json.Unmarshal(w.Body.Bytes(), &updatedItem)
	assert.NoError(t, err)
	assert.Equal(t, item, updatedItem)
}

func TestDeleteItem(t *testing.T) {
	inventory := goinv.NewMockInventory()
	srv := goinv.NewServer(inventory)

	// Create an Item
	item := goinv.Item{
		ID:       1,
		Name:     "Test Item",
		Qty:      10,
		Category: goinv.Cable,
		Location: goinv.StorageLocation("A1"),
	}

	// Add the item to the inventory
	err := inventory.CreateItem(item)
	assert.NoError(t, err)

	// Create a new HTTP request to the "/item/:id" endpoint
	req, err := http.NewRequest(http.MethodDelete, "/item/1", nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Serve the request using the router
	srv.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusNoContent, w.Code)

	// Check the response body
	assert.Empty(t, w.Body.String())
}
