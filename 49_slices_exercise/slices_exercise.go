// This exercise is needs to be tested in the playground

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic_slice := make([][]uint8, dy)
	
	for i := 0; i < len(pic_slice); i++ {
		pic_slice[i] = make([]uint8, dx)
	}
	
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			//pic_slice[uint8(x)][uint8(y)] = uint8(x ^ y)
			//pic_slice[uint8(x)][uint8(y)] = uint8(x * y)
			pic_slice[uint8(x)][uint8(y)] = uint8((x + y)/2)
		}
	}
	
	return pic_slice
}


func main() {
	pic.Show(Pic)
}