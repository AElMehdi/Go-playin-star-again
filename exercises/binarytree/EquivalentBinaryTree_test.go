package main

import (
	"golang.org/x/tour/tree"
	"testing"
)

func TestSameReturnsTrueOnEquivalentTrees(t *testing.T) {
	firstTree := tree.New(1)
	secondTree := tree.New(1)

	isSame := Same(firstTree, secondTree)

	if !isSame {
		t.Error("Must be true but was false")
	}
}