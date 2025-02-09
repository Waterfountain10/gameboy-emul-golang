package emulator

import "errors"

// Cartridge represents the Game Boy cartridge, holding the ROM data and any mapper state.
type Cartridge struct {
	ROM []byte
	// For more advanced cartridges, add fields for bank switching, RAM, etc.
}

// NewCartridge creates a new Cartridge instance from the provided ROM data.
func NewCartridge(romData []byte) (*Cartridge, error) {
	if len(romData) == 0 {
		return nil, errors.New("ROM data is empty")
	}
	return &Cartridge{
		ROM: romData,
	}, nil
}
