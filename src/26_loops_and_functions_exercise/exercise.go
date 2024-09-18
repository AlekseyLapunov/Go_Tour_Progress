package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x < 0 {
		fmt.Println("Negative input, returning 0")
		return 0
	}

	z := 1.0
	memo_odd := 0.0
	memo_even := 0.0

	for i := 0;; i++ {
		if i % 2 == 0 {
			memo_even = z;
		} else {
			memo_odd = z;
		}

		z -= (z*z - x)/(2*z)
		
		if z == memo_odd || z == memo_even {
			break
		}
	}

	return z
}

func main() {
	const Value float64 = 2

	my_value := Sqrt(Value)
	math_value := math.Sqrt(Value)

	fmt.Printf("main.Sqrt(%g): %g\nmath.Sqrt(%g): %g\n",
		Value, my_value,
		Value, math_value)
}