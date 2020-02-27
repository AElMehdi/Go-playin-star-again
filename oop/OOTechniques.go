// Examples of OOP techniques without inheritance
package main

import "fmt"

func main() {
	shape := Shape{color: "black"}
	circle := Circle{shape: shape, otherProps: "someProp"}
	fmt.Println(circle)
}

// Code reuse by composition
type Shape struct {
	color string
}

// Circle needs part of Shape properties
type Circle struct {
	shape      Shape
	otherProps string
}
