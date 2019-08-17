package main

import (
	"fmt"
	"math"
)

type bucket struct {
	max int
	min int
}

// TODO: refactor it
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	if len(nums) == 2 {
		if nums[1] > nums[0] {
			return nums[1] - nums[0]
		} else {
			return nums[0] - nums[1]
		}
	}
	var max, min = nums[0], nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
		if min > num {
			min = num
		}
	}

	bucketSize := int(math.Ceil(float64(max-min) / float64(len(nums)-1)))
	if bucketSize == 0 {
		return 0
	}

	buckets := make([]bucket, len(nums)-1)
	for _, num := range nums {
		idx := (num - min) / bucketSize

		if num == max && (num-min)%bucketSize == 0 {
			idx -= 1
		}

		if buckets[idx].min == 0 || buckets[idx].min > num {
			buckets[idx].min = num
		}
		if buckets[idx].max == 0 || buckets[idx].max < num {
			buckets[idx].max = num
		}
	}

	validBuckets := make([]bucket, 0)
	var gap int
	for _, b := range buckets {
		if b.max != 0 && b.min != 0 {
			validBuckets = append(validBuckets, b)
		}
	}

	for i, b := range validBuckets {
		if i == len(validBuckets)-1 {
			if gap == 0 || b.max-b.min > gap {
				gap = b.max - b.min
			}
		} else {
			if gap == 0 || b.max-b.min > gap {
				gap = b.max - b.min
			}
			if gap == 0 || validBuckets[i+1].min-b.max > gap {
				gap = validBuckets[i+1].min - b.max
			}
		}
	}
	return gap
}

func main() {
	input := []int{100, 3, 2, 1}
	fmt.Println(maximumGap(input))
}
