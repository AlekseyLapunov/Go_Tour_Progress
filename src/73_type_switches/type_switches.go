package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("[%T] Twice %v is %v\n", v, v, v*2)
	case string:
		fmt.Printf("[%T] %q is %v bytes long\n", v, v, len(v))
	default:
		fmt.Printf("[%T] This type is not supported by this switch\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
