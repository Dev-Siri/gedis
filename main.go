package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Dev-Siri/gedis/core"
	"github.com/Dev-Siri/gedis/db"
)

func main() {
	log.Println("Starting Gedis server")

	port := os.Getenv("PORT")

	if port == "" {
		log.Printf("WARN: Port environment variable not found. Using 5000 instead.")
		port = "5000"
	}

	go core.StartServer(port)
	log.Printf("Gedis server listening on port %s\n", port)
	
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	
	go func() {
		<-interrupt
		fmt.Println("\nShutting down Gedis server")
		
		os.Exit(0)
	}()
		
	if err := db.InitializeDB(); err != nil {
		log.Fatalf("Database initialization failed")
		return
	}

	core.RunCLI()
}