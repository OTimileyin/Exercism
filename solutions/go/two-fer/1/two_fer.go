// Package twofer is a simulation of the two-for-one offer
// made in the market place - buy two articles for the price of one.
package twofer

// ShareWith returns a customer's message from the concept of two for one
func ShareWith(name string) string {
    message := ""
    if name == ""{
		message = "One for you, one for me." 
        return message
    }
    message = "One for " + name + ", one for me."        
	return message
}
