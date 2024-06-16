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
	HCS1 StorageLocation = "half_crate_stealth_1"
	HCO1 StorageLocation = "half_crate_orange_1"
	FCB1 StorageLocation = "full_crate_black_1"
	FCB2 StorageLocation = "full_crate_black_2"
	FCG1 StorageLocation = "full_crate_gray_1"
	FCG2 StorageLocation = "full_crate_gray_2"
	FCS1 StorageLocation = "full_crate_stealth_1"
	FCS2 StorageLocation = "full_crate_stealth_2"
)

func AllLocations() []string {
	return []string{
		string(HCW1),
		string(HCW2),
		string(HCS1),
		string(HCO1),
		string(FCB1),
		string(FCB2),
		string(FCG1),
		string(FCG2),
		string(FCS1),
		string(FCS2),
	}
}

const SeedItemsCSV = `location,category,item,quantity
half_crate_white_1,adapter,AC to USB-A with button,1
half_crate_white_1,adapter,Govee 12V - 1.5A,2
half_crate_white_1,adapter,Single 5V 1A,7
half_crate_white_1,adapter,AC to USB-C,1
half_crate_white_1,adapter,19V - 2.1A,1
half_crate_white_1,adapter,5V - 2A,2
half_crate_white_1,adapter,12V - 1.5A ,1
half_crate_white_1,adapter,5V - 2.1A - 2x USB-A,1
half_crate_white_1,adapter,5V - 2.4A - 2x USB-A,1
half_crate_white_1,cable,USB-A to Micro-USB,14
full_crate_black_1,cable,USB-A to USB-B,8
full_crate_black_1,cable,USB-A to USB-C,8
full_crate_black_1,cable,USB-C to Lightning,2
full_crate_black_1,cable,USB-A to Lightning,4
full_crate_black_1,cable,USB-C to USB-C,2
half_crate_white_2,cable,VGA to VGA,1
half_crate_stealth_1,cable,AC - D (Nema 15),6
half_crate_stealth_1,cable,CAT6 - 100ft,1
full_crate_black_2,cable,CAT6 - 6ft,20
full_crate_black_2,cable,CAT6 - 3ft,13
full_crate_black_2,cable,CAT6 - 10ft,2
full_crate_black_1,cable,Miscellaneos Power Cords,6
full_crate_stealth_1,cable,Intel Nuc Power,1
full_crate_stealth_1,cable,Lenovo Mini Power,1
full_crate_stealth_1,cable,Half of Lenovo mini,1
half_crate_white_1,device,Asus mini laptop,1
half_crate_white_1,device,Sony Cybershot Battery,1
full_crate_black_1,device,USB Hub (4/3 ports),1
half_crate_white_2,device,USB N64,1
half_crate_white_2,device,USB PS2,1
half_crate_white_2,device,USB NES,1
half_crate_white_2,device,USB SNES,1
half_crate_white_2,device,USB Sega,1
half_crate_white_2,device,iPhone,2
half_crate_white_2,device,Raspberry Pi 4,1
half_crate_white_2,device,Raspberry Pi Zero,1
half_crate_white_2,device,Intel Stick Computer,1
half_crate_white_2,device,iPod Color,1
half_crate_white_2,device,External Hard Drive,2
half_crate_white_2,device,Power Bank,2
half_crate_white_2,device,Pendo Box,1
half_crate_white_2,device,Mini Monitor,1
half_crate_orange_1,device,Powered Hub USB-C - x8 USBA,1
half_crate_orange_1,device,Hub USB-C - 4x USB-A,1
half_crate_orange_1,device,Hub USB-C - 3x USB-A 1x RJ45,1
full_crate_stealth_1,device,Intel Nuc,1
full_crate_stealth_1,device,Lenovo Mini PC,2
full_crate_gray_1,device,Dream Router,1
full_crate_gray_1,device,UAP-Beacon HD,1
full_crate_gray_1,device,Swiss Army Knife,1
full_crate_gray_1,device,External Hard Drive Enclosure,1
full_crate_gray_1,device,Server HD 1TB,2
full_crate_gray_1,device,NAS HD 4TB,2
full_crate_gray_1,misc,Server HD Mounts,10
full_crate_gray_2,cable,1/4in Stereo Patch",9
full_crate_gray_2,cable,1/4in Insert Patch",2
full_crate_gray_2,cable,MIDI to USB-A,3
full_crate_gray_2,cable,MIDI to MIDI,2
full_crate_gray_2,cable,1/4in Mono 10ft,1
full_crate_gray_2,cable,1/4in to RCA 2x2,1
full_crate_gray_2,cable,1/8in to RCA 1x2,2
full_crate_gray_2,cable,1/8in to RCA(F) 1x2,1
full_crate_gray_2,cable,1/8in Stereo Patch,5
full_crate_gray_2,cable,1/8in to MIDI Jumper,1
full_crate_gray_2,cable,1/8in to Lightning,1
full_crate_gray_2,device,Shure Beta 58,1
`
