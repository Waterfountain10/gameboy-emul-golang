# gameboy-emul-golang
Gameboy Emulator written entirely in Go.

this project was designed to help me learn Golang.

I started it a couple months ago, but only recently continued it and put it on github.

## it is not working yet.

here is a project overview :

### cmd/main.go
- The entry point for your emulator. This file parses command-line arguments (such as the ROM path), initializes the emulator, and starts the emulation loop.

### pkg/emulator/

This package contains the core components of your emulator:

- emulator.go: The top-level struct that ties all the components together and contains the main emulation loop.
- cpu.go: Contains the CPU implementation. In a Game Boy, this is a variant of the Z80 processor.
- memory.go: Implements the memory map and handles reading/writing data (including integration with the cartridge).
- gpu.go: Handles the graphics processing unit, responsible for rendering the Game Boy’s display.
- apu.go: (Optional, at first) Implements sound emulation.
- cartridge.go: Loads and manages the ROM, including bank switching logic if you plan to support more than basic games.


## Building and Running

1. Build the emulator:
   ```bash
   go build -o gameboy-emulator ./cmd
   ```

2. Run the emulator with a downloadwd ROM
  ```bash
  ./gameboy-emulator path/to/your/game.rom
  ```
