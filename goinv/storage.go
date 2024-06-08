package goinv

// StorageLocation represents a location where items are stored
// in an inventory system.
type StorageLocation struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

// String implements the fmt.Stringer interface for StorageLocation.
func (s StorageLocation) String() string {
	return s.Description + " (" + s.Location + ")"
}
