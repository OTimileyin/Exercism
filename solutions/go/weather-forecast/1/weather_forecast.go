// Package weather is a program that can forecast 
// the current weather condition of various cities in Goblinocus. 
package weather

var (
    // CurrentCondition states the current condition of the weather.
	CurrentCondition string 
    // CurrentLocation  states the current location where the condition exist.
	CurrentLocation  string 
)

// Forecast() displays the CurrentLocation and the CurrentCondtion  of the weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
