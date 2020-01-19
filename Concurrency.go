package main

import (
	"fmt"
	"time"
)

func main() {
	goroutines()
	channels()
}

func goroutines() {
	go say("world")
	say("hello")
}

func say(word string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(word)
	}
}

func channels() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _,v := range s {
		sum += v
	}
	c <- sum
}
