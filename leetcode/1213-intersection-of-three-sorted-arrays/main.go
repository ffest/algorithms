package main

import "fmt"

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	intersection := make([]int, 0)
	for len(arr1) > 0 && len(arr2) > 0 && len(arr3) > 0 {
		if arr1[0] == arr2[0] && arr2[0] == arr3[0] {
			intersection = append(intersection, arr1[0])
			arr1 = arr1[1:]
			arr2 = arr2[1:]
			arr3 = arr3[1:]
		} else {
			currentMax := max(arr1[0], arr2[0])
			currentMax = max(currentMax, arr3[0])
			if arr1[0] < currentMax {
				arr1 = arr1[1:]
			}
			if arr2[0] < currentMax {
				arr2 = arr2[1:]
			}
			if arr3[0] < currentMax {
				arr3 = arr3[1:]
			}
		}
	}
	return intersection
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 5, 7, 9}
	arr3 := []int{1, 3, 4, 5, 8}
	fmt.Println(arraysIntersection(arr1, arr2, arr3))
}
