package main

import (
	"fmt"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	starts, ends := make([]int, len(intervals)), make([]int, len(intervals))
	for i := range intervals {
		starts[i] = intervals[i][0]
		ends[i] = intervals[i][1]
	}
	sort.Ints(starts)
	sort.Ints(ends)

	var rooms, endIter int
	for i := range starts {
		if starts[i] < ends[endIter] {
			rooms++
		} else {
			endIter++
		}
	}
	return rooms
}

func main() {
	intervals := [][]int{
		{0, 30},
		{5, 10},
		{15, 20},
		{6, 11},
	}
	fmt.Println(minMeetingRooms(intervals))
}
