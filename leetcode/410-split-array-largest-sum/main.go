package main

import (
	"fmt"
)

func splitArray(nums []int, m int) int {
	var max, sum int
	for _, n := range nums {
		sum += n
		if n > max {
			max = n
		}
	}
	if m == 1 {
		return sum
	}

	left, right := max, sum
	for left <= right {
		mid := (left + right) / 2
		if helper(mid, m, nums) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func helper(target, m int, nums []int) bool {
	var sum int
	count := 1
	for _, n := range nums {
		sum += n
		if sum > target {
			sum = n
			count++
			if count > m {
				return false
			}
		}
	}
	return true
}

func main() {
	nums := []int{7, 2, 5, 10, 8}
	fmt.Println(splitArray(nums, 4))
}
