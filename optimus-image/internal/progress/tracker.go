package progress

// ProgressTracker defines the interface for tracking optimization progress.
type ProgressTracker interface {
	StartSpinner()
	StopSpinner()
	UpdateProgress(file string, percent int)
	CompleteFile(file string, originalSize, optimizedSize int64)
}
