package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Method is just a function but with the receiver argument (that is commented below)
func /*(v Vertex)*/ Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}