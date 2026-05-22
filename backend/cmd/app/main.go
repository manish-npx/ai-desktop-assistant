package main

import (
	"context"
	"fmt"
	"log"

	"github.com/manish-npx/ai-desktop-assistant/backend/internal/logging"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

var (
	appVersion = "0.1.0"
	appName    = "AI Desktop Assistant"
)

func main() {
	// Initialize logging
	logger, err := logging.New(logging.Config{
		Level:  "info",
		Format: "json",
	})
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	logger.Infof("Starting %s v%s", appName, appVersion)

	// Create application with options
	err = createApp(logger)
	if err != nil {
		logger.Fatalf("Error creating application: %v", err)
	}
}

func createApp(logger *logging.Logger) error {
	// Create context
	ctx := context.Background()

	// Create the application with options
	err := wails.Run(&options.App{
		Title:             appName,
		Width:             1200,
		Height:            800,
		MinWidth:          400,
		MinHeight:         300,
		MaxWidth:          2560,
		MaxHeight:         1440,
		DisableFrameless:  false,
		Transparent:       false,
		BackgroundColour:  &options.RGBA{R: 27, G: 27, B: 27, A: 255},
		OnStartup:         func(ctx context.Context) { onStartup(ctx, logger) },
		OnDomReady:        func(ctx context.Context) { onDomReady(ctx, logger) },
		OnShutdown:        func(ctx context.Context) { onShutdown(ctx, logger) },
		OnBeforeClose:     func(ctx context.Context) (prevent bool) { return false },
		StartHidden:       false,
		HideWindowOnClose: false,
		EnableModals:      true,
	})

	if err != nil {
		return fmt.Errorf("failed to run application: %w", err)
	}

	return nil
}

func onStartup(ctx context.Context, logger *logging.Logger) {
	logger.Infof("Application startup - context initialized")
	// Initialize services, database, etc.
}

func onDomReady(ctx context.Context, logger *logging.Logger) {
	logger.Infof("DOM ready - UI initialized")
	// Setup event listeners, etc.
}

func onShutdown(ctx context.Context, logger *logging.Logger) {
	logger.Infof("Application shutting down")
	// Cleanup resources
}
