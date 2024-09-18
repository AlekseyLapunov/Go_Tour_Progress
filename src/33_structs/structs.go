package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type myRune uint32

func main() {
	fmt.Println(Vertex{1, 2})

	var a myRune = 1

	fmt.Println(a)
}