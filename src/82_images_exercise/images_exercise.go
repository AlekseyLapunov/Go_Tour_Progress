// This exercise needs to be tested at go tour playground

package main

import "golang.org/x/tour/pic"
import ("image"
	    "image/color")

type Image struct {
	W int
	H int
	C int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.W, i.H)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(i.C + x), uint8(i.C + y), 255, 255}
}

func main() {
	m := Image{350, 150, 250}
	pic.ShowImage(m)
}
