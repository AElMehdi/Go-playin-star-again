package main

func main() {
	forLoop()
	forLoopContinued()
	whileLoop()
}

func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum)
}

func forLoopContinued() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	println(sum)
}

func whileLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	println(sum)
}