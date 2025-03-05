package cmd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFileUtils struct {
	mock.Mock
}

func (m *MockFileUtils) GetFilePath() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockFileUtils) GetDirectoryPath() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockFileUtils) GetFilesInDirectory(dirPath string) ([]string, error) {
	args := m.Called(dirPath)
	return args.Get(0).([]string), args.Error(1)
}

type MockFileOptimizer struct {
	mock.Mock
}

func (m *MockFileOptimizer) OptimizeFiles(files []string) error {
	args := m.Called(files)
	return args.Error(0)
}

func TestExecute_SingleFile(t *testing.T) {
	// Creating mocks for file utils and optimizer
	mockFileUtils := new(MockFileUtils)
	mockFileOptimizer := new(MockFileOptimizer)

	// Mocking the GetUserSelection function
	mockGetSelection := func() (string, error) {
		return "Single File", nil
	}

	// Mocking the behavior of GetFilePath
	mockFileUtils.On("GetFilePath").Return("test.jpg", nil)
	// Mocking the behavior of OptimizeFiles
	mockFileOptimizer.On("OptimizeFiles", []string{"test.jpg"}).Return(nil)

	// Creating the executor with the mocked dependencies
	executor := &Executor{
		FileUtils:     mockFileUtils,
		FileOptimizer: mockFileOptimizer,
	}

	// Execute and test
	result, err := executor.Execute(mockGetSelection)
	assert.NoError(t, err)
	assert.Equal(t, "Single File", result)

	// Assert that the mocks were called as expected
	mockFileUtils.AssertExpectations(t)
	mockFileOptimizer.AssertExpectations(t)
}

func TestProcessSingleFile_Failure(t *testing.T) {
	// Creating mocks for file utils
	mockFileUtils := new(MockFileUtils)

	// Mocking GetFilePath to return an error
	mockFileUtils.On("GetFilePath").Return("", errors.New("failed to get file path"))

	// Creating the executor with the mocked dependencies
	executor := &Executor{
		FileUtils: mockFileUtils,
	}

	// Test failure in processing a single file
	err := executor.processSingleFile()
	assert.Error(t, err)
	assert.Equal(t, "failed to get file path", err.Error())

	mockFileUtils.AssertExpectations(t)
}
