package main

import (
	"fmt"
)

func main()  {
	fmt.Println("The sum is", add(2, 3))
}

func add(x int, y int) int {
	return x + y
}