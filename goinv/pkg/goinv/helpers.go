package goinv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ParseInventoryCSV(filename string) ([]Item, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var items []Item

	// Skip the header
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading record: %w", err)
		}

		quantity, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, fmt.Errorf("invalid quantity: %w", err)
		}

		item := Item{
			Name:     record[1],
			Qty:      uint(quantity),
			Category: ParseItemCategory(record[0]),
			Location: StorageLocation(record[3]),
		}

		items = append(items, item)
	}

	return items, nil
}

func ParseItemCategory(category string) ItemCategory {
	switch strings.ToLower(category) {
	case "cable":
		return Cable
	case "adapter":
		return Adapter
	case "device":
		return Device
	case "battery":
		return Misc
	default:
		return Unknown
	}
}
