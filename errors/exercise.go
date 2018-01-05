package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(1)
	for i := 0; i < 10; i++ {
		fa := (z * z) - x
		faprima := 2 * z

		z = z - (fa / faprima)
	}

	return z, nil
}

func exercise() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
