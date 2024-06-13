package main

type ItemCategory string

const (
	Cable   ItemCategory = "cable"
	Adapter ItemCategory = "adapter"
	Device  ItemCategory = "device"
	Misc    ItemCategory = "misc"
	Unknown ItemCategory = "unknown"
)

func AllCategories() []string {
	return []string{
		string(Cable),
		string(Adapter),
		string(Device),
		string(Misc),
		string(Unknown),
	}
}

type StorageLocation string

const (
	HCW1 StorageLocation = "half_crate_white_1"
	HCW2 StorageLocation = "half_crate_white_2"
	FCB1 StorageLocation = "full_crate_black_1"
	FCG1 StorageLocation = "full_crate_gray_1"
)

func AllLocations() []string {
	return []string{
		string(HCW1),
		string(HCW2),
		string(FCB1),
		string(FCG1),
	}
}

const SeedItemsCSV = `Category,Item,Quantity,Location
adapter,AC to USB-A with button,1,half_crate_white_1
device,Asus mini laptop,1,half_crate_white_1
adapter,Govee 12V - 1.5A,2,half_crate_white_1
adapter,Single 5V 1A,7,half_crate_white_1
adapter,AC to USB-C,1,half_crate_white_1
adapter,19V - 2.1A,1,half_crate_white_1
adapter,5V - 2A,2,half_crate_white_1
adapter,12V - 1.5A,1,half_crate_white_1
adapter,Double 5V - 2.1A,1,half_crate_white_1
adapter,Double 5V - 2.4A,1,half_crate_white_1
device,Sony Cybershot Battery,1,half_crate_white_1
cable,USB-A to Micro-USB,14,half_crate_white_1
device,USB Hub (4/3 ports),1,full_crate_black_1
cable,USB-A to USB-B,8,full_crate_black_1
cable,USB-A to USB-C,8,full_crate_black_1
cable,USB-C to Lightning,2,full_crate_black_1
cable,USB-A to Lightning,4,full_crate_black_1
cable,USB-C to USB-C,2,full_crate_black_1
device,USB N64,1,half_crate_white_2
device,USB PS2,1,half_crate_white_2
device,USB NES,1,half_crate_white_2
device,USB SNES,1,half_crate_white_2
device,USB Sega,1,half_crate_white_2
device,iPhone,2,half_crate_white_2
device,Raspberry Pi 4,1,half_crate_white_2
device,Raspberry Pi Zero,1,half_crate_white_2
device,Intel Stick Computer,1,half_crate_white_2
device,iPod Color,1,half_crate_white_2
device,External Hard Drive,2,half_crate_white_2
device,Power Bank,2,half_crate_white_2
device,Pendo Box,1,half_crate_white_2
device,Mini Monitor,1,half_crate_white_2
cable,VGA to VGA,1,half_crate_white_2
`
