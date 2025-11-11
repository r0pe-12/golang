package main

import (
	"fmt"
	. "math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter the amount you would like to calculate your investment: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the years you would like to calculate your investment: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
