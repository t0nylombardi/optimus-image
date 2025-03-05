package progress

import (
	"fmt"
	"sync"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

// DefaultProgressTracker implements ProgressTracker.
type DefaultProgressTracker struct {
	mu          sync.Mutex
	totalFiles  int
	currentFile int
	spinner     *spinner.Spinner
	progress    map[string]int
	results     map[string]string
}

// NewTracker creates a new progress tracker.
func NewTracker(total int) ProgressTracker {
	return &DefaultProgressTracker{
		totalFiles: total,
		spinner:    spinner.New(spinner.CharSets[9], 100),
		progress:   make(map[string]int),
		results:    make(map[string]string),
	}
}

func (p *DefaultProgressTracker) StartSpinner() {
	p.spinner.Suffix = fmt.Sprintf(" Just a moment... (0/%d)", p.totalFiles)
	p.spinner.Start()
}

func (p *DefaultProgressTracker) StopSpinner() {
	p.spinner.Stop()
}

func (p *DefaultProgressTracker) UpdateProgress(file string, percent int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.progress[file] = percent
}

func (p *DefaultProgressTracker) CompleteFile(file string, originalSize, optimizedSize int64) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.currentFile++
	p.spinner.Suffix = fmt.Sprintf(" Just a moment... (%d/%d)", p.currentFile, p.totalFiles)
	reduction := 100 - (float64(optimizedSize)/float64(originalSize))*100

	var status string
	if reduction > 30 {
		status = color.GreenString("🟩 %dKB → %dKB (-%.1f%%) ✅", originalSize, optimizedSize, reduction)
	} else {
		status = color.YellowString("🟨 %dKB → %dKB (-%.1f%%) ⚠️", originalSize, optimizedSize, reduction)
	}

	p.results[file] = fmt.Sprintf("📷 %s\n%s", file, status)
}
