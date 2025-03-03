package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// Execute starts the CLI
func Execute() error {
	fmt.Println("Welcome to Optimus-Image! 🚀")

	// Prompt for file or directory input
	prompt := promptui.Select{
		Label: "Choose input type",
		Items: []string{"Single File", "Directory"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %v", err)
	}

	fmt.Printf("You selected: %s\n", result)
	return nil
}
