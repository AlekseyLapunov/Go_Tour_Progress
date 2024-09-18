package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	time.Sleep(time.Second) // waiting demonstration

	c <- sum // send
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // recv and assign; 
                     // so here is the "waiting" for the values
                     // to be available for fetching from the channel;
                     // therefore channels provide goroutines synchronization

	fmt.Println(x, y, x+y)
}