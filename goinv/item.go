package goinv

import "fmt"

type Item struct {
	ID         uint            `json:"id" gorm:"primaryKey"`
	Name       string          `json:"name"`
	Qty        uint            `json:"qty"`
	Category   ItemCategory    `json:"category"`
	LocationID uint            `json:"location_id"`
	Location   StorageLocation `json:"-" gorm:"foreignKey:LocationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// String implements the fmt.Stringer interface for Item
func (i Item) String() string {
	return fmt.Sprintf("%s (%v) x%d -- %v", i.Name, i.Category, i.Qty, i.Location)
}

type ItemCategory string

const (
	Cable   ItemCategory = "Cable"
	Adapter ItemCategory = "Adapter"
	Device  ItemCategory = "Device"
	Misc    ItemCategory = "Misc"
	Unknown ItemCategory = "Unknown"
)
