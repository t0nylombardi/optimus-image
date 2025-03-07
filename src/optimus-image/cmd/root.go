package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/t0nylombardi/optimus-image/src/optimus-image/internal/file"
	"github.com/t0nylombardi/optimus-image/src/optimus-image/internal/optimizer"
	"github.com/t0nylombardi/optimus-image/src/optimus-image/internal/progress"
)

// ExecutorImpl handles CLI execution.
type ExecutorImpl struct {
	FileUtils     file.FileUtils
	FileOptimizer optimizer.FileOptimizer
	Tracker       progress.ProgressTracker
}

// NewExecutor creates a new ExecutorImpl instance.
func NewExecutor(fileUtils file.FileUtils, fileOptimizer optimizer.FileOptimizer, tracker progress.ProgressTracker) *ExecutorImpl {
	return &ExecutorImpl{
		FileUtils:     fileUtils,
		FileOptimizer: fileOptimizer,
		Tracker:       tracker,
	}
}

// Execute runs the CLI workflow based on user selection.
func (e *ExecutorImpl) Execute(getSelectionFunc func(PromptRunner) (string, error)) (string, error) {
	selection, _ := getSelectionFunc(func(p promptui.Select) (int, string, error) {
		return p.Run()
	})

	switch selection {
	case "Single File":
		err := e.processSingleFile()
		return selection, err
	case "Directory":
		err := e.processDirectory()
		return selection, err
	default:
		return "", fmt.Errorf("invalid selection: %s", selection)
	}
}

// processSingleFile optimizes a single image file.
func (e *ExecutorImpl) processSingleFile() error {
	filePath, err := e.FileUtils.GetFilePath()
	if err != nil {
		return err
	}

	files := []string{filePath}
	if err := e.FileOptimizer.OptimizeFiles(files, e.Tracker); err != nil {
		return err
	}

	fmt.Println("✅ Image optimization complete!")
	return nil
}

// processDirectory optimizes all images in a directory.
func (e *ExecutorImpl) processDirectory() error {
	dirPath, err := e.FileUtils.GetDirectoryPath()
	if err != nil {
		return err
	}

	files, err := e.FileUtils.GetFilesInDirectory(dirPath)
	if err != nil {
		return err
	}

	if err := e.FileOptimizer.OptimizeFiles(files, e.Tracker); err != nil {
		return err
	}

	fmt.Println("✅ Directory optimization complete!")
	return nil
}
