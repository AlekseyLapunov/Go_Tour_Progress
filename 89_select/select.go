package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		select { // blocking until one of the cases can run and executes it
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)

	quit := make(chan int)

	go func(lim int) {
		for i := 0; i < lim; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}(10)

	fibonacci(c, quit)
}