package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	for i := 0; i < 10; i++ {
		adjustmentValue := adjustmentValue(z, x)

		if adjustmentValue == 0 {
			break
		}

		z -= adjustmentValue
		fmt.Printf("Current computed square %f\n", z)
	}

	return z
}

func adjustmentValue(z float64, x float64) float64 {
	return (square(z) - x) / (2 * z)
}

func square(number float64) float64 {
	return math.Pow(number, 2)
}
