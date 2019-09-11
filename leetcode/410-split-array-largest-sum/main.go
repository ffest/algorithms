package main

import (
	"fmt"
)

// Binary search
func splitArray(nums []int, m int) int {
	var max, sum int
	for _, num := range nums {
		sum += num
		if max < num {
			num = max
		}
	}
	if m == 1 {
		return sum
	}

	left, right := max, sum
	for left <= right {
		mid := (left + right) / 2
		if helper(nums, mid, m) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func helper(nums []int, target, m int) bool {
	var sum, counter int
	for _, num := range nums {
		sum += num
		if sum > target {
			counter++
			sum = num
			if counter == m {
				return false
			}
		}
	}

	return true
}

func main() {
	nums := []int{7, 2, 5, 10, 8}
	fmt.Println(splitArray(nums, 2))
}
