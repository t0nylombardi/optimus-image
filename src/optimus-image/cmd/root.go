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
		fmt.Println("Directory processing is not yet implemented.")
		return selection, nil
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

	// // Ask user if they want to overwrite or rename the optimized file
	// overwrite, err := utils.AskOverwriteOption()
	// if err != nil {
	// 	return err
	// }

	// // Ask user where they want to save the optimized image
	// saveLocation, err := utils.AskSaveLocation()
	// if err != nil {
	// 	return err
	// }

	// Run optimization process
	files := []string{filePath}

	if err := optimizer.OptimizeFiles(files); err != nil {
		return err
	}

	fmt.Println("Image optimization complete!")
	return nil
}
