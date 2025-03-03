package main

import (
	"fmt"
	"os"

	"github.com/t0nylombardi/optimus-image/src/optimus-image/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
