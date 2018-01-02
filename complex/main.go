package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	fmt.Println(cbrt(2.0))
	fmt.Println(cmplx.Pow(2.0, 1.0/3.0))
}

func cbrt(x complex128) complex128 {
	var z, z0 complex128 = x, x
	for {
		z = z - ((z*z*z - x) / (3.0 * z * z))
		if cmplx.Abs(z-z0) < 1e-10 {
			break
		}
		z0 = z
	}
	return z
}
