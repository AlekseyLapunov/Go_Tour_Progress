package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	comically_small_buffer := make([]byte, 8)

	for {
		n, err := r.Read(comically_small_buffer)

		fmt.Printf("n = %v err = %v b = %v\n", n, err, comically_small_buffer)
		fmt.Printf("b[:n] = %q\n", comically_small_buffer[:n])

		if err == io.EOF {
			break
		}
	}
}