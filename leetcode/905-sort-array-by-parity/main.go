package main

import (
	"fmt"
)

func sortArrayByParity(A []int) []int {
	elementMap := make(map[int]int, len(A))
	for _, element := range A {
		elementMap[element]++
	}
	cursorLeft := 0
	cursorRight := len(A) - 1
	for el, count := range elementMap {
		if el%2 == 0 {
			for count > 0 {
				A[cursorLeft] = el
				count--
				cursorLeft++
			}
		} else {
			for count > 0 {
				A[cursorRight] = el
				count--
				cursorRight--
			}
		}
	}
	return A
}

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 2, 3, 5, 8, 7, 4, 2, 12, 33, 2, 1, 2, 4}))
}
