package purchase
import (
    "fmt"
)
// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    if kind == "car"|| kind == "truck"{
        return true
    } 
    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    var vehicle string
	if option1 < option2 {
        vehicle = option1
    } else {
        vehicle = option2
    }
    return fmt.Sprintf("%v is clearly the better choice.", vehicle)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
   if age < 0 {
		return 0
	}
    var cost float64 
    
    if age < 3 {
    	cost = (originalPrice * (80.0/100.0))
    } else if age >= 10 {
    	cost = (originalPrice * (50.0/100.0))
    } else if age >=3 && age < 10  {
    	cost = (originalPrice * (70.0/100.0))
    }

	return cost

	// switch {
	// case age < 3:
	// 	percentage = 0.80
	// case age >= 10:
	// 	percentage = 0.50
	// default:
	// 	// Covers the 3 to 9.99... range
	// 	percentage = 0.70
	// }

	// return originalPrice * percentage
}
