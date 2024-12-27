package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futuredValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futuredRealValue := futuredValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futuredValue)
	fmt.Println(futuredRealValue)

}
