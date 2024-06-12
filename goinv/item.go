package goinv

import "fmt"

type Item struct {
	ID       uint            `json:"id" gorm:"primaryKey"`
	Name     string          `json:"name"`
	Qty      uint            `json:"qty"`
	Category ItemCategory    `json:"category"`
	Location StorageLocation `json:"location"`
}

// Implement the charmbracelet/list widget's Item interface
func (i Item) Title() string {
	return i.Name
}

// Description returns a string that describes the item
func (i Item) Description() string {
	return fmt.Sprintf("( %d ) %s -- %s", i.Qty, i.Category, i.Location)
}

// FilterValue returns the value of the item that should be used for filtering
func (i Item) FilterValue() string {
	return fmt.Sprintf("%s %s", i.Name, i.Category)
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

var AllCategories = []string{
	string(Cable),
	string(Adapter),
	string(Device),
	string(Misc),
	string(Unknown),
}

type StorageLocation string

func (l StorageLocation) String() string {
	return string(l)
}

const (
	HalfCrate_White_1   StorageLocation = "half_crate_white_1"
	HalfCrate_White_2   StorageLocation = "half_crate_white_2"
	FullCrate_Black_1   StorageLocation = "full_crate_black_1"
	FullCrate_Black_2   StorageLocation = "full_crate_black_2"
	FullCrate_Gray_1    StorageLocation = "full_crate_gray_1"
	FullCrate_Stealth_1 StorageLocation = "full_crate_stealth_1"
	FullCrate_Stealth_2 StorageLocation = "full_crate_stealth_2"
	HalfCrate_Stealth_1 StorageLocation = "half_crate_stealth_1"
	HalfCrate_Stealth_2 StorageLocation = "half_crate_stealth_2"
	HalfCrate_Orange_1  StorageLocation = "half_crate_orange_1"
)

var AllLocations = []string{
	string(HalfCrate_White_1),
	string(HalfCrate_White_2),
	string(FullCrate_Black_1),
	string(FullCrate_Black_2),
	string(FullCrate_Gray_1),
	string(FullCrate_Stealth_1),
	string(FullCrate_Stealth_2),
	string(HalfCrate_Stealth_1),
	string(HalfCrate_Stealth_2),
	string(HalfCrate_Orange_1),
}
