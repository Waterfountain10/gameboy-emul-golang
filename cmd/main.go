package main

import (
	"fmt"
	"log"
	"os"

	"gameboy-emul-golang/pkg/emulator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-gameboy-emulator [ROM file]")
		os.Exit(1)
	}
	romPath := os.Args[1]

	emu, err := emulator.NewEmulator(romPath)
	if err != nil {
		log.Fatalf("Failed to initialize emulator: %v", err)
	}

	emu.Run() // start the emulation loop
}
