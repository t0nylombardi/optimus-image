package optimizer

import "github.com/t0nylombardi/optimus-image/optimus-image/internal/progress"

// FileOptimizer defines the contract for optimizing images.
type FileOptimizer interface {
	OptimizeFiles(images []string, tracker progress.ProgressTracker) error
}
