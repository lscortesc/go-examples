package main

import (
	"fmt"
)

var groceries []string

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
	fmt.Println("Capacity", cap(groceries))
	groceries = append(groceries, grocery)
}

func listGroceries() {
	fmt.Println("Grocery list:")

	for i := 0; i < len(groceries); i++ {
		fmt.Println(groceries[i])
	}
}
