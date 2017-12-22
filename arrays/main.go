package main

import (
	"fmt"
)

const gCap int = 5 // Capacity of list
var gGroceries [gCap]string
var gLen int // Total numbers of grocery items in our current array

func main() {

	addGrocery("Bread")
	addGrocery("Cucumbers")
	addGrocery("Salt")
	addGrocery("Fruit Cake")
	addGrocery("Pokemon Game")
	addGrocery("Ice Cream")

	listGroceries()
}

func addGrocery(grocery string) {

	if gLen < gCap {
		gGroceries[gLen] = grocery
		gLen++

		return
	}

	fmt.Println("You can't add more items, we are full")
	return
}

func listGroceries() {
	fmt.Println("Grocery List:")

	for i := 0; i < gLen; i++ {
		fmt.Println(gGroceries[i])
	}
}
