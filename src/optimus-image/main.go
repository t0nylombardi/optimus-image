package main

import (
	"log"

	"github.com/t0nylombardi/optimus-image/src/optimus-image/cmd"
	"github.com/t0nylombardi/optimus-image/src/optimus-image/internal/optimizer"
	"github.com/t0nylombardi/optimus-image/src/optimus-image/internal/utils"
)

func main() {
	// Instantiate the dependencies
	fileUtils := &utils.FileUtilsImpl{}
	fileOptimizer := &optimizer.FileOptimizerImpl{}

	// Create an Executor instance with the dependencies
	executor := &cmd.Executor{
		FileUtils:     fileUtils,
		FileOptimizer: fileOptimizer,
	}

	// Call Execute with GetUserSelection as the input function
	_, err := executor.Execute(cmd.GetUserSelection)
	if err != nil {
		log.Fatal(err)
	}
}
