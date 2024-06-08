package goinv

// StorageLocation represents a location where items are stored
// in an inventory system.
type StorageLocation struct {
	ID          uint
	Description string
	Location    string
	Items       []Item
}

// String implements the fmt.Stringer interface for StorageLocation.
func (s StorageLocation) String() string {
	return s.Description + " (" + s.Location + ")"
}
