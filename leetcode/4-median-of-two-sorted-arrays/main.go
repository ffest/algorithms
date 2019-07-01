package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	var maxLeft, minRight int
	left, right, half := 0, len(nums1), (len(nums1)+len(nums2)+1)/2
	var i, j int
	for left <= right {
		i = (left + right) / 2
		j = half - i
		if i < len(nums1) && nums2[j-1] > nums1[i] {
			left = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else if nums1[i-1] > nums2[j-1] {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = nums2[j-1]
			}
			if (len(nums1)+len(nums2))%2 != 0 {
				return float64(maxLeft)
			}
			if i == len(nums1) {
				minRight = nums2[j]
			} else if j == len(nums2) {
				minRight = nums1[i]
			} else if nums1[i] < nums2[j] {
				minRight = nums1[i]
			} else {
				minRight = nums2[j]
			}
			return float64(maxLeft+minRight) / 2
		}
	}
	return 0
}

func main() {
	nums1 := []int{1, 2, 4, 5, 6, 8}
	nums2 := []int{1, 7}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
