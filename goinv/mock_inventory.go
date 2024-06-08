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

func (m *MockInventory) GetItemsByStorageLocation(locationID uint) ([]Item, error) {
	for _, location := range m.StorageLocations {
		if location.ID == locationID {
			var items []Item
			for _, item := range m.Items {
				if item.LocationID == locationID {
					items = append(items, item)
				}
			}
			return items, nil
		}
	}

	return nil, fmt.Errorf("location not found")
}

func (m *MockInventory) Populate(items []Item, locs []StorageLocation) error {
	for _, location := range locs {
		if err := m.CreateStorageLocation(location); err != nil {
			return err
		}
	}

	for _, item := range items {
		if err := m.CreateItem(item); err != nil {
			return err
		}
	}

	return nil
}
