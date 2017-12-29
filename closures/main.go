package main

import (
	"fmt"
	"math"
)

func main() {
	hypot()
	closures()
	showFibonacci()
}

func hypot() {
	fmt.Println("\n*** Hypotenuse ***\n")

	hypot := func(a, b float64) float64 {
		return math.Sqrt(a*a + b*b)
	}

	fmt.Println(hypot(3.4, 4.123))
}

func closures() {
	fmt.Println("\n*** Closures ***\n")

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func showFibonacci() {
	fmt.Println("\n*** Fibonacci Series ***\n")
	f := fibonacci()

	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	current, next := 0, 1
	return func() int {
		current, next = next, current+next
		return current
	}
}

func adder() func(x int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
