package main

import (
	"fmt"
	"math/cmplx"
)

var (
	toBe bool = false
	maxInt uint64 = 1 << 64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var i int
	var b bool
	var f float64
	var s string
	//fmt.Printf("Type %T, value %v\n", toBe, toBe)
	//fmt.Printf("Type %T, value %v\n", maxInt, maxInt)
	//fmt.Printf("Type %T, value %v\n", z, z)
	fmt.Printf("%v %v %v %v", i, b, f, s)

}