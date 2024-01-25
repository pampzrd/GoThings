package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	//productionRate é o total de carros produzido e o sucessRate é a porcentagem de sucesso
	//O resultado deve ser a porcentagem(float64) de carros bem-feitos.
	return float64(productionRate) * successRate / 100
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	PerHour := float64(productionRate) * float64(successRate) / 100
	PerMinute := PerHour / 60
	return int(PerMinute)
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	//IndividualCosts := 10.000
	// 10 = 95.000
	// 37 = 3*10+7
	//Separa os números unidade e dezena e cada um terá um calculo correspondente
	//(3*95.000)+(7*10.000) = 355.000
	deci := carsCount / 10
	uni := carsCount % 10
	result := uint((deci * 95000) + (uni * 10000))
	return result
	panic("CalculateCost not implemented")

}
