package techpalace

import (
    "fmt"
    "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    name := strings.ToUpper(customer)
	message:= fmt.Sprintf("Welcome to the Tech Palace, %s", name)
    return message
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var result string = ""
	for i := 0; i < numStarsPerLine; i++ {
		result += "*"
	}
	result += "\n" + welcomeMsg + "\n"

	for i := 0; i < numStarsPerLine; i++ {
		result += "*"
	}
	return result
 //    border := strings.Repeat("*", numStarsPerLine)
	// return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	
	nostar := strings.ReplaceAll(oldMsg, "*", "")
	return strings.TrimSpace(nostar)
}
