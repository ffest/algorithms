package main

import "fmt"

func missingNumber(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return len(nums)*(len(nums)+1)/2 - sum
}

func main() {
	nums := []int{0, 1}
	fmt.Println(missingNumber(nums))
}
