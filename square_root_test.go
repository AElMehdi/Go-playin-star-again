package main

import "testing"

func TestSquareRootOfValueOne(t *testing.T) {
	result := Sqrt(1)
	if result != 1 {
		t.Errorf("Expecting %f to be %d", result, 1)
	}
}

func TestSquareRootValueTwo(t *testing.T) {
	result := Sqrt(2)
	if result != 1 {
		t.Errorf("Expecting %f to be %d", result, 1)
	}
}
