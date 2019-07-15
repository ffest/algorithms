package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	var number string
	sort.Slice(nums, func(i, j int) bool {
		first, second := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		if first+second >= second+first {
			return true
		}
		return false
	})
	for _, num := range nums {
		number += strconv.Itoa(num)
	}
	if number[0] == '0' {
		return number[:1]
	}
	return number
}

func main() {
	nums := []int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247}
	fmt.Println(largestNumber(nums))
}
