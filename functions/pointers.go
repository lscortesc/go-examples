package main

import "fmt"

func pointer(p *int) {
	fmt.Println(&p)
	fmt.Println(*p)
	*p = 100
}
