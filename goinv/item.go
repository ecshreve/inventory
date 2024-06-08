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

type ItemFilter struct {
	Category string
	Location string
}

type ItemCategory string

const (
	Cable   ItemCategory = "cable"
	Adapter ItemCategory = "adapter"
	Device  ItemCategory = "device"
	Misc    ItemCategory = "misc"
	Unknown ItemCategory = "unknown"
)

type StorageLocation string

const (
	HalfCrate_White_1 StorageLocation = "half_crate_white_1"
	HalfCrate_White_2 StorageLocation = "half_crate_white_2"
	FullCrate_Black_1 StorageLocation = "full_crate_black_1"
	FullCrate_Black_2 StorageLocation = "full_crate_black_2"
	FullCrate_Gray_1  StorageLocation = "full_crate_gray_1"
	HalfCrate_Gray_1  StorageLocation = "half_crate_gray_1"
	HalfCrate_Gray_2  StorageLocation = "half_crate_gray_2"
)
