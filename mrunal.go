package main

import (
	"fmt"
)

// Calculator provides methods for mathematical calculations
type Calculator struct{}

// CalculateSquare calculates the square of a number
func (c *Calculator) CalculateSquare(number int) int {
	return number * number
}

// CalculateCube calculates the cube of a number
func (c *Calculator) CalculateCube(number int) int {
	return number * number * number
}

// CalculateConcurrently performs square and cube calculations concurrently
func (c *Calculator) CalculateConcurrently(number int) (square, cube int) {
	chSqr := make(chan int)
	chCub := make(chan int)

	go func() { chSqr <- c.CalculateSquare(number) }()
	go func() { chCub <- c.CalculateCube(number) }()

	return <-chSqr, <-chCub
}

func main() {
	calc := &Calculator{}
	number := 2
	square, cube := calc.CalculateConcurrently(number)
	fmt.Printf("Number: %d, Square: %d, Cube: %d\n", number, square, cube)
}
