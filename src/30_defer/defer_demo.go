package main

import "fmt"

func main() {
	fmt.Println("рыба")

	defer fmt.Println("world")

	fmt.Println("hello")
}