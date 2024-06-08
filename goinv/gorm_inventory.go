// goinv/gorm_inventory.go

package goinv

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GormInventory struct {
	db *gorm.DB
}

func NewGormInventory() (*GormInventory, error) {
	db, err := gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Item{})
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
