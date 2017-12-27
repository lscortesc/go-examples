// Bucles & Functions
// Square Root aprox with Newton's method
// Formule -> z = z - ((z²-x)/2z)
package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i < 1000; i++ {
		j := (float64(i)*7.894)*3.7908 + (float64(i) / 12.48)

		fmt.Println("Square root of", j)
		fmt.Println("Math Sqrt Golang", math.Sqrt(j))
		fmt.Println("Newton-Raphson Improved", newtonImproved(j))
		fmt.Println("Newton - Raphson", newton(j))
		fmt.Println()
	}
}

func newton(x float64) float64 {

	z := float64(1)
	for i := 0; i < 10; i++ {
		fa := (z * z) - x
		faprima := 2 * z

		z = z - (fa / faprima)
	}

	return z
}

func newtonImproved(x float64) float64 {

	z := float64(1)
	delta := float64(0.00000000001)

	for absolute(z*z-x) >= delta {
		fa := (z * z) - x
		faprima := 2 * z

		z = z - (fa / faprima)
	}

	return z
}

func absolute(x float64) float64 {
	if x < 0 {
		return -x
	}

	return x
}
