package main

import (
	"fmt"
	. "math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	outputText("Enter the amount you would like to calculate your investment: ")
	//fmt.Print("Enter the amount you would like to calculate your investment: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the years you would like to calculate your investment: ")
	fmt.Scan(&years)

	//futureValue := investmentAmount * Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calcFutureValues(investmentAmount, expectedReturnRate, years)

	//futureValueF := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	//futureRealValueF := fmt.Sprintf("Future Real Value: %.2f\n", futureRealValue)
	//fmt.Println("Future Value: ", futureValue)
	//fmt.Println("Future Real Value: ", futureRealValue)
	fmt.Printf("Future Value: %.2f\nFuture Real Value: %.2f\n", futureValue, futureRealValue)
	//fmt.Print(futureValueF, futureRealValueF)

	//	fmt.Printf(`Future Value: %.2f
	//Future Real Value: %.2f`, futureValue, futureRealValue)
}

func outputText(output string) {
	fmt.Print(output)
}

func calcFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * Pow(1+expectedReturnRate/100, years)
	rfv := fv / Pow(1+inflationRate/100, years)
	return fv, rfv
}
