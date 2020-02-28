// Examples of OOP techniques without inheritance
package main

import "fmt"

func main() {
	simpleCodeCodeReuse()
	delegation()
}

func simpleCodeCodeReuse() {
	shape := Shape{color: "black"}
	circle := Circle{shape: shape, otherProps: "someProp"}
	fmt.Println(circle)
	circle.shape.shapeFunc()
}

func delegation() {
	fmt.Println("Second example")
	dog := Dog{}
	dog.eat()
}

// Code reuse by composition
type Shape struct {
	color string
}

func (s *Shape) shapeFunc() {
	fmt.Println("Shape function called")
}

// Circle needs part of Shape properties
type Circle struct {
	shape      Shape
	otherProps string
}

// Another example with a lot of functions
// Using delegation
type Animal struct {
}

func (a *Animal) eat() {
	fmt.Println("Animal eats")
}

func (a *Animal) sleep() {
	fmt.Println("Animal sleeps")
}

func (a *Animal) jump() {
	fmt.Println("Animal jumps")
}

type Dog struct {
	Animal // Embedding
	//a Animal
}

//func (d *Dog) eat() {
//	d.a.eat()
//}
//
//func (d *Dog) sleep() {
//	d.a.sleep()
//}
//
//func (d *Dog) jump() {
//	d.a.jump()
//}
