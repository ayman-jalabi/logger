package main

import (
	"fmt"
)

func main() {
	// Create a new Logger instance
	logger, err := NewLogger(DebugLevel, true, "app.log")
	if err != nil {
		fmt.Println("Error creating logs:", err)
		return
	}
	defer logger.Close()

	// Example usage
	logger.Info("Application started")
	logger.Warning("This is a warning message")
	logger.Error("An error occurred")
	logger.Debug("Debugging details")
}
