package main

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/aldysp34/deeptech-test/routes"
	"github.com/joho/godotenv"
)

func main() {
	_, currentFile, _, _ := runtime.Caller(0)
	appDir := filepath.Dir(currentFile)

	envPath := filepath.Join(appDir, "../.env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file from %s: %v", envPath, err)
	}

	// Start HTTP server
	addr := "8080"
	r := routes.Routes{Port: addr}
	r.Init()
}
