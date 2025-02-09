package emulator

import (
	"io/ioutil"
)

type Emulator struct {
	CPU       *CPU
	GPU       *GPU
	Memory    *Memory
	APU       *APU
	Cartridge *Cartridge
	// Future fields: Input, Timer, etc.
}

// NewEmulator loads the ROM and sets up the emulator components.
func NewEmulator(romPath string) (*Emulator, error) {
	romData, err := ioutil.ReadFile(romPath)
	if err != nil {
		return nil, err
	}

	cart, err := NewCartridge(romData)
	if err != nil {
		return nil, err
	}

	mem := NewMemory(cart)
	cpu := NewCPU(mem)
	gpu := NewGPU(mem)
	apu := NewAPU(mem)

	return &Emulator{
		CPU:       cpu,
		GPU:       gpu,
		Memory:    mem,
		APU:       apu,
		Cartridge: cart,
	}, nil
}

// Run starts the main emulation loop.
func (e *Emulator) Run() {
	// A very basic loop: in a real emulator you might want to implement timing,
	// frame limiting, and event handling.
	for {
		// Execute a CPU instruction and get the number of cycles consumed.
		cycles := e.CPU.Step()

		// Advance the GPU and APU state by the number of cycles.
		e.GPU.Step(cycles)
		e.APU.Step(cycles)

		// Input handling and other periodic updates would go here.
	}
}
