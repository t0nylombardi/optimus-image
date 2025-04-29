package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/t0nylombardi/optimus-image/optimus-image/cmd"
	"github.com/t0nylombardi/optimus-image/optimus-image/internal/mocks"
	"github.com/t0nylombardi/optimus-image/optimus-image/internal/optimizer"
)

func TestMain_Success(t *testing.T) {
	mockFileUtils := new(mocks.MockFileUtils)
	mockFileOptimizer := new(mocks.MockFileOptimizer)
	mockTracker := new(mocks.MockTracker)

	executor := cmd.NewExecutor(mockFileUtils, mockFileOptimizer, mockTracker)

	mockFileUtils.On("GetFilePath").Return("image.jpg", nil)
	mockFileOptimizer.On("OptimizeFiles", []string{"image.jpg"}, mockTracker).Return(nil)

	t.Run("successful execution", func(t *testing.T) {
		selectionFunc := func(cmd.PromptRunner) (string, error) {
			return "Single File", nil
		}

		result, err := executor.Execute(selectionFunc)
		require.NoError(t, err)
		require.Equal(t, "Single File", result)

		mockFileUtils.AssertExpectations(t)
		mockFileOptimizer.AssertExpectations(t)
	})

	t.Run("execution error", func(t *testing.T) {
		selectionFunc := func(cmd.PromptRunner) (string, error) {
			return "", errors.New("mock error")
		}

		_, err := executor.Execute(selectionFunc)
		require.Error(t, err)
		t.Logf("Actual error: %v", err) // 🔍 Debugging step

		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid selection")
	})
}

func TestFileOptimizer_OptimizeFile(t *testing.T) {
	imageProcessor := optimizer.DefaultImageProcessor{}
	fileOptimizer := optimizer.NewFileOptimizer(&imageProcessor)
	mockTracker := new(mocks.MockTracker) // ✅ Mock tracker

	// ✅ Mock StartSpinner() to prevent unexpected method call error
	mockTracker.On("StartSpinner").Return()
	mockTracker.On("UpdateProgress", mock.Anything, mock.Anything).Return()
	mockTracker.On("CompleteFile", mock.Anything, mock.Anything, mock.Anything).Return()
	mockTracker.On("StopSpinner").Return()

	t.Run("successful optimization", func(t *testing.T) {
		err := fileOptimizer.OptimizeFiles([]string{"image1.jpg", "image2.jpg"}, mockTracker)
		require.NoError(t, err)

		// ✅ Verify that StartSpinner() was called
		mockTracker.AssertExpectations(t)
	})
}
