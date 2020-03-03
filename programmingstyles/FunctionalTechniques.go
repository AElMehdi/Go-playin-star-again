package main

import (
	"fmt"
	"sort"
)

func main() {
	closureExample()
	closuresAndScope()
}

// Closures
func closureExample() {
	people := []string{"Alice", "Bob", "Dave"}

	less := func(i, j int) bool {
		return len(people[i]) < len(people[j])
	}

	sort.Slice(people, less)
	fmt.Println(people)
}

// Another example of closures and variables accessibility
func closuresAndScope() {
	New := func() (Count func()){
		n := 0
		return func() {
			n++
			fmt.Println(n)
		}
	}

	f1, f2 := New(), New()
	f1()
	f1()
	f2()
	f1()
}