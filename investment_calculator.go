package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5 // never change this value
	var investmentAmount,
		years float64 // you can declare a varible with no value - go automatically allows (doesn't work with constants)
	expectedReturnRate := 5.5

	// Show a user a message
	fmt.Print("Investment Amount: ")
	// Gets inputs from the terminal
	// need to add a 'pointer' in front of the variable
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years of Investment: ")
	fmt.Scan(&years)

	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
