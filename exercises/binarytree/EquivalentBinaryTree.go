package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walks the tree t and send the values in the channel
func Walk(t *tree.Tree, ch chan int) {
	recursiveWalker(t, ch)
	close(ch)
}

func recursiveWalker(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	recursiveWalker(t.Left, ch)
	ch <- t.Value
	recursiveWalker(t.Right, ch)
}

// Check if the values sent are the same
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	aTree := tree.New(1)
	aChannel := make(chan int)

	go Walk(aTree, aChannel)

	for v := range aChannel {
		fmt.Println(v)
	}
}


// Steps to implement the logic:
/*
1- Walk a single tree DONE
2- Print the values from the channel passed to the Walk function DONE
3- Implement the Same function which will just compare two tree, through the values passed in the channels!
*/