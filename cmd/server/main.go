package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Default values
	path := "public"
	port := 8080

	// Check if environment variables are set for path and port
	envPath := os.Getenv("FILE_SERVER_PATH")
	if envPath != "" {
		path = envPath
	}

	envPort := os.Getenv("FILE_SERVER_PORT")
	if envPort != "" {
		envPortInt, err := strconv.Atoi(envPort)
		if err == nil {
			port = envPortInt
		}
	}

	// Start server
	fmt.Printf("Starting file server at %s on port %d\n", path, port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(path)))
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}
