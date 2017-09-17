// Variables are explicitly declared and used by the compliler
package main

import "fmt"

func main() {
	var a = "initial" // var declares 1 or more variables
	fmt.Println(a)

	var b, c int = 1, 2 // We can declare multiple variables at once
	fmt.Println(b, c)

	var d = true // Go will infer the type of initialized variable
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued
	// For example, the zero value for an 'int' is 0
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable
	// e.g. for var f string = "short" in this case
	f := "short"
	fmt.Println(f)
}
