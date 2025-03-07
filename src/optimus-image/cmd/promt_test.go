package cmd

import (
	"errors"
	"testing"

	"github.com/manifoldco/promptui"
	"github.com/stretchr/testify/assert"
)

func TestGetUserSelection_SingleFile(t *testing.T) {
	mockRunner := func(prompt promptui.Select) (int, string, error) {
		return 0, "Single File", nil
	}

	result, err := GetUserSelection(mockRunner)
	assert.NoError(t, err)
	assert.Equal(t, "Single File", result)
}

func TestGetUserSelection_Directory(t *testing.T) {
	mockRunner := func(prompt promptui.Select) (int, string, error) {
		return 1, "Directory", nil
	}

	result, err := GetUserSelection(mockRunner)
	assert.NoError(t, err)
	assert.Equal(t, "Directory", result)
}

func TestGetUserSelection_PromptError(t *testing.T) {
	mockRunner := func(prompt promptui.Select) (int, string, error) {
		return 0, "", errors.New("mocked prompt failure")
	}

	result, err := GetUserSelection(mockRunner)
	assert.Error(t, err)
	assert.Equal(t, "", result)
	assert.Contains(t, err.Error(), "mocked prompt failure")
}
