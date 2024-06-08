package main

import (
	"fmt"
	"goinv"
)

func main() {
	inv := goinv.Inventory{
		Items: []goinv.Item{
			{
				ID:       1,
				Name:     "USB-C to USB-A",
				Qty:      10,
				Category: goinv.Cable,
				Location: goinv.StorageLocation{
					ID:          1,
					Description: "Top shelf",
					Location:    "Shelf 1",
				},
			},
		},
	}

	fmt.Println(inv)
}
