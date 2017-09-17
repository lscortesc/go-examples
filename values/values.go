// Go has various value types including strings, integers, floats, booleans, etc.
package main

import "fmt"

func main() {
	// Operator "sum" in strings works like a concat operator
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
