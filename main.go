package main

import (
	"fmt"
)

func main() {
	calc := &Calculator{}
	number := 2
	square, cube := calc.CalculateConcurrently(number)
	fmt.Printf("Number: %d, Square: %d, Cube: %d\n", number, square, cube)
}
