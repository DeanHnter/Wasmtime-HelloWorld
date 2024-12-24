package main

import (
	"fmt"
	"log"

	"github.com/bytecodealliance/wasmtime-go"
)

func main() {
	// Configure the Wasmtime engine and store
	engine := wasmtime.NewEngine()
	store := wasmtime.NewStore(engine)

	// Convert WAT to Wasm bytes
	wasm, err := wasmtime.Wat2Wasm(`
        (module
            (import "env" "print_string" (func $print_string (param i32 i32)))
            (func (export "print") (param i32 i32)
                local.get 0  ;; string pointer
                local.get 1  ;; string length
                call $print_string
            )
            (memory (export "memory") 1)
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	// Compile the Wasm module
	module, err := wasmtime.NewModule(engine, wasm)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new linker
	linker := wasmtime.NewLinker(engine)

	// Define the function import
	err = linker.DefineFunc(store, "env", "print_string",
		func(caller *wasmtime.Caller, ptr int32, length int32) {
			memory := caller.GetExport("memory").Memory()
			data := memory.UnsafeData(store)
			str := string(data[ptr : ptr+length])
			fmt.Println("Wasm says:", str)
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Instantiate the module
	instance, err := linker.Instantiate(store, module)
	if err != nil {
		log.Fatal(err)
	}

	// Get the memory export
	memory := instance.GetExport(store, "memory").Memory()

	// Function to write string to WebAssembly memory and call print function
	writeAndPrint := func(s string) error {
		// Get the print function
		printFunc := instance.GetFunc(store, "print")
		if printFunc == nil {
			return fmt.Errorf("failed to find function export 'print'")
		}

		// Write the string to memory
		data := memory.UnsafeData(store)
		stringBytes := []byte(s)
		copy(data[0:], stringBytes)

		// Call the print function with pointer (0) and length of the string
		_, err := printFunc.Call(store, 0, len(stringBytes))
		return err
	}

	// Test with different strings
	strings := []string{
		"Hello from Go!",
		"This is a test",
		"WebAssembly is cool!",
	}

	for _, s := range strings {
		if err := writeAndPrint(s); err != nil {
			log.Fatal(err)
		}
	}
}
