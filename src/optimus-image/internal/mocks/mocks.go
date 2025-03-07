package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/t0nylombardi/optimus-image/src/optimus-image/internal/progress"
)

// ✅ MockFileUtils implements file.FileUtils
type MockFileUtils struct {
	mock.Mock
}

func (m *MockFileUtils) GetDirectoryPath() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockFileUtils) GetFilePath() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockFileUtils) GetFilesInDirectory(path string) ([]string, error) {
	args := m.Called(path)
	return args.Get(0).([]string), args.Error(1)
}

// ✅ MockFileOptimizer implements optimizer.FileOptimizer correctly
type MockFileOptimizer struct {
	mock.Mock
}

func (m *MockFileOptimizer) OptimizeFile(path string) error {
	args := m.Called(path)
	return args.Error(0)
}

func (m *MockFileOptimizer) OptimizeFiles(images []string, tracker progress.ProgressTracker) error {
	args := m.Called(images, tracker)
	return args.Error(0)
}

// ✅ MockTracker implements progress.ProgressTracker
type MockTracker struct {
	mock.Mock
}

func (m *MockTracker) StartSpinner() {
	m.Called()
}

func (m *MockTracker) StopSpinner() {
	m.Called()
}

func (m *MockTracker) UpdateProgress(file string, percent int) {
	m.Called(file, percent)
}

func (m *MockTracker) CompleteFile(file string, originalSize, optimizedSize int64) {
	m.Called(file, originalSize, optimizedSize)
}

// ✅ MockExecutor implements cmd.Executor
type MockExecutor struct {
	mock.Mock
}

func (m *MockExecutor) Execute(selectionFunc func() (string, error)) (string, error) {
	args := m.Called(selectionFunc)
	return args.String(0), args.Error(1)
}
