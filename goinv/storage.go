package goinv

type StorageLocation struct {
	ID          int
	Description string
	Location    string
	Items       []Item
}

// String implements the fmt.Stringer interface for StorageLocation
func (s StorageLocation) String() string {
	return s.Description + " (" + s.Location + ")"
}
