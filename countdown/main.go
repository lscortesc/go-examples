// Happy new year countdown
package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 10; i >= 1; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("Happy New Year!")
}
