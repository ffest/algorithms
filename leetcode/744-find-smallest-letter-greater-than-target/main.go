package main

import (
	"fmt"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	left := 0
	right := len(letters)
	if target >= letters[len(letters)-1] {
		return letters[0]
	} else {
		target++
	}
	for {
		median := (left + right) / 2
		if left == right {
			return letters[right]
		} else if letters[median] == target {
			return letters[median]
		} else if letters[median] < target {
			left = median + 1
		} else {
			right = median
		}
	}
}

func main() {
	input := []byte{'c', 'f', 'j'}
	target := byte('c')
	fmt.Println(string(nextGreatestLetter(input, target)))
}
