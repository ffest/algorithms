package main

import "fmt"

func diStringMatch(S string) []int {
	result := make([]int, 0, len(S)+1)

	min := 0
	max := len(S)

	for _, letter := range S {
		switch letter {
		case 'D':
			result = append(result, max)
			max--
		case 'I':
			result = append(result, min)
			min++
		}
	}
	result = append(result, min)
	return result
}

func main() {
	input := "DIDI"
	fmt.Println(diStringMatch(input))
}
