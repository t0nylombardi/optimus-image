package main

import (
	"fmt"
	"os"

	"github.com/t0nylombardi/optimus-image/src/optimus-image/cmd"
)

func main() {
	selection, err := cmd.Execute(cmd.GetUserSelection)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("User selected:", selection)
}
