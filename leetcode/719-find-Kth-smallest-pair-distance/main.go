package main

import (
	"fmt"
	"math"
	"sort"
)

// min heap solution O(n^2 logK) -- too slow
/*type distances []int

func (d distances) Len() int           { return len(d) }
func (d distances) Less(i, j int) bool { return d[i] > d[j] }
func (d distances) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func (d *distances) Push(x interface{}) {
	*d = append(*d, x.(int))
}

func (d *distances) Pop() interface{} {
	var x interface{}
	x, *d = (*d)[len(*d)-1], (*d)[:len(*d)-1]
	return x
}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	distances := distances{}
	heap.Init(&distances)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			d := int(math.Abs(float64(nums[j] - nums[i])))
			if distances.Len() < k {
				heap.Push(&distances, d)
			} else if d < distances[0] {
				distances[0] = d
				heap.Fix(&distances, 0)
			}
		}
	}
	return distances[0]
}*/

// Binary search O(nlogn)
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	left := math.MaxInt32
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] < left {
			left = nums[i] - nums[i-1]
		}
	}
	right := nums[len(nums)-1] - nums[0]

	for left < right {
		mid := (left + right) / 2
		count := 0

		for i, j := 0, 1; i < len(nums); i++ {
			for j < len(nums) && nums[j]-nums[i] <= mid {
				j++
			}
			count += j - i - 1
		}

		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func main() {
	nums := []int{9, 10, 7, 10, 6, 1, 5, 4, 9, 8}
	k := 18
	fmt.Println(smallestDistancePair(nums, k))
}
