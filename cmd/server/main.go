package main

import (
	"goblossom/api"
	"goblossom/config"
	"goblossom/internal/blob"
	"goblossom/pkg/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load configuration from config.yaml
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Configure logging
	utils.SetLogLevel(cfg.Logging.Level)

	// Initialize the BlobRepository based on storage type
	var blobRepo blob.BlobRepository
	if cfg.Storage.Type == "file" {
		blobRepo = blob.NewFileBlobRepository(cfg.Storage.FilePath)
	} else {
		log.Fatalf("Unsupported storage type: %s", cfg.Storage.Type)
	}

	// Create the handler implementation
	handler := api.HandlerImpl{
		BlobRepo: blobRepo,
	}

	// Initialize the router with the handler
	router := api.NewRouter(handler)

	// Start the server on the configured port
	log.Printf("Starting server on port %d...\n", cfg.Server.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func watchConfigReload(configFile string) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGHUP)

	for {
		<-signals
		log.Println("Reloading configuration...")
		cfg, err := config.LoadConfig(configFile)
		if err != nil {
			log.Printf("Failed to reload config: %v", err)
		} else {
			// Apply the new configuration dynamically
			currentConfig = cfg
		}
	}
}
