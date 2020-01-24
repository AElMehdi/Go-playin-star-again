package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walks the tree t and send the values in the channel
func Walk(t *tree.Tree, ch chan int)  {
	if t.Left != nil {
		ch<- t.Value
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		ch<- t.Value
		Walk(t.Right, ch)
	}
}

// Check if the values sent are the same
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	// create two trees
	aTree := tree.New(1)
	//anotherTree := tree.New(1)

	// Create a channel and send the tree through it
	aChannel := make(chan int)
	go Walk(aTree, aChannel)
	for v := range aChannel{
		fmt.Println(v)
	}


}
