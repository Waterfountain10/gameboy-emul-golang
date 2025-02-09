package emulator

// Memory represents the Game Boy’s memory, including RAM and memory-mapped I/O.
type Memory struct {
	data []byte
	cart *Cartridge
}

// NewMemory creates a new memory instance and initializes the memory map.
func NewMemory(cart *Cartridge) *Memory {
	mem := &Memory{
		data: make([]byte, 0x10000), // 64KB of addressable space.
		cart: cart,
	}
	// Additional initialization (e.g., I/O registers, VRAM) can be done here.
	return mem
}

// Read returns the byte at the given address.
func (m *Memory) Read(addr uint16) byte {
	// For this simple example, if the address is within ROM, return from cartridge.
	if addr < 0x8000 {
		return m.cart.ROM[addr]
	}
	return m.data[addr]
}

// Write sets the byte at the given address.
func (m *Memory) Write(addr uint16, value byte) {
	// Prevent writes to ROM areas.
	if addr < 0x8000 {
		return
	}
	m.data[addr] = value
}
