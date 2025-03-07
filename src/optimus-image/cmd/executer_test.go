package cmd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/t0nylombardi/optimus-image/src/optimus-image/internal/mocks"
)

func TestExecutor_Execute(t *testing.T) {
	mockFileUtils := new(mocks.MockFileUtils)
	mockFileOptimizer := new(mocks.MockFileOptimizer)
	mockTracker := new(mocks.MockTracker)

	// Ensure mocks return expected values
	mockFileUtils.On("GetDirectoryPath").Return("/mock/directory")
	mockFileOptimizer.On("OptimizeFiles", mock.Anything).Return(nil)
	mockTracker.On("CompleteFile", mock.Anything)

	executor := NewExecutor(mockFileUtils, mockFileOptimizer, mockTracker)

	t.Run("successful execution", func(t *testing.T) {
		result, err := executor.Execute(func() (string, error) {
			return "Optimized Successfully", nil
		})
		require.NoError(t, err)
		require.Equal(t, "Optimized Successfully", result)
	})

	t.Run("execution error", func(t *testing.T) {
		_, err := executor.Execute(func() (string, error) {
			return "", errors.New("mock error")
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "mock error")
	})

	mockFileUtils.AssertExpectations(t)
	mockFileOptimizer.AssertExpectations(t)
	mockTracker.AssertExpectations(t)
}
