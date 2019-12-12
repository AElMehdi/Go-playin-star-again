package main

import (
	"fmt"
	"math"
)

//func main() {
//	forLoop()
//	forLoopContinued()
//	whileLoop()
//	//forever()
//	fmt.Println(sqrt(2), sqrt(-4))
//	fmt.Println(
//		pow(3, 2, 10),
//		pow(3, 3, 20),
//		)
//}

func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum)
}

func forLoopContinued() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	println(sum)
}

func whileLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	println(sum)
}

//func forever() {
//	for {
//	}
//}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, limit)
	}
	return limit
}
