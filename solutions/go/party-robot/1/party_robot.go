package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age )
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	greetings := fmt.Sprintf("Welcome to my party, %s!\n", name)
    
	// Using %03d for the table to get that "027" look
	// Using %.1f to round the distance to one decimal place (23.8)
	direc := fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\n", table, direction, distance)
    
	neig := fmt.Sprintf("You will be sitting next to %s.", neighbor)
    
	return greetings + direc + neig
}

