// goinv/mock_inventory.go

package goinv

import "fmt"

type MockInventory struct {
	Items            []Item
	StorageLocations []StorageLocation
}

func NewMockInventory() *MockInventory {
	return &MockInventory{
		Items:            []Item{},
		StorageLocations: []StorageLocation{},
	}
}

func (m *MockInventory) CreateItem(item Item) error {
	m.Items = append(m.Items, item)
	return nil
}

func (m *MockInventory) GetItems() ([]Item, error) {
	return m.Items, nil
}

func (m *MockInventory) UpdateItem(id uint, newItem Item) error {
	for i, item := range m.Items {
		if item.ID == id {
			m.Items[i] = newItem
			return nil
		}
	}
	return fmt.Errorf("item not found")
}

func (m *MockInventory) DeleteItem(id uint) error {
	for i, item := range m.Items {
		if item.ID == id {
			m.Items = append(m.Items[:i], m.Items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item not found")
}

func (m *MockInventory) GetItemsByCategory(category string) ([]Item, error) {
	var items []Item
	for _, item := range m.Items {
		if string(item.Category) == category {
			items = append(items, item)
		}
	}
	return items, nil
}

func (m *MockInventory) CreateStorageLocation(location StorageLocation) error {
	m.StorageLocations = append(m.StorageLocations, location)
	return nil
}

func (m *MockInventory) GetStorageLocations() ([]StorageLocation, error) {
	return m.StorageLocations, nil
}

func (m *MockInventory) Populate() error {
	locations := []StorageLocation{
		{ID: 1, Description: "HalfCrate_White_1", Location: "Office"},
		{ID: 2, Description: "FullCrate_Black_1", Location: "Office"},
		{ID: 3, Description: "HalfCrate_White_2", Location: "Office"},
	}

	for _, location := range locations {
		if err := m.CreateStorageLocation(location); err != nil {
			return err
		}
	}

	return nil
}
