package main

import "fmt"

// TODO: make in faster. O(n^2) solution
func oddEvenJumps(arr []int) int {
	var count int
	odds := make(map[int]int)
	evens := make(map[int]int)

	for i := range arr {
		temp := make([]int, 0)
		position := i
		stepNumber := 1
		for position != len(arr)-1 {
			nextPosition := findNextStep(arr, position, stepNumber, odds, evens)
			if nextPosition == position {
				break
			}
			stepNumber++
			position = nextPosition

			temp = append(temp, nextPosition)
		}
		if position == len(arr)-1 {
			for i := range temp {
				if i%2 == 0 {
					odds[temp[i]] = len(arr) - 1
				} else {
					evens[temp[i]] = len(arr) - 1
				}
			}
			count++
		}
	}
	return count
}

func findNextStep(arr []int, position, stepNum int, odds, evens map[int]int) int {
	bestPosition := position
	value := -1
	if stepNum%2 == 0 {
		if nextStep, ok := odds[position]; ok {
			return nextStep
		}
		for i := position + 1; i < len(arr); i++ {
			if arr[i] > arr[position] {
				continue
			}
			if value == -1 || arr[i] > value {
				value = arr[i]
				bestPosition = i
			}
		}
	} else {
		if nextStep, ok := evens[position]; ok {
			return nextStep
		}
		for i := position + 1; i < len(arr); i++ {
			if arr[i] < arr[position] {
				continue
			}
			if value == -1 || arr[i] < value {
				value = arr[i]
				bestPosition = i
			}
		}
	}

	return bestPosition
}

func main() {
	input := []int{2, 3, 1, 1, 4}
	fmt.Println(oddEvenJumps(input))
}
