package main

import (
	"strconv"

	"github.com/pterm/pterm"
)

type Input struct {
	numberInput int
}

func main() {
	pterm.Info.Println("Input your Number")
	input, _ := pterm.DefaultInteractiveTextInput.Show("Enter you Guess Number")
	newInput, _ := strconv.ParseInt(input, 0, 64)

	trackInput(newInput)
}

// track the user input
func trackInput(input int64) {
}
