package main

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Width  int
	Height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i Image) At(x, y int) color.Color {
	imageClosure := func(x, y int) uint8 {
		return uint8(x ^ y)
	}

	v := imageClosure(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	exercise()
}

func exercise() {
	m := Image{256, 64}
	pic.ShowImage(m)
}
