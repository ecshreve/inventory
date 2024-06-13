// goinv/mock_inventory.go

package goinv

import (
	"fmt"

	"github.com/charmbracelet/log"
)

type MockInventory struct {
	Items            []Item
	StorageLocations []StorageLocation
}

func NewMockInventory() *MockInventory {
	log.Info("NewMockInventory")
	return &MockInventory{
		Items:            []Item{},
		StorageLocations: []StorageLocation{},
	}
}

func (m *MockInventory) CreateItem(item Item) error {
	log.Info("CreateItem")
	m.Items = append(m.Items, item)
	return nil
}

func (m *MockInventory) GetItems() ([]Item, error) {
	log.Info("GetItems")
	return m.Items, nil
}

func (m *MockInventory) UpdateItem(id uint, newItem Item) error {
	log.Info("UpdateItem")
	for i, item := range m.Items {
		if item.ID == id {
			m.Items[i] = newItem
			return nil
		}
	}
	return fmt.Errorf("item not found")
}

func (m *MockInventory) DeleteItem(id uint) error {
	log.Info("DeleteItem")
	for i, item := range m.Items {
		if item.ID == id {
			m.Items = append(m.Items[:i], m.Items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item not found")
}

func (m *MockInventory) GetItemsByCategory(category string) ([]Item, error) {
	log.Info("GetItemsByCategory")
	var items []Item
	for _, item := range m.Items {
		if string(item.Category) == category {
			items = append(items, item)
		}
	}
	return items, nil
}

func (m *MockInventory) GetItemsByLocation(location string) ([]Item, error) {
	log.Info("GetItemsByLocation")
	var items []Item
	for _, item := range m.Items {
		if string(item.Location) == location {
			items = append(items, item)
		}
	}
	return items, nil
}

func (m *MockInventory) GetItemsWithFilter(filter ItemFilter) ([]Item, error) {
	log.Info("GetItemsWithFilter")
	var items []Item
	for _, item := range m.Items {
		if filter.Category != "" && string(item.Category) != filter.Category {
			continue
		}
		if filter.Location != "" && string(item.Location) != filter.Location {
			continue
		}
		items = append(items, item)
	}
	return items, nil
}
