package main

import (
	"fmt"
	"math"
)

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
		if max < num {
			max = num
		}
	}

	bucketSize := int(math.Ceil(float64(max-min) / float64(len(nums)-1)))

	bucketsMin, bucketsMax := make([]int, len(nums)-1), make([]int, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		bucketsMin[i], bucketsMax[i] = math.MaxInt32, math.MinInt32
	}

	for _, num := range nums {
		if num == min || num == max {
			continue
		}
		bucketIdx := (num - min) / bucketSize
		if num < bucketsMin[bucketIdx] {
			bucketsMin[bucketIdx] = num
		}
		if num > bucketsMax[bucketIdx] {
			bucketsMax[bucketIdx] = num
		}
	}

	maxGap := math.MinInt32
	prev := min
	for i := 0; i < len(nums)-1; i++ {
		if bucketsMin[i] == math.MaxInt32 && bucketsMax[i] == math.MinInt32 {
			//empty bucket
			continue
		}
		if bucketsMin[i]-prev > maxGap {
			maxGap = bucketsMin[i] - prev
		}
		prev = bucketsMax[i]
	}
	if max-prev > maxGap {
		maxGap = max - prev
	}

	return maxGap
}

func main() {
	input := []int{100, 3, 2, 1}
	fmt.Println(maximumGap(input))
}
