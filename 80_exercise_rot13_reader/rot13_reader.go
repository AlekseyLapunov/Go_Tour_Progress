package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r13.r.Read(b)
	
	for i := range b {
		b[i] = asciiSafeLatinSubstract(b[i], 13)
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func isAsciiCapital(letter byte) bool {
	return letter >= 65 && letter <= 90
}

func isAsciiLowercase(letter byte) bool {
	return letter >= 97 && letter <= 122
}

func asciiSafeLatinSubstract(letter byte, x int) byte {
	if isAsciiCapital(letter) {
		if (int(letter) - x) < 'A' {
			return byte( 'Z' - 'A' + int(letter) - x + 1 ) 
		}

		return byte(int(letter) - x) 
	}

	if isAsciiLowercase(letter) {
		if (int(letter) - x) < 'a' {
			return byte( 'z' - 'a' + int(letter) - x + 1 )
		}

		return byte(int(letter) - x) 
	}

	return letter
}