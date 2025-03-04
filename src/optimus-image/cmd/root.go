package cmd

import (
	"fmt"

	"github.com/t0nylombardi/optimus-image/src/optimus-image/internal/optimizer"
	"github.com/t0nylombardi/optimus-image/src/optimus-image/internal/utils"

	"github.com/manifoldco/promptui"
)

// GetUserSelection prompts the user to choose between processing a single file or a directory.
func GetUserSelection() (string, error) {
	prompt := promptui.Select{
		Label: "Choose input type",
		Items: []string{"Single File", "Directory"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %v", err)
	}
	return result, nil
}

// Execute runs the CLI workflow based on user selection.
func Execute(getSelectionFunc func() (string, error)) (string, error) {
	// Get user input for input type (Single File / Directory)
	selection, err := getSelectionFunc()
	if err != nil {
		return "", err
	}

	switch selection {
	case "Single File":
		err := processSingleFile()
		return selection, err
	case "Directory":
		err := processDirectory()
		return selection, err
	default:
		return "", fmt.Errorf("invalid selection: %s", selection)
	}
}

// processSingleFile handles the workflow for optimizing a single image file.
func processSingleFile() error {
	// Get file path from the user
	filePath, err := utils.GetFilePath()
	if err != nil {
		return err
	}

	// Run optimization process
	files := []string{filePath}

	if err := optimizer.OptimizeFiles(files); err != nil {
		return err
	}

	fmt.Println("Image optimization complete!")
	return nil
}

// processDirectory handles the workflow for optimizing all image files in a directory.
func processDirectory() error {
	// Get directory path from the user
	dirPath, err := utils.GetDirectoryPath()
	if err != nil {
		return err
	}

	// Run optimization process
	files, err := utils.GetFilesInDirectory(dirPath)
	if err != nil {
		return err
	}

	if err := optimizer.OptimizeFiles(files); err != nil {
		return err
	}

	return nil
}
