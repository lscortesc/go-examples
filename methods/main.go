package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Method to Struct Type

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Method to any type in the same package
// You can't define a method to a type from antoher package neither to a basic type

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
