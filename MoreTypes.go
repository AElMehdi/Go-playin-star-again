package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	pointers()
	accessingFields()
	accessingFieldsFromPointers()
}

func accessingFields() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}


func accessingFieldsFromPointers() {
	v := Vertex{1, 2}
	p :=  &v
	p.X = 8
	fmt.Println(v.X)
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
