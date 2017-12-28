// Structs & Pointers
package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

var (
	r = Vertex{1, 2}  // It has Vertex type
	s = &Vertex{1, 2} // It has *Vertex type
	t = Vertex{X: 1}  // Y:0 is implicit - Zero value
	u = Vertex{}      // X:0 and Y:0 are implicit - Zero value
)

func main() {
	firstVertex()
	secondVertex()
	fmt.Println(r, s, t, u)
	newVertex()
}

func firstVertex() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func secondVertex() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e3
	fmt.Println(p)
}

func newVertex() {
	// Stores at memory a 'T' value with zero value and
	// return a pointer
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 10, 13
	fmt.Println(v)
}
