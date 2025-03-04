package optimizer

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"github.com/t0nylombardi/optimus-image/src/optimus-image/internal/utils"
)

// OptimizeFiles processes a single file or multiple images with stacked progress bars.
func OptimizeFiles(images []string) error {
	if len(images) == 0 {
		return fmt.Errorf("no images provided for optimization")
	}

	var wg sync.WaitGroup
	tracker := utils.NewTracker(len(images))
	tracker.StartSpinner()

	// Channel for errors
	errChan := make(chan error, len(images))

	// Process images concurrently
	for _, img := range images {
		wg.Add(1)
		go func(imagePath string) {
			defer wg.Done()

			// Simulate optimization
			err := optimizeImage(imagePath, tracker)
			if err != nil {
				errChan <- err
			}
		}(img)
	}

	wg.Wait()
	close(errChan) // Close error channel when done
	tracker.StopSpinner()

	// Collect any errors
	var errList []error
	for err := range errChan {
		errList = append(errList, err)
	}
	if len(errList) > 0 {
		return fmt.Errorf("some files failed to optimize: %v", errList)
	}

	// Show total savings (only for multiple images)
	if len(images) > 1 {
		displayTotalSavings(images)
	}

	return nil
}

// optimizeImage handles the processing of a single image.
func optimizeImage(imagePath string, tracker *utils.ProgressTracker) error {
	// Simulate optimization progress
	for i := 0; i <= 100; i += 10 {
		time.Sleep(200 * time.Millisecond)
		tracker.UpdateProgress(imagePath, i)
	}

	// Simulate final optimization result
	originalSize, optimizedSize := calculateOptimization(imagePath)
	tracker.CompleteFile(filepath.Base(imagePath), originalSize, optimizedSize)

	return nil
}

// calculateOptimization simulates the image size reduction.
func calculateOptimization(imagePath string) (int64, int64) {
	originalSize := int64(1024 * (1 + len(imagePath)%5)) // Mock file size in KB
	reductionPercent := 30 + (len(imagePath) % 50)       // Mock percentage
	optimizedSize := originalSize * (100 - int64(reductionPercent)) / 100
	return originalSize, optimizedSize
}

// displayTotalSavings calculates and prints the total space saved.
func displayTotalSavings(images []string) {
	var totalOriginalSize, totalOptimizedSize int64

	for _, img := range images {
		originalSize, optimizedSize := calculateOptimization(img)
		totalOriginalSize += originalSize
		totalOptimizedSize += optimizedSize
	}

	totalReduction := float64(100) - (float64(totalOptimizedSize)/float64(totalOriginalSize))*100
	fmt.Printf("\n✨ Total Saved: %.1f%% (%d KB) 🎉\n", totalReduction, totalOriginalSize-totalOptimizedSize)
}
