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
	log.Info("NewGormInventory")
	db_file := "inventory.db"
	if os.Getenv("ENV") == "test" {
		db_file = "test.db"
	}

	db, err := gorm.Open(sqlite.Open(db_file), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Item{})
	return &GormInventory{db: db}, nil
}

func (g *GormInventory) CreateItem(item Item) error {
	log.Info("CreateItem")
	return g.db.Create(&item).Error
}

func (g *GormInventory) GetItems() ([]Item, error) {
	log.Info("GetItems")
	var items []Item
	if err := g.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (g *GormInventory) UpdateItem(id uint, newItem Item) error {
	log.Info("UpdateItem")
	return g.db.Model(&Item{}).Where("id = ?", id).Updates(newItem).Error
}

func (g *GormInventory) DeleteItem(id uint) error {
	log.Info("DeleteItem")
	return g.db.Delete(&Item{}, id).Error
}

func (g *GormInventory) GetItemsByCategory(category string) ([]Item, error) {
	log.Info("GetItemsByCategory")
	var items []Item
	if err := g.db.Where("category = ?", category).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (g *GormInventory) GetItemsByLocation(location string) ([]Item, error) {
	log.Info("GetItemsByLocation")
	var items []Item
	if err := g.db.Where("location = ?", location).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (g *GormInventory) GetItemsWithFilter(filter ItemFilter) ([]Item, error) {
	log.Info("GetItemsWithFilter")
	var items []Item
	query := g.db
	if filter.Category != "" {
		query = query.Where("category = ?", filter.Category)
	}
	if filter.Location != "" {
		query = query.Where("location = ?", filter.Location)
	}
	if err := query.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
