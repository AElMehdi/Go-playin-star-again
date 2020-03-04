package main

import (
	"fmt"
	"sort"
)

func main() {
	closureExample()
	closuresAndScope()
	higherOrderFunction()
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
	New := func() (Count func()) {
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

// Higher order functions example
func higherOrderFunction() {
	fruits := []string{"Orange", "Apple", "Strawberry", "Pineapple"}
	transformed := mapForEach(fruits, func(it string) int {
		return len(it)
	})

	fmt.Println(transformed)

}

func mapForEach(list []string, function func(it string) int) []int {
	var result []int
	for _, word := range list {
		result = append(result, function(word))
	}
	return result
}
