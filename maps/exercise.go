package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func mapExercise() {
	fmt.Println("\n** Exercise Maps **\n")

	wc.Test(counterWords)
}

func counterWords(s string) map[string]int {
	words := strings.Fields(s)
	counter := make(map[string]int)

	for _, word := range words {
		_, ok := counter[word]

		if ok {
			counter[word]++
			continue
		}

		counter[word] = 1
	}

	return counter
}
