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
	typeInference()
	constants()
	numericConstants()
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

func typeInference() {
	v := 42.4i // change it
	fmt.Printf("v is of type %T \n", v)
}

const Pi = 3.14

func constants() {
	const World = "عالم"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

const (
	Big   = 1 << 30
	Small = Big >> 99
)

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func numericConstants() {
	fmt.Println(needInt(Small))
	fmt.Println(needInt(Big))

	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}