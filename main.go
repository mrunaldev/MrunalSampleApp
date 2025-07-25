package main

import (
	"fmt"
	"log"
)

func main() {
	calc := &Calculator{}
	number := 2

	// Calculate square and cube concurrently with error handling
	square, cube, err := calc.CalculateConcurrently(number)
	if err != nil {
		log.Fatalf("Error in concurrent calculation: %v", err)
	}
	fmt.Printf("Number: %d, Square: %d, Cube: %d\n", number, square, cube)

	// Demonstrate power calculation
	base, power := 2, 3
	result, err := calc.CalculatePower(base, power)
	if err != nil {
		log.Fatalf("Error calculating power: %v", err)
	}
	fmt.Printf("%d raised to power %d is: %d\n", base, power, result)
}
