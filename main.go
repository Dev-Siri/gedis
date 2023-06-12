package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Dev-Siri/gedis/core"
	"github.com/Dev-Siri/gedis/db"
	"github.com/Dev-Siri/gedis/embeds"
)

//go:embed pages
var PagesFS embed.FS

//go:embed public
var PublicFS embed.FS

func main() {
	log.Println("Starting Gedis server")

	port := os.Getenv("PORT")

	if port == "" {
		log.Printf("WARN: Port environment variable not found. Using 5000 instead.")
		port = "5000"
	}

	embeds.Pages = PagesFS
	embeds.StaticFiles = PublicFS

	go core.StartServer(port)
	log.Printf("Gedis server listening on port %s\n", port)
	
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	
	go func() {
		<-interrupt
		fmt.Println("\n(Info) Shutting down Gedis server")
		
		os.Exit(0)
	}()

	if err := db.InitializeDB(); err != nil {
		log.Fatalf("Database initialization failed")
		return
	}

	// Runs every 5 seconds and checks the db
	// for expired values & other stuff then deletes it
	go db.CleanDB()
	core.RunCLI()
}