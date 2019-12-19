package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // Has type vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // Both fields are set to 0
	p  = &Vertex{1, 2} // Has type *Vertex
)

func main() {
	pointers()
	accessingFields()
	accessingFieldsFromPointers()
	structLiterals()
}

func accessingFields() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func accessingFieldsFromPointers() {
	v := Vertex{1, 2}
	p := &v
	p.X = 8
	fmt.Println(v.X)
}

func structLiterals() {
	fmt.Println(v1, p, v2, v3)
}

func pointers() {
	i, j := 42, 72

	p := &i         // Point to i
	fmt.Println(*p) // Read i through the pointer
	*p = 21         // Set i through the pointer
	fmt.Println(i)  // See the new value of i

	p = &j         // Point to j
	*p = *p / 37   // Divide j through the pointer
	fmt.Println(j) // See the new value of j
}
