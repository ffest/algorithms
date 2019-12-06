package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type mh [][]int

func (mh mh) Len() int           { return len(mh) }
func (mh mh) Less(i, j int) bool { return mh[i][2] < mh[j][2] }
func (mh mh) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }

func (mh *mh) Push(x interface{}) {
	*mh = append(*mh, x.([]int))
}

func (mh *mh) Pop() interface{} {
	var x interface{}
	x, *mh = (*mh)[len(*mh)-1], (*mh)[:len(*mh)-1]
	return x
}

func carPooling(trips [][]int, capacity int) bool {
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})

	mh := mh{}
	heap.Init(&mh)

	for _, trip := range trips {
		heap.Push(&mh, trip)
		for mh[0][2] <= trip[1] {
			capacity += mh[0][0]
			heap.Remove(&mh, 0)
		}
		capacity -= trip[0]
		if capacity < 0 {
			return false
		}
	}
	return true
}

// Just saving every start/end
/*func carPooling(trips [][]int, capacity int) bool {
	stops := [1001]int{}
	for _, trip := range trips {
		stops[trip[1]] -= trip[0]
		stops[trip[2]] += trip[0]
	}
	for i := 0; i < 1001; i++ {
		capacity += stops[i]
		if capacity < 0 {
			return false
		}
	}
	return true
}*/

func main() {
	capacity := 11
	trips := [][]int{
		{3, 2, 7},
		{3, 7, 9},
		{8, 3, 9},
	}
	fmt.Println(carPooling(trips, capacity))
}
