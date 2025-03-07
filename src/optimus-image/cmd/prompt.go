package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// PromptRunner defines a function signature for running a prompt.
type PromptRunner func(promptui.Select) (int, string, error)

// GetUserSelection prompts the user to choose an input type.
func GetUserSelection(runPrompt PromptRunner) (string, error) {
	prompt := promptui.Select{
		Label: "Choose input type",
		Items: []string{"Single File", "Directory"},
	}

	_, result, err := runPrompt(prompt)
	if err != nil {
		return "", fmt.Errorf("prompt failed: %v", err)
	}
	return result, nil
}
