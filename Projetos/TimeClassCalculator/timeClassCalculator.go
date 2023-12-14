package main

/////////PROJECT NOT FINISHED YET
import "fmt"

// If add array with qtd of index, it needs to be clear to the function.
func calculationTimeArray(arr [20]float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return total
}

func HourMinutesFormat(result float64) {

}

func main() {
	times := 0
	timesEntry := [20]float64{}
	fmt.Println("How many videos do you need to watch? - Maximum 20.") //How to increment like a list, without settled quantity of indexes?
	fmt.Scanln(&times)                                                 //pass

	fmt.Println("Please, input the time of classes and press 'Enter' ") //pass
	for i := 0; i < times; i++ {
		fmt.Scan(&timesEntry[i])
	}

	result := calculationTimeArray(timesEntry)
	// To Do -> How to format like hours and minutes?
	fmt.Println("Time of Classes: ", result)
}
