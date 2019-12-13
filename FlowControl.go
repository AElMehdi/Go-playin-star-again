package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	//forLoop()
	//forLoopContinued()
	//whileLoop()
	////forever()
	//fmt.Println(sqrt(2), sqrt(-4))
	//fmt.Println(
	//	pow(3, 2, 10),
	//	pow(3, 3, 20),
	//	)
	switchCase()
}

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

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//func forever() {
//	for {
//	}
//}

func pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, limit)
	}
	return limit
}

func switchCase() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}
}
