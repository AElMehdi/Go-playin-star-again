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
	firstChan, secondChan := make(chan int), make(chan int)

	go Walk(t1, firstChan)
	go Walk(t2, secondChan)

	for {
		x1, ok1 := <-firstChan
		x2, ok2 := <-secondChan

		if areChannelsEmpty(ok1, ok2) {
			return haveSameSize(ok1, ok2)
		}
		if x1 != x2 {
			return false
		}
	}
}

func haveSameSize(ok1 bool, ok2 bool) bool {
	return ok1 == ok2
}

func areChannelsEmpty(ok1 bool, ok2 bool) bool {
	return !ok1 || !ok2
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(10), tree.New(1)))
}

// Steps to implement the logic:
/*
1- Walk a single tree DONE
2- Print the values from the channel passed to the Walk function DONE
3- Implement the Same function which will just compare two tree, through the values passed in the channels! DONE
*/
