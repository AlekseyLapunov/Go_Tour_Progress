package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
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

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}