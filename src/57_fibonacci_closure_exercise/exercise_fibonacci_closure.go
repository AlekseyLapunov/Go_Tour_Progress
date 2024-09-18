package main

import "fmt"

func fibonacci() func() int {
	current, previous := -1, 0

	fmt.Printf("Entered fibonacci() with current=%d and previous=%d\n", current, previous)
	
	return func() int {
		if current < 1 {
			current++
			return current
		}

		previous, current = current, current + previous
		return current
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}