package main

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	return atMost(nums, k) - atMost(nums, k-1)
}

func atMost(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}
	var count int
	var left, right, currentOdds int
	for right = 0; right < len(nums); right++ {
		if nums[right]%2 == 1 {
			currentOdds++
		}
		for currentOdds > k {
			if nums[left]%2 == 1 {
				currentOdds--
			}
			left++
		}
		count += right - left + 1
	}
	return count
}

func main() {
	nums := []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}
	k := 2
	fmt.Println(numberOfSubarrays(nums, k))
}
