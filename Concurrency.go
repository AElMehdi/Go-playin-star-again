package main

import (
	"fmt"
	"time"
)

func main() {
	goroutines()
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
