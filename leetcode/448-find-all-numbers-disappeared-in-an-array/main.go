package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	result := make([]int, 0)
	for i := range nums {
		tmp := nums[i]
		if tmp < 0 {
			tmp *= -1
		}
		if nums[tmp-1] > 0 {
			nums[tmp-1] *= -1
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Print(findDisappearedNumbers(nums))
}
