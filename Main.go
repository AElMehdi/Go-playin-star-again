package main

import (
	"fmt"
)

func main()  {
	create("hello", false, true, 1, 2,3)
	fmt.Println("The sum is", add(2, 3))
	fmt.Println(swap("first", "last"))
}

func add(x, y int) int {
	return x + y
}
func create(s string, b, b1 bool, i, i2, i3 int ) {
	fmt.Println(s, b, b1, i, i2, i3)
}

func swap(x string, y string) (string, string) {
	return y, x
}
