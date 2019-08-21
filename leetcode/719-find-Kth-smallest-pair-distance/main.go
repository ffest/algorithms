package main

import (
	"fmt"
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

func smallestDistancePair(nums []int, k int) int {
	return 0
}

func main() {
	nums := []int{9, 10, 7, 10, 6, 1, 5, 4, 9, 8}
	k := 18
	fmt.Println(smallestDistancePair(nums, k))
}
