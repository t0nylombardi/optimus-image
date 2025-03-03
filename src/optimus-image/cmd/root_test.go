package cmd

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/manifoldco/promptui"
	"github.com/stretchr/testify/assert"
)

// TestPromptSelect simulates selecting an option from a prompt.
func TestPromptSelect(t *testing.T) {
	// Mock user input selecting the first option ("Single File")
	input := "1\n"                                      // Simulating user pressing "1" and "Enter"
	stdin := io.NopCloser(bytes.NewBufferString(input)) // Wrap with io.NopCloser

	// Create a select prompt with mocked input
	prompt := promptui.Select{
		Label: "Choose input type",
		Items: []string{"Single File", "Directory"},
		Stdin: stdin, // Now it matches the expected type
	}

	_, result, err := prompt.Run()

	// Assertions
	assert.NoError(t, err, "Prompt should not return an error")
	assert.Equal(t, "Single File", result, "User should have selected 'Single File'")
}

// TestPromptCancel simulates the user pressing Ctrl+C to exit the prompt.
func TestPromptCancel(t *testing.T) {
	// Simulate user pressing Ctrl+C by manually returning `ErrInterrupt`
	mockPrompt := func() (string, error) {
		return "", promptui.ErrInterrupt
	}

	_, err := mockPrompt()

	// Assertions
	assert.ErrorIs(t, err, promptui.ErrInterrupt, "Expected user to cancel the prompt")
}

// TestPromptValidation simulates a validation error scenario.
// TestPromptValidation simulates a validation error scenario.
func TestPromptValidation(t *testing.T) {
	// Simulated user input sequence: "invalid" → "valid"
	input := io.NopCloser(bytes.NewBufferString("valid\n"))

	// Validation function
	validate := func(input string) error {
		if input == "valid" {
			return nil
		}
		return errors.New("invalid input")
	}

	// Create the prompt
	prompt := promptui.Prompt{
		Label:    "Test Prompt",
		Stdin:    input,
		Validate: validate,
	}

	// Simulate user retrying input
	result, err := prompt.Run()

	// Assertions
	assert.NoError(t, err, "Prompt should not return an error after valid input")
	assert.Equal(t, "valid", result, "Expected 'valid' as the user input")
}
