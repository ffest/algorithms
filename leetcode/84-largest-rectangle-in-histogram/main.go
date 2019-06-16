package main

import "fmt"

func largestRectangleArea(heights []int) int {
	var largest int
	var tmpHeight, tmpPosition, tmpSize int
	pStack := make([]int, 0)
	hStack := make([]int, 0)

	for i := 0; i < len(heights); i++ {
		if len(hStack) == 0 || hStack[len(hStack)-1] < heights[i] {
			hStack = append(hStack, heights[i])
			pStack = append(pStack, i)
		} else if hStack[len(hStack)-1] > heights[i] {
			for len(hStack) > 0 && hStack[len(hStack)-1] > heights[i] {
				tmpHeight = hStack[len(hStack)-1]
				hStack = hStack[:len(hStack)-1]
				tmpPosition = pStack[len(pStack)-1]
				pStack = pStack[:len(pStack)-1]
				tmpSize = tmpHeight * (i - tmpPosition)
				if tmpSize > largest {
					largest = tmpSize
				}
			}

			hStack = append(hStack, heights[i])
			pStack = append(pStack, tmpPosition)
		}
	}

	for len(hStack) > 0 {
		tmpHeight = hStack[len(hStack)-1]
		hStack = hStack[:len(hStack)-1]
		tmpPosition = pStack[len(pStack)-1]
		pStack = pStack[:len(pStack)-1]
		tmpSize = tmpHeight * (len(heights) - tmpPosition)
		if tmpSize > largest {
			largest = tmpSize
		}
	}

	return largest
}

func main() {
	input := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(input))
}
