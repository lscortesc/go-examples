package main

import (
	"fmt"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // Sum sends to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c // Recive from c

	fmt.Println(x, y, x+y)

	buffer()
}

func buffer() {
	c := make(chan int, 3)
	c <- 1
	c <- 10

	fmt.Println(<-c)
	fmt.Println(<-c)
}
