# Unverity: Verity Dissections Tool

## Info

- Unverity is a website used to generate a solution for the Verity encounter dissections for the Salvation's Edge raid in Destiny 2.
- The tool works by recursively brute forcing all possible dissections up to a certain max amount of dissections and returning the shortest solution found.
- The first solution is returned if multiple solutions of the same length are found.
- If a solution cannot be found, in the case of the max search depth being too short, no solution existing, or invalid input being given, an error message will be returned instead.
- The core logic is programmed in Go and built into WebAssembly (Wasm) to be executed on the browser.

# Building and Running

## Prerequisites

- Go 1.20

## Building the wasm file

- Build the .wasm file using the command `go build -o public/wasm/dissect.wasm unverity/cmd/wasm` in the root project directory

- Note: Make sure the environment variables `GOARCH=wasm` and `GOOS=js` are set

## Running the server

- Build the server using the command `go build unverity/cmd/server` in the root project directory
- Run the server using the command `server` in the root project directory

- Note: Make sure the environment variables `GOARCH` and `GOOS` are set and correct for your Architecture and OS

Environment Variables:
- `FILE_SERVER_PATH`: Set the fileserver path. This should be the directory of the public folder. Default `public`.
- `FILE_SERVER_PORT`: The port of the fileserver. Default 8080.

## Updating the `wasm_exec.js` file

- The `wasm_exec.js` is a runtime support script for running Go-compiled WebAssembly modules in web browsers, found in the Go source directory.
- To update the `wasm_exec.js` file, copy the file from the Go source directory to the project:

```sh
cp $(go env GOROOT)/misc/wasm/wasm_exec.js /path/to/your/project/static/wasm/
```
