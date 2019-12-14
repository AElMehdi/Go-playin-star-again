package main

import (
	"fmt"
	"math"
)

var result = float64(1)

func Sqrt(x float64) float64 {
	for i := 0; i < 10; i++ {
		adjustResult(square(result), x)
		fmt.Printf("Current computed square %f\n", result)
	}
	return result
}

func adjustResult(squareResult float64, x float64) {
	result -= (squareResult - x) / (2 * result)
}

func square(number float64) float64 {
	return math.Pow(number, 2)
}
