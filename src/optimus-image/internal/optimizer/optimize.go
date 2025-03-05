package optimizer

import (
	"fmt"
	"sync"

	"github.com/t0nylombardi/optimus-image/src/optimus-image/internal/progress"
)

// FileOptimizerImpl is the default implementation of FileOptimizer.
type FileOptimizerImpl struct {
	processor ImageProcessor // Injected dependency
}

// NewFileOptimizer creates a new optimizer instance with dependencies injected.
func NewFileOptimizer(processor ImageProcessor) *FileOptimizerImpl {
	return &FileOptimizerImpl{processor: processor}
}

// OptimizeFiles processes multiple images concurrently.
func (o *FileOptimizerImpl) OptimizeFiles(images []string, tracker progress.ProgressTracker) error {
	if len(images) == 0 {
		return fmt.Errorf("no images provided for optimization")
	}

	var wg sync.WaitGroup
	tracker.StartSpinner()

	errChan := make(chan error, len(images))

	for _, img := range images {
		wg.Add(1)
		go func(imagePath string) {
			defer wg.Done()
			err := o.processor.Process(imagePath, tracker)
			if err != nil {
				errChan <- err
			}
		}(img)
	}

	wg.Wait()
	close(errChan)
	tracker.StopSpinner()

	var errList []error
	for err := range errChan {
		errList = append(errList, err)
	}
	if len(errList) > 0 {
		return fmt.Errorf("some files failed to optimize: %v", errList)
	}

	return nil
}
