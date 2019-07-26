package main

import (
	"fmt"
	"sort"
)

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{
		k:    k,
		nums: nums,
	}
}

func (kl *KthLargest) Add(val int) int {
	place := kl.binary(val)
	kl.nums = append(kl.nums[:place], append([]int{val}, kl.nums[place:]...)...)
	return kl.nums[len(kl.nums)-kl.k]
}

func (kl *KthLargest) binary(key int) int {
	left := 0
	right := len(kl.nums)
	for left < right {
		median := (left + right) / 2
		if key < kl.nums[median] {
			right = median
		} else if key > kl.nums[median] {
			left = median + 1
		} else {
			return median
		}
	}
	return left
}

func main() {
	k := 3
	arr := []int{4, 5, 8, 2}
	kl := Constructor(k, arr)
	fmt.Println(kl.Add(3))
	fmt.Println(kl.Add(5))
	fmt.Println(kl.Add(10))
	fmt.Println(kl.Add(9))
	fmt.Println(kl.Add(4))
}
