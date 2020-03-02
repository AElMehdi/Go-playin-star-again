// Examples of OOP techniques without inheritance
package main

import "fmt"

func main() {
	methodsExample()
	simpleCodeCodeReuse()
	delegation()
	polymorphism()
}

// Methods in Go
// Go doesn't have classes but you can have methods on any type declared by type definition by using the func receiver
type myInt int

func (m myInt) IsPositive() bool {
	return m > 0
}

func methodsExample() {
	m := myInt(12)
	fmt.Println(m.IsPositive())

	m = myInt(-78)
	fmt.Println(m.IsPositive())
}

func simpleCodeCodeReuse() {
	fmt.Println("Code reuse example:")
	shape := Shape{color: "black"}
	circle := Circle{shape: shape, otherProps: "someProp"}
	fmt.Println(circle)
	circle.shape.shapeFunc()
}

func delegation() {
	fmt.Println("Delegation example:")
	dog := Dog{}
	dog.Eat()
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

func (a *Animal) Eat() {
	fmt.Println("Animal eats")
}

func (a *Animal) Sleep() {
	fmt.Println("Animal sleeps")
}

func (a *Animal) Jump() {
	fmt.Println("Animal jumps")
}

type Dog struct {
	Animal // Embedding
	//a Animal
}

//func (d *Dog) Eat() {
//	d.a.Eat()
//}
//
//func (d *Dog) Sleep() {
//	d.a.Sleep()
//}
//
//func (d *Dog) Jump() {
//	d.a.Jump()
//}


// Dynamic dispatch/Polymorphism using interface
type Sleeper interface {
	Sleep()
}

type Cat struct {
	Animal
}
func polymorphism()  {
	fmt.Println("Polymorphism example:")
	pets := []Sleeper{new(Dog), new(Cat)}
	for _, pet := range pets {
		pet.Sleep()
	}

}