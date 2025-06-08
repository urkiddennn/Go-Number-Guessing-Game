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
		newInput, _ := strconv.Atoi(input)
		if inpNumber != 0 {
			npNumber := trackInput(newInput, gNumber, inpNumber)
			pterm.Info.Println("You have ", npNumber, " guesses left")
		}

	}
}

// track the user input
func trackInput(input int, gNumber int, npNumber int) int {
	if input > gNumber {
		pterm.Warning.Println("The number is less than the Input")
		npNumber -= 1
	} else if input < gNumber {
		pterm.Warning.Println("The number is greater than the input")
		npNumber -= 1
	} else if input == gNumber {
		pterm.Success.Println("You guess it right")
	}

	return npNumber
}

// select random number
func randomNumber() int {
	ranNumber := rand.Intn(100)

	return ranNumber
}
