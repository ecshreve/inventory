package goinv

import "fmt"

type Item struct {
	ID       int             `json:"id"`
	Name     string          `json:"name"`
	Qty      int             `json:"qty"`
	Category ItemCategory    `json:"category"`
	Location StorageLocation `json:"location"`
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
