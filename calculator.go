package main

import (
	"fmt"
	"math"
)

// Calculator provides methods for mathematical calculations
type Calculator struct{}

// CalculateSquare calculates the square of a number
// Returns an error if the result would overflow int
func (c *Calculator) CalculateSquare(number int) (int, error) {
	if number > math.MaxInt32/2 || number < math.MinInt32/2 {
		return 0, fmt.Errorf("square calculation would overflow: %d", number)
	}
	result := number * number
	return result, nil
}

// CalculateCube calculates the cube of a number
// Returns an error if the result would overflow int
func (c *Calculator) CalculateCube(number int) (int, error) {
	maxCube := int(math.Cbrt(float64(math.MaxInt32)))
	if number > maxCube || number < -maxCube {
		return 0, fmt.Errorf("cube calculation would overflow: %d", number)
	}
	result := number * number * number
	return result, nil
}

// CalculatePower calculates number raised to the given power
// Returns an error if the result would overflow int or if power is negative
func (c *Calculator) CalculatePower(base, power int) (int, error) {
	if power < 0 {
		return 0, fmt.Errorf("negative power not supported: %d", power)
	}
	if power == 0 {
		return 1, nil
	}
	if power == 1 {
		return base, nil
	}
	if base == 0 && power > 0 {
		return 0, nil
	}

	result := 1
	absBase := base
	if absBase < 0 {
		absBase = -absBase
	}

	for i := 0; i < power; i++ {
		if result > math.MaxInt32/absBase {
			return 0, fmt.Errorf("power calculation would overflow: %d^%d", base, power)
		}
		result *= base
	}
	return result, nil
}

// CalculateConcurrently performs square and cube calculations concurrently
func (c *Calculator) CalculateConcurrently(number int) (square int, cube int, err error) {
	chSqr := make(chan struct{ result int; err error })
	chCub := make(chan struct{ result int; err error })

	go func() {
		sqr, err := c.CalculateSquare(number)
		chSqr <- struct{ result int; err error }{sqr, err}
	}()
	go func() {
		cub, err := c.CalculateCube(number)
		chCub <- struct{ result int; err error }{cub, err}
	}()

	sqrResult := <-chSqr
	cubResult := <-chCub

	if sqrResult.err != nil {
		return 0, 0, sqrResult.err
	}
	if cubResult.err != nil {
		return 0, 0, cubResult.err
	}

	return sqrResult.result, cubResult.result, nil
}
