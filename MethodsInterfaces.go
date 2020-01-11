package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	methods()
	pointerReceivers()
	interfaces()
	noImplicitImplementation()
	emptyInterface()
	typeAssertions()
	typeSwitches()
	stringers()
	ipAddressStringer()
	errors()
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

	// Nil interface without type
	var iAsNil I
	describe(iAsNil)
	// This call will result on a runtime error
	// No type inside the interface tuple to indicate which concrete method to call
	// iAsNil.M()
}

func describe(i I) {
	fmt.Printf("(%v , %T)\n", i, i)
}

func emptyInterface() {
	var i interface{}
	describeEmpty(i)

	i = 42
	describeEmpty(i)

	i = "A feeeeeeen"
	describeEmpty(i)
}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v , %T)\n", i, i)
}

func typeAssertions() {
	var i interface{} = "Hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // panic
	fmt.Println(f)
}

func typeSwitches() {
	do(21)
	do("Wa feen")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}

// The User struct implements Stringer interface
func (u User) String() string {
	return fmt.Sprintf("%v (%v years)", u.name, u.age)
}

func stringers() {
	marouane := User{"Marouane", 21,}
	youssef := User{"Youssef", 22,}

	fmt.Println(marouane, youssef)
}

// Exercise on Stringers
type IpAddr [4]byte

func (ip IpAddr) String() string {
	return fmt.Sprintf("[%v.%v.%v.%v]", ip[0], ip[1], ip[2], ip[3])
}

func ipAddressStringer() {
	hosts := map[string]IpAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type MyError struct {
	when time.Time
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}

func run() error {
	return &MyError{time.Now(),
		"it didn't work",}
}

func errors() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
