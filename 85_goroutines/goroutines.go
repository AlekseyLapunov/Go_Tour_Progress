package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	defer say("deferred") // just for test

	go say("world")
	say("hello")
}