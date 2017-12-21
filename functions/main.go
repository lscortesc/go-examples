package main

import (
	"fmt"
)

func main() {

	// Functions
	j := 15
	fmt.Println("In main - j is:", j, "Located at:", &j)
	resetValue(j)
	fmt.Println("In main - j is:", j, "Located at:", &j)

	// Max Function
	a, b := 20, 10
	fmt.Println("Max number between", a, "&", b, "is:", maximo(a, b))

	c, d := 1000, 40
	fmt.Println("Max number between", c, "&", d, "is:", maximo(c, d))
}

func resetValue(i int) {
	fmt.Println("In resetValue() - i is:", i, "Located at:", &i)
	i = 0
	fmt.Println("After i=0 resetValue() - i is:", i, "Located at:", &i)
}
