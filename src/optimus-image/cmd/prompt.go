package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// GetUserSelection prompts the user to choose an input type.
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
