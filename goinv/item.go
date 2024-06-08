package goinv

import "fmt"

type Item struct {
	ID       uint            `json:"id" gorm:"primaryKey"`
	Name     string          `json:"name"`
	Qty      uint            `json:"qty"`
	Category ItemCategory    `json:"category"`
	Location StorageLocation `json:"location"`
}

// String implements the fmt.Stringer interface for Item
func (i Item) String() string {
	return fmt.Sprintf("%s (%v) x%d -- %v", i.Name, i.Category, i.Qty, i.Location)
}

type ItemCategory string

const (
	Cable   ItemCategory = "CABLE"
	Adapter ItemCategory = "ADAPTER"
	Device  ItemCategory = "DEVICE"
	Misc    ItemCategory = "MISC"
	Unknown ItemCategory = "UNKNOWN"
)

type StorageLocation string

const (
	HalfCrate_White_1 StorageLocation = "HalfCrate_White_1"
	HalfCrate_White_2 StorageLocation = "HalfCrate_White_2"
	FullCrate_Black_1 StorageLocation = "FullCrate_Black_1"
)
