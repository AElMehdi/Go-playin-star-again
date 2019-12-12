package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const z = float64(1)

	if squareZ := square(z); squareZ == x {
		fmt.Printf("Current computed square %f\n", z)
		return z
	}
	return z
}

func square(z float64) float64 {
	squareZ := math.Pow(z, 2)
	return squareZ
}