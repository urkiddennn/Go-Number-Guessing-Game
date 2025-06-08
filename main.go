package main

import (
	"math/rand"
	"strconv"

	"github.com/pterm/pterm"
)

func main() {
	pterm.Info.Println("Input your Number")

	gNumber := randomNumber()
	inpNumber := 5

	for {
		input, _ := pterm.DefaultInteractiveTextInput.Show("Enter you Guess Number")
		inpNumber -= 1
		newInput, _ := strconv.Atoi(input)
		if inpNumber != 0 {
			trackInput(newInput, gNumber)
			pterm.Info.Println("You have ", inpNumber, " guesses left")
		}

	}
}

// track the user input
func trackInput(input int, gNumber int) {
	if input > gNumber {
		pterm.Warning.Println("The number is less than the Input")
	} else if input < gNumber {
		pterm.Warning.Println("The number is greater than the input")
	} else if input == gNumber {
		pterm.Success.Println("You guess it right")
	}
}

// select random number
func randomNumber() int {
	ranNumber := rand.Intn(100)

	return ranNumber
}
