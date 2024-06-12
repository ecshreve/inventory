package goinv

import "fmt"

// InMemoryInventory is an in-memory implementation of the Inventory interface
// that stores items as a map of id to Item.
type InMemoryInventory struct {
	nextID uint
	items  map[uint]Item
}

// NewInMemoryInventory creates a new InMemoryInventory.
func NewInMemoryInventory() *InMemoryInventory {
	return &InMemoryInventory{
		nextID: 1,
		items:  make(map[uint]Item),
	}
}

func (i *InMemoryInventory) ItemLocationMap() map[StorageLocation][]Item {
	locationMap := make(map[StorageLocation][]Item)
	for _, item := range i.items {
		locationMap[item.Location] = append(locationMap[item.Location], item)
	}
	return locationMap
}

func (i *InMemoryInventory) ItemCategoryMap() map[ItemCategory][]Item {
	categoryMap := make(map[ItemCategory][]Item)
	for _, item := range i.items {
		categoryMap[item.Category] = append(categoryMap[item.Category], item)
	}
	return categoryMap
}

// CreateItem adds an item to the inventory.
func (i *InMemoryInventory) CreateItem(item Item) error {
	item.ID = i.nextID
	i.items[item.ID] = item
	i.nextID++

	return nil
}

// GetItems returns all items in the inventory.
func (i *InMemoryInventory) GetItems() ([]Item, error) {
	items := make([]Item, 0, len(i.items))
	for _, item := range i.items {
		items = append(items, item)
	}
	return items, nil
}

// UpdateItem updates an item in the inventory.
func (i *InMemoryInventory) UpdateItem(id uint, newItem Item) error {
	if _, ok := i.items[id]; !ok {
		return fmt.Errorf("item with ID %d not found", id)
	}
	newItem.ID = id
	i.items[id] = newItem
	return nil
}

// DeleteItem removes an item from the inventory.
func (i *InMemoryInventory) DeleteItem(id uint) error {
	if _, ok := i.items[id]; !ok {
		return fmt.Errorf("item with ID %d not found", id)
	}
	delete(i.items, id)
	return nil
}

// GetItemsByCategory returns all items in the inventory with the given category.
func (i *InMemoryInventory) GetItemsByCategory(category string) ([]Item, error) {
	items := make([]Item, 0)
	for _, item := range i.items {
		if string(item.Category) == category {
			items = append(items, item)
		}
	}
	return items, nil
}

// GetItemsByLocation returns all items in the inventory with the given location.
func (i *InMemoryInventory) GetItemsByLocation(location string) ([]Item, error) {
	items := make([]Item, 0)
	for _, item := range i.items {
		if string(item.Location) == location {
			items = append(items, item)
		}
	}
	return items, nil
}

// GetItemsWithFilter returns all items in the inventory that match the given filter.
func (i *InMemoryInventory) GetItemsWithFilter(filter ItemFilter) ([]Item, error) {
	items := make([]Item, 0)
	for _, item := range i.items {
		if (filter.Category == "" || string(item.Category) == filter.Category) &&
			(filter.Location == "" || string(item.Location) == filter.Location) {
			items = append(items, item)
		}
	}
	return items, nil
}
