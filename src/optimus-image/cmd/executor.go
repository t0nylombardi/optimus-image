package cmd

// Executor defines the contract for CLI execution.
type Executor interface {
	Execute(getSelectionFunc func() (string, error)) (string, error)
}
