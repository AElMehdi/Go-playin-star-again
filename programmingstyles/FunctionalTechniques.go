package main

import (
	"fmt"
	"sort"
)

func main() {
	closureExample()
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
