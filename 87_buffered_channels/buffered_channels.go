package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println(<- ch)
	fmt.Println(<- ch)

	// testing overfill
	ch <- 3
	ch <- 4 // deadlock starts here causing fatal error
	ch <- 5
	fmt.Println(<- ch)
	fmt.Println(<- ch)
	fmt.Println(<- ch)
}
