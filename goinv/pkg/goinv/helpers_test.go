package goinv_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"goinv"
)

func TestParseInventoryCSV(t *testing.T) {
	items, err := goinv.ParseInventoryCSV("testdata/test_inventory.csv")
	assert.NoError(t, err)
	assert.Len(t, items, 2)

	// Category,Item,Quantity,Location
	// cable,USB-A to Micro-USB,14,half_crate_white_1
	// device,Raspberry Pi Zero,1,half_crate_white_2

	assert.Equal(t, goinv.Item{
		Name:     "USB-A to Micro-USB",
		Qty:      14,
		Category: goinv.Cable,
		Location: goinv.StorageLocation("half_crate_white_1"),
	}, items[0])

	assert.Equal(t, goinv.Item{
		Name:     "Raspberry Pi Zero",
		Qty:      1,
		Category: goinv.Device,
		Location: goinv.StorageLocation("half_crate_white_2"),
	}, items[1])
}
