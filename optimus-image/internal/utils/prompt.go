package utils

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// GetFilePath prompts the user to enter the image file path
func (f *FileUtilsImpl) GetFilePath() (string, error) {
	prompt := promptui.Prompt{
		Label: "Enter the path to the image file",
	}

	filePath, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get file path: %v", err)
	}

	if !IsValidImage(filePath) {
		return "", fmt.Errorf("invalid file type: %s", filePath)
	}

	return filePath, nil
}

// GetDirectoryPath prompts the user to enter the directory path
func (f *FileUtilsImpl) GetDirectoryPath() (string, error) {
	prompt := promptui.Prompt{
		Label: "Enter the path to the directory",
	}

	dirPath, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get directory path: %v", err)
	}

	return dirPath, nil
}

// AskThumbnailOption prompts the user if they want to generate a thumbnail
func AskThumbnailOption() (bool, error) {
	prompt := promptui.Select{
		Label: "Do you want to create a thumbnail?",
		Items: []string{"Yes", "No"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return false, fmt.Errorf("failed to get thumbnail option: %v", err)
	}
	return result == "Yes", nil
}

// AskOverwriteOption prompts the user if they want to overwrite or rename
func AskOverwriteOption() (bool, error) {
	prompt := promptui.Select{
		Label: "Do you want to overwrite the original file?",
		Items: []string{"Yes", "No, rename it"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return false, fmt.Errorf("failed to get overwrite option: %v", err)
	}
	return result == "Yes", nil
}

// AskSaveLocation prompts the user to choose where to save the optimized image
func AskSaveLocation() (string, error) {
	prompt := promptui.Select{
		Label: "Where do you want to save the optimized image?",
		Items: []string{"Same location", "Different location"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get save location: %v", err)
	}

	if result == "Different location" {
		prompt := promptui.Prompt{
			Label: "Enter the destination folder",
		}
		destination, err := prompt.Run()
		if err != nil {
			return "", fmt.Errorf("failed to get destination folder: %v", err)
		}
		return destination, nil
	}

	return "Same location", nil
}
