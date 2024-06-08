// goinv/mock_inventory.go

package goinv

import "fmt"

type MockInventory struct {
	Items []Item
}

func NewMockInventory() *MockInventory {
	return &MockInventory{
		Items: []Item{},
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
