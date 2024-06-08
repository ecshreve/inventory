// goinv/gorm_inventory.go

package goinv

import (
	"os"

	"github.com/charmbracelet/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GormInventory struct {
	db *gorm.DB
}

func NewGormInventory() (*GormInventory, error) {
	db_file := "inventory.db"
	if os.Getenv("ENV") == "test" {
		db_file = "test.db"
		if err := os.Remove(db_file); err != nil {
			log.Errorf("Failed to remove test database: %v", err)
		}
	}

	db, err := gorm.Open(sqlite.Open(db_file), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&StorageLocation{}, &Item{})
	return &GormInventory{db: db}, nil
}

func (g *GormInventory) CreateItem(item Item) error {
	return g.db.Create(&item).Error
}

func (g *GormInventory) GetItems() ([]Item, error) {
	var items []Item
	if err := g.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (g *GormInventory) UpdateItem(id uint, newItem Item) error {
	return g.db.Model(&Item{}).Where("id = ?", id).Updates(newItem).Error
}

func (g *GormInventory) DeleteItem(id uint) error {
	return g.db.Delete(&Item{}, id).Error
}

func (g *GormInventory) GetItemsByCategory(category string) ([]Item, error) {
	var items []Item
	if err := g.db.Where("category = ?", category).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (g *GormInventory) CreateStorageLocation(location StorageLocation) error {
	return g.db.Create(&location).Error
}

func (g *GormInventory) GetStorageLocations() ([]StorageLocation, error) {
	var locations []StorageLocation
	if err := g.db.Find(&locations).Error; err != nil {
		return nil, err
	}
	return locations, nil
}

func (g *GormInventory) Populate(items []Item, locs []StorageLocation) error {
	for _, location := range locs {
		if err := g.CreateStorageLocation(location); err != nil {
			return err
		}
	}

	for _, item := range items {
		if err := g.CreateItem(item); err != nil {
			return err
		}
	}

	return nil
}
