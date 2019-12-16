package main

import (
	"fmt"
)

// Package level declaration
var i, j int = 1, 2

func main() {
	//BasicTypes()
	FlowControl()
}

//func BasicTypes() {
//	// Function level declaration
//	c, python, java := true, false, "Yes!"
//	create("hello", false, true, 1, 2,3)
//	fmt.Println("The sum is", add(2, 3))
//	fmt.Println(swap("first", "last"))
//	fmt.Println(split(17))
//	fmt.Println(i, j, c, python, java)
//	basicTypes()
//	defaultValues()
//	typeConversions()
//	typeInference()
//	constants()
//	numericConstants()
//}

func FlowControl() {
	forLoop()
	forLoopContinued()
	whileLoop()
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	switchCase()
	switchCaseEvaluationOrder()
	switchCaseWithoutCondition()
}
