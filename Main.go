package main

import (
	"fmt"
)

// Package level declaration
var i, j int = 1, 2

func main() {
	// Function level declaration
	c, python, java := true, false, "Yes!"
	//create("hello", false, true, 1, 2,3)
	//fmt.Println("The sum is", add(2, 3))
	//fmt.Println(swap("first", "last"))
	//fmt.Println(split(17))
	fmt.Println(i, j, c, python, java)
}

// Simple function
func add(x, y int) int {
	return x + y
}

// Params with same type
func create(s string, b, b1 bool, i, i2, i3 int) {
	fmt.Println(s, b, b1, i, i2, i3)
}

// Multiple returns
func swap(x string, y string) (string, string) {
	return y, x
}

// Named return values, naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
