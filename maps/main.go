package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
	showMap()
	literalMaps()
	mapOperations()

	mapExercise()
}

func showMap() {
	// Maps are created with 'make' and never with 'new'
	m := make(map[string]Vertex)
	m["BellLabs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["BellLabs"])
}

func literalMaps() {
	// Like literal structs but the keys are mandatory
	var m = map[string]Vertex{
		"BellLabs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

	// If top name is of a type, you can omit it of the elements
	var ma = map[string]Vertex{
		"BellLabs": {40.68433, -74.39967},
		"Google":   {37.42202, -122.08408},
	}
	fmt.Println(ma)
}

func mapOperations() {
	m := make(map[string]int)

	m["Response"] = 42
	fmt.Println("Response:", m["Response"])

	m["Response"] = 48
	fmt.Println("Response:", m["Response"])

	delete(m, "Response")
	fmt.Println("Response", m["Response"])

	v, ok := m["Response"]
	fmt.Println("Value:", v, "Exists?", ok)
}
