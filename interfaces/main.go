package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs()
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	a.Abs()
	f.Abs()
}

type MyFloat float64

func (f MyFloat) Abs() {
	fmt.Println("Abs() Method of MyFloat")
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() {
	fmt.Println("Abs() method of Vertex Struct")
}
