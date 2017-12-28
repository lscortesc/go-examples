package main

import "fmt"

var groceries []string

func main() {
	addGrocery("Bread")
	addGrocery("Cucumbers")
	addGrocery("Salt")
	addGrocery("Fruit Cake")
	addGrocery("Pokemon Game")
	addGrocery("Ice Cream")

	addGroceries("Corn", "Peas", "Another")

	listGroceries()

	antoherSlice()
	slicingSlices()
	createSlices()
	nilSlice()
	rangeSlice()
	slicesExercise()
}

func addGrocery(grocery string) {
	groceries = append(groceries, grocery)
}

func listGroceries() {
	fmt.Println("Grocery list:")

	for i := 0; i < len(groceries); i++ {
		fmt.Println(groceries[i])
	}

	for _, grocery := range groceries {
		fmt.Println(grocery)
	}
}

func addGroceries(list ...string) {

	for _, grocery := range list {
		addGrocery(grocery)
	}
}
