package goinv

import "fmt"

type Inventory struct {
	Items []Item
}

// String implements the fmt.Stringer interface for Inventory
func (i Inventory) String() string {
	return fmt.Sprintf("%+v", i.Items)
}
