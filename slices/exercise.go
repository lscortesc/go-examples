package main

import (
	"golang.org/x/tour/pic"
)

func Picture(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for x := range pic {
		pic[x] = make([]uint8, dx)

		for y := range pic[x] {
			pic[x][y] = uint8(x * y)
		}

	}

	return pic
}

func slicesExercise() {
	pic.Show(Picture)

}
