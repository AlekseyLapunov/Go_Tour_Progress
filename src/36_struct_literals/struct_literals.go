package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}	// type: Vertex
	v2 = Vertex{X: 1}	// X:1 Y:0, Y:0 is implicit
	v3 = Vertex{}		// X:0 and Y:0 are implicit
	p = &Vertex{1, 2}	// type: *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}