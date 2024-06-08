package goinv

type Inventory interface {
	CreateItem(item Item) error
	GetItems() ([]Item, error)
	UpdateItem(id uint, newItem Item) error
	DeleteItem(id uint) error
	GetItemsByCategory(category string) ([]Item, error)
	GetItemsByLocation(location string) ([]Item, error)
	GetItemsWithFilter(filter ItemFilter) ([]Item, error)
}
