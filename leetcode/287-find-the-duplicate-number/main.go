package main

import (
	"fmt"
	"sort"
)

// Solution with array modification
/*func findDuplicate(nums []int) int {
	for i := range nums {
		tmp := int(math.Abs(float64(nums[i])))
		if nums[tmp-1] < 0 {
			return int(math.Abs(float64(nums[i])))
		}
		nums[tmp-1] = -nums[tmp-1]
	}
	return -1
}*/

// We can do kinda binary search here. Count the number than greater than mid and move borders
func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return 0
}

func main() {
	nums := []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicate(nums))
}
