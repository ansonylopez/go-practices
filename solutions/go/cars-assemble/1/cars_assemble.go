package cars

const perCentum float64 = 100.00
const minutesThatHasOneHour int = 60
const costOfTenCars int = 95000
const costEachCar int = 10000


// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    var productionRateFloat = float64(productionRate)
	return ((successRate/perCentum)*productionRateFloat)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    var productionRatePerMinute = float64(productionRate/minutesThatHasOneHour)
    return  int((successRate/perCentum)*productionRatePerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	return uint((carsCount/10*costOfTenCars)+(carsCount%10*costEachCar))
}
