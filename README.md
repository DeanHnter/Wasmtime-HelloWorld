# Wasm with Go and Wasmtime: Argument Passing Example

This project demonstrates how to use [Wasmtime](https://github.com/bytecodealliance/wasmtime) with Go to pass arguments between the Go runtime and a WebAssembly (Wasm) module, which is defined using WebAssembly Text Format (WAT). The example showcases how you can pass strings from Go to a WebAssembly function, and have them printed using a host-defined function.

## Prerequisites

- Go 1.16 or later installed
- [Wasmtime](https://github.com/bytecodealliance/wasmtime-go) Go package installed

You can install the Wasmtime Go package using:

```sh
go get github.com/bytecodealliance/wasmtime-go
```

## Project Structure

- `main.go`: The main program that sets up the Wasm environment, passes arguments from Go to Wasm, and prints them using a Wasm-hosted function.

## Features

- **Argument Passing**: Demonstrates how to pass string arguments from Go to Wasm.
- **WAT to Wasm Compilation**: Shows how to compile WAT (WebAssembly Text Format) to Wasm bytes.
- **Function Import/Export**: Includes a host function that is imported into the Wasm module, and a Wasm function that is called from Go.

## How It Works

1. **Wasmtime Setup**:
    - Create a new engine and store for managing Wasm execution.
    - Compile WAT into Wasm bytes and then create a Wasm module.

2. **Function Import**:
    - Define an `env.print_string` function in Go which is hooked to print strings passed from Wasm.

3. **Memory Management**:
    - Allocate memory in Wasm and manipulate it directly from Go to write strings.

4. **Wasm Function Invocation**:
    - Call an exported Wasm function with a string's memory location and length to trigger the host-defined `print_string`.

5. **Example Execution**:
    - Test the setup by passing a series of strings from Go to the Wasm module, printing each using the defined function.

## Running the Example

1. Clone this repository:

   ```sh
   git clone <repository-url>
   cd wasm-with-go-wasmtime
   ```

2. Execute the example:

   ```sh
   go run main.go
   ```

3. The output should display:

   ```
   Wasm says: Hello from Go!
   Wasm says: This is a test
   Wasm says: WebAssembly is cool!
   ```

## License

This project is open source and available under the MIT License.

---

This tutorial should serve as a basic guide to building more advanced WebAssembly applications using Wasmtime and Go. Feel free to expand on this example by exploring more complex data types, error handling, and interacting with other WebAssembly modules.