package main

import (
	"fmt"
	"log"

	"github.com/t0nylombardi/optimus-image/optimus-image/cmd"
	"github.com/t0nylombardi/optimus-image/optimus-image/internal/file"
	"github.com/t0nylombardi/optimus-image/optimus-image/internal/optimizer"
	"github.com/t0nylombardi/optimus-image/optimus-image/internal/progress"
)

func main() {
	fmt.Println("🚀 Optimus Image - Image Optimization CLI")

	// Initialize dependencies
	fileUtils := file.NewFileUtils()
	imageProcessor := optimizer.DefaultImageProcessor{}
	fileOptimizer := optimizer.NewFileOptimizer(&imageProcessor)
	tracker := progress.NewTracker(1)

	executor := cmd.NewExecutor(fileUtils, fileOptimizer, tracker)

	// Run CLI execution
	selection, err := executor.Execute(cmd.GetUserSelection)
	if err != nil {
		log.Fatalf("❌ Error: %v", err)
	}

	fmt.Printf("✨ Operation completed: %s\n", selection)
}
