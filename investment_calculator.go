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

	// fmt.Println("Future Value", futureValue)
	// Print text in a formatted way
	// values go in order for placeholders
	// \n is a line break fyi

	/**
	Printing also lets you use some other values as well
	AND control the values of the outputs as well.

	Check out the documentation about how GO formats values via fmt package:
	https://pkg.go.dev/fmt
	*/
	fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for Inflation): %.2f", futureValue, futureRealValue)

}
