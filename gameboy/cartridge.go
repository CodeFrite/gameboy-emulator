package gameboy

import (
	"errors"
	"fmt"
	"os"
)

type Cartridge struct {
	cartridgePath string
	cartridgeName string
	rom []byte
	header []byte
	entry_point []byte
	nintendo_logo []byte
	title []byte
	manufacturer_code []byte
	cgb_flag []byte
	new_licensee_code []byte
	sgb_flag []byte
	cartridge_type []byte
	rom_size []byte
	ram_size []byte
	destination_code []byte
	old_licensee_code []byte
	mask_rom_version []byte
	header_checksum []byte
	global_checksum []byte
}

func NewCartridge(uri string, name string) *Cartridge {
	var c Cartridge
	err := c.loadRom(uri + "/" + name)
	if err != nil {
		fmt.Println("Error loading ROM:", err)
		return nil
	}
	c.cartridgePath = uri
	c.cartridgeName = name
	c.parseHeader()
	return &c
}

func (c *Cartridge) loadRom(uri string) error {
	var err error
	c.rom, err = os.ReadFile(uri)
	if err != nil {
		errText := fmt.Sprint("Error loading ROM:", err)
		return errors.New(errText)
	}
	return nil
}

func (c *Cartridge) parseHeader() {
	c.header = c.rom[0x0100:0x014F]
	c.entry_point = c.rom[0x0100:0x0103]
	c.nintendo_logo = c.rom[0x0104:0x0133]
	c.title = c.rom[0x0134:0x0143]
	c.manufacturer_code = c.rom[0x013F:0x0142]
	c.cgb_flag = c.rom[0x0143:0x0144]
	c.new_licensee_code = c.rom[0x0144:0x0145]
	c.sgb_flag = c.rom[0x0146:0x0147]
	c.cartridge_type = c.rom[0x0147:0x0148]
	c.rom_size = c.rom[0x0148:0x0149]
	c.ram_size = c.rom[0x0149:0x014A]
	c.destination_code = c.rom[0x014A:0x014B]
	c.old_licensee_code = c.rom[0x014B:0x014C]
	c.mask_rom_version = c.rom[0x014C:0x014D]
	c.header_checksum = c.rom[0x014D:0x014E]
	c.global_checksum = c.rom[0x014E:0x0150]
}

func (c *Cartridge) PrintInfo() {
	// metadata about the cartridge
	fmt.Println("Cartridge Path:", c.cartridgePath)
	fmt.Println("Cartridge Name:", c.cartridgeName)
	fmt.Println("Cartridge Size:", len(c.rom), "bytes")

	// header information
	fmt.Println("Header:", c.header)
	fmt.Println("Entry Point:", c.entry_point)
	fmt.Println("Nintendo Logo:", c.nintendo_logo)
	fmt.Println("Title:", string(c.title))
	fmt.Println("Manufacturer Code:", c.manufacturer_code)
	fmt.Println("CGB Flag:", c.cgb_flag)
	fmt.Println("New Licensee Code:", c.new_licensee_code)
	fmt.Println("SGB Flag:", c.sgb_flag)
	fmt.Println("Cartridge Type:", c.cartridge_type)
	fmt.Println("ROM Size:", c.rom_size)
	fmt.Println("RAM Size:", c.ram_size)
	fmt.Println("Destination Code:", c.destination_code)
	fmt.Println("Old Licensee Code:", c.old_licensee_code)
	fmt.Println("Mask ROM Version:", c.mask_rom_version)
	fmt.Println("Header Checksum:", c.header_checksum)
	fmt.Println("Global Checksum:", c.global_checksum)
}

func (c *Cartridge) Read(addr [2]byte) byte {
	addr16 := bytesToUint16(addr)
	return c.rom[addr16]
}

func (c *Cartridge) Write(addr [2]byte, value byte) {
	addr16 := bytesToUint16(addr)
	c.rom[addr16] = value
}