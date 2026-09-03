package cars

import "fmt"
// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	fmt.Println(float64(productionRate) * (successRate/100))
    return float64(productionRate) * (successRate/100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	fmt.Println(int(CalculateWorkingCarsPerHour(productionRate, successRate)/60))
    return int(CalculateWorkingCarsPerHour(productionRate, successRate)/60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	carsInd := carsCount % 10
    carsGrp := (carsCount - carsInd)/10
    return uint(carsInd * 10000 + carsGrp * 95000)
}
