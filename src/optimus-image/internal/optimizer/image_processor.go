package optimizer

import (
	"path/filepath"
	"time"

	"github.com/t0nylombardi/optimus-image/src/optimus-image/internal/progress"
)

// ImageProcessor defines the interface for image processing.
type ImageProcessor interface {
	Process(imagePath string, tracker progress.ProgressTracker) error
}

// DefaultImageProcessor is the default implementation of ImageProcessor.
type DefaultImageProcessor struct{}

// Process simulates image optimization.
func (p *DefaultImageProcessor) Process(imagePath string, tracker progress.ProgressTracker) error {
	for i := 0; i <= 100; i += 10 {
		time.Sleep(50 * time.Millisecond)
		tracker.UpdateProgress(imagePath, i)
	}

	originalSize, optimizedSize := calculateOptimization(imagePath)
	tracker.CompleteFile(filepath.Base(imagePath), originalSize, optimizedSize)
	return nil
}

// calculateOptimization simulates image size reduction.
func calculateOptimization(imagePath string) (int64, int64) {
	originalSize := int64(1024 * (1 + len(imagePath)%5)) // Mock file size in KB
	reductionPercent := 30 + (len(imagePath) % 50)       // Mock reduction percentage
	optimizedSize := originalSize * (100 - int64(reductionPercent)) / 100
	return originalSize, optimizedSize
}
