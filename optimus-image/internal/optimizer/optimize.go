package optimizer

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/h2non/bimg"
	"github.com/t0nylombardi/optimus-image/optimus-image/internal/utils"
)

type FileOptimizer interface {
	OptimizeFiles(images []string) error
}

type FileOptimizerImpl struct{}

// OptimizeFiles processes a single file or multiple images with stacked progress bars.
func (o *FileOptimizerImpl) OptimizeFiles(images []string) error {
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
		displayTotalSavings(images, tracker)
	}

	return nil
}

// optimizeImage handles the processing of a single image.
func optimizeImage(imagePath string, tracker *utils.ProgressTracker) error {
	tracker.UpdateProgress(imagePath, 10) // Start

	buffer, err := os.ReadFile(imagePath)
	if err != nil {
		return fmt.Errorf("failed to read image %s: %w", imagePath, err)
	}
	tracker.UpdateProgress(imagePath, 30) // After read

	originalSize := int64(len(buffer))

	image := bimg.NewImage(buffer)
	processed, err := image.Process(bimg.Options{
		Quality: 80, // Actual compression
	})
	if err != nil {
		return fmt.Errorf("failed to process image %s: %w", imagePath, err)
	}
	tracker.UpdateProgress(imagePath, 70) // After processing

	err = bimg.Write(imagePath, processed) // Overwrite original
	if err != nil {
		return fmt.Errorf("failed to write optimized image %s: %w", imagePath, err)
	}

	optimizedSize := int64(len(processed))
	tracker.UpdateProgress(imagePath, 100)
	tracker.CompleteFile(filepath.Base(imagePath), originalSize, optimizedSize)

	return nil
}

// displayTotalSavings calculates and prints the total space saved.
func displayTotalSavings(images []string, tracker *utils.ProgressTracker) {
	var totalOriginalSize, totalOptimizedSize int64

	for _, img := range images {
		data, err := os.ReadFile(img)
		if err != nil {
			continue
		}
		optimizedSize := int64(len(data))
		originalSize := tracker.OriginalSize(filepath.Base(img))

		if originalSize > 0 {
			totalOriginalSize += originalSize
			totalOptimizedSize += optimizedSize
		}
	}

	if totalOriginalSize == 0 {
		fmt.Println("Could not compute total savings.")
		return
	}

	totalReduction := float64(100) - (float64(totalOptimizedSize)/float64(totalOriginalSize))*100
	fmt.Printf("\n✨ Total Saved: %.1f%% (%d KB) 🎉\n", totalReduction, (totalOriginalSize-totalOptimizedSize)/1024)
}
