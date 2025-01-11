package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate = 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)
	futuredValue, futuredRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// futuredValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futuredRealValue := futuredValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futuredValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futuredRealValue)
	// fmt.Println("Future Value:", futuredValue)
	// fmt.Printf(`Future Value: %.1f\n

	// Future Value (adjusted for Inflation): %.1f\n`, futuredValue, futuredRealValue)
	// fmt.Println("Future Value (adjusted for Inflation):", futuredRealValue)
	fmt.Print(formattedFV, formattedRFV)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
