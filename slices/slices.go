package main

import "fmt"

func antoherSlice() {
	fmt.Println("\n* Another Slice *\n")
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

func slicingSlices() {
	fmt.Println("\n* Slicing Slices *\n")
	p := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("p =", p)
	fmt.Println("p[1:4] =", p[1:4])

	// Init value omitted, it implicates 0
	fmt.Println("p[:3] =", p[:3])

	// End value omitted, it implicates len(s)
	fmt.Println("p[4:] =", p[4:])
}

func createSlices() {
	fmt.Println("\n* Creating Slices *\n")

	// Slices are created with 'make' function
	a := make([]int, 5) // Create a slice with length = 5
	printSlice("a", a)

	b := make([]int, 0, 5) // Create a slice with length 0 and capacity 5
	printSlice("b", b)

	c := b[:2] // Create a slice of 'b' from 0:2
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Println("\n* Nil Slices *\n")
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func nilSlice() {
	var z []int // Default value of slice is 'nil'. nil len = 0, cap = 0
	fmt.Println(z, len(z), cap(z))

	if z == nil {
		fmt.Println("Slice is null!")
	}
}

func rangeSlice() {
	fmt.Println("\n* Range Slices *\n")
	// Range iterate over elements of a slice or a map
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	fmt.Println("\n* Range Slices Without Index/Value *\n")
	pow = make([]int, 10)

	// You can ignore the value
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	// You can forget the index with an underscore '_'
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
