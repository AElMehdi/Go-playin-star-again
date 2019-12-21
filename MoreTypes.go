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
	arrays()
	slices()
	slicesAndSHaredData()
	slicesAndLiterals()
	slicesDefaultBounds()
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

func arrays() {
	var greeting [2]string
	greeting[0] = "khouna"
	greeting[1] = "fl JouGhouna"

	fmt.Println(greeting[0], greeting[1])
	fmt.Println(greeting)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]

	fmt.Println(s)
}

func slicesAndSHaredData() {
	names := [4]string{
		"Nasser",
		"Kamal",
		"Redouane",
		"Youness",
	}

	var a = names[0:2]
	var b = names[1:3]

	b[0] = "XXXX"

	fmt.Println(a, b)
	fmt.Println(names)
}

func slicesAndLiterals() {
	booleans := []bool{true, false, false,}
	fmt.Println(booleans)

	integers := []int{1, 5, 6, 8, 9,}
	fmt.Println(integers)

	structTypes := [] struct {
		i int
		b bool
	}{
		{1, true},
		{0, false},
	}

	fmt.Println(structTypes)
}

func slicesDefaultBounds() {
	numbers := []int{3, 5, 8, 7, 9,}
	fmt.Println(numbers)

	numbers = numbers[2:5]
	fmt.Println(numbers) // 8 7 9

	numbers = numbers[:2]
	fmt.Println(numbers) // 8 7

	numbers = numbers[1:]
	fmt.Println(numbers) // 7
}
