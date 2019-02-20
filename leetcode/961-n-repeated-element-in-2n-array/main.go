package main

import (
	"fmt"
)

func repeatedNTimes(A []int) int {
	elementMap := make(map[int]struct{})
	for _, element := range A {
		if _, ok := elementMap[element]; ok {
			return element
		}
		elementMap[element] = struct{}{}
	}
	return 0
}

func main() {
	input := []int{5, 1, 5, 2, 5, 3, 5, 4}
	fmt.Println(repeatedNTimes(input))
}
