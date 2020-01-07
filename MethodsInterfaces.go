package main

import (
	"fmt"
	"math"
)

func main() {
	methods()
	pointerReceivers()
	interfaces()
	noImplicitImplementation()
}

type MyInt int

type User struct {
	name string
	age  int
}

func (u User) isAdult() bool {
	return u.age >= 18
}

func setAge(u *User, age int) {
	u.age = u.age + age
}

func (age MyInt) isAdult() bool {
	return age >= 18
}

func isAdult(user User) bool {
	return user.age >= 18
}

func methods() {
	anAdult := User{"Dean", 45,}
	fmt.Println(anAdult.isAdult())
	fmt.Println(isAdult(anAdult))

	anAge := MyInt(13)
	fmt.Println(anAge.isAdult())
}

func pointerReceivers() {
	youssef := User{"Youssef", 15}
	fmt.Println(youssef.isAdult())
	setAge(&youssef, 3)
	fmt.Println(youssef.isAdult())
}

// Interfaces
// Abser interface
type Abser interface {
	abs() float64
}

// A type MyVertex: Implements Abser
type MyVertex struct {
	x, y float64
}

func (v *MyVertex) abs() float64 {
	fmt.Println("Vertex called")
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Float wrapper
type MyFloat float64

func (f MyFloat) abs() float64 {
	fmt.Println("MyFloat called")
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) M() {
	fmt.Println(f)
}

func interfaces() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := MyVertex{3, 4,}

	// Make MyFloat implements the interface Abser
	a = f
	// Make *MyVertex implements the interface Abser
	a = &v

	fmt.Println(a.abs())
}

type I interface {
	M()
}

type T struct {
	S string
}

// Implicit implementation of an interface
//func (t T) M() {
//	fmt.Println(t.S)
//}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func noImplicitImplementation() {
	var i I = &T{"Hello"}
	describe(i)
	i.M()

	i = MyFloat(math.Pi)
	describe(i)
	i.M()

	var t *T
	i = t
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v , %T)\n", i, i)
}
