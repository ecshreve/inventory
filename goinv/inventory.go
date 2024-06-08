package goinv

type Inventory interface {
	CreateItem(item Item) error
	GetItems() ([]Item, error)
	UpdateItem(id uint, newItem Item) error
	DeleteItem(id uint) error
	GetItemsByCategory(category string) ([]Item, error)
	CreateStorageLocation(location StorageLocation) error
	GetStorageLocations() ([]StorageLocation, error)
	Populate([]Item, []StorageLocation) error
}
