package utils

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

// ProgressTracker manages multiple stacked progress bars.
type ProgressTracker struct {
	mu            sync.Mutex
	totalFiles    int
	currentFile   int
	spinner       *spinner.Spinner
	progress      map[string]int
	results       map[string]string
	originalSizes map[string]int64
}

// NewTracker initializes a progress tracker.
func NewTracker(total int) *ProgressTracker {
	return &ProgressTracker{
		totalFiles:    total,
		spinner:       spinner.New(spinner.CharSets[9], 100*time.Millisecond),
		progress:      make(map[string]int),
		results:       make(map[string]string),
		originalSizes: make(map[string]int64),
	}
}

// StartSpinner begins the overall progress spinner.
func (p *ProgressTracker) StartSpinner() {
	p.spinner.Suffix = fmt.Sprintf(" Just a moment... (0/%d)", p.totalFiles)
	p.spinner.Start()
}

// UpdateProgress updates the progress for a specific image.
func (p *ProgressTracker) UpdateProgress(file string, percent int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.progress[file] = percent
	p.render()
}

// CompleteFile marks an image as optimized and updates the results.
func (p *ProgressTracker) CompleteFile(file string, originalSize, optimizedSize int64) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.currentFile++
	p.spinner.Suffix = fmt.Sprintf(" Just a moment... (%d/%d)", p.currentFile, p.totalFiles)

	reduction := 100 - (float64(optimizedSize)/float64(originalSize))*100
	var status string

	if reduction > 30 {
		status = color.GreenString("🟩 [%-10s] %dKB → %dKB (-%.1f%%) ✅", strings.Repeat("█", int(reduction/10)), originalSize, optimizedSize, reduction)
	} else if reduction > 10 {
		status = color.YellowString("🟨 [%-10s] %dKB → %dKB (-%.1f%%) ⚠️", strings.Repeat("█", int(reduction/10)), originalSize, optimizedSize, reduction)
	} else {
		status = color.RedString("🟥 [%-10s] %dKB → %dKB (-%.1f%%) ❌", strings.Repeat("█", int(reduction/10)), originalSize, optimizedSize, reduction)
	}

	p.originalSizes[file] = originalSize
	p.results[file] = fmt.Sprintf("📷 %s\n%s", file, status)
	p.render()
}

// render updates the terminal output dynamically.
func (p *ProgressTracker) render() {
	fmt.Print("\033[H\033[J") // Clear terminal

	// Display spinner with current file count
	fmt.Println(p.spinner.Suffix)

	// Show each file's progress
	for file, percent := range p.progress {
		fmt.Printf("📷 %s\n[%s%s] %d%%\n", file, strings.Repeat("█", percent/10), strings.Repeat(".", 10-percent/10), percent)
	}

	// Show completed files with color-coded results
	for _, result := range p.results {
		fmt.Println(result)
	}
}

// StopSpinner stops the spinner.
func (p *ProgressTracker) StopSpinner() {
	p.spinner.Stop()
}

func (p *ProgressTracker) OriginalSize(file string) int64 {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.originalSizes[file]
}
