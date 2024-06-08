package goinv

// StorageLocation represents a location where items are stored
// in an inventory system.
type StorageLocation struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Items       []Item `json:"items" gorm:"foreignKey:LocationID"`
}

// String implements the fmt.Stringer interface for StorageLocation.
func (s StorageLocation) String() string {
	return s.Description + " (" + s.Location + ")"
}
