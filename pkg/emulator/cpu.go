package emulator

// CPU represents the Game Boy’s CPU, including registers and a pointer to memory.
type CPU struct {
	Memory *Memory
	// Registers (for example purposes, not complete)
	A, F, B, C, D, E, H, L uint8
	PC, SP                 uint16
}

// NewCPU initializes a new CPU instance.
func NewCPU(memory *Memory) *CPU {
	// Initialize registers to the Game Boy’s default startup values as needed.
	return &CPU{
		Memory: memory,
		PC:     0x0100, // Typically, the Game Boy starts execution at 0x0100.
		SP:     0xFFFE,
	}
}

// Step emulates a single CPU instruction. Returns the number of cycles taken.
func (cpu *CPU) Step() int {
	opcode := cpu.Memory.Read(cpu.PC)
	// Decode and execute the opcode.
	// Here we only do a simple placeholder implementation.
	// You would implement a full instruction decoder and executor here.
	cpu.PC++ // Advance program counter for this example.
	return 4 // Assume each instruction takes 4 cycles (placeholder).
}
