package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	toBe          = false
	maxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

func main() {
	basicTypes()
	defaultValues()
	typeConversions()
}

func basicTypes() {
	fmt.Printf("Type %T, value %v\n", toBe, toBe)
	fmt.Printf("Type %T, value %v\n", maxInt, maxInt)
	fmt.Printf("Type %T, value %v\n", z, z)
}

func defaultValues() {
	var i int
	var b bool
	var f float64
	var s string
	fmt.Printf("%v %v %v %v \n", i, b, f, s)
}

func typeConversions() {
	var x, y int = 4, 5

	var f float64 = math.Sqrt(float64(x*x + y*y))
	var u uint = uint(f)

	fmt.Println(x, y, u)
}
