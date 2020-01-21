package main

import (
	"fmt"
	"sort"
)

func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) <= 1 {
		return true
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	finish := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < finish {
			return false
		}
		finish = intervals[i][1]
	}
	return true
}

func main() {
	intervals := [][]int{
		{0, 30},
		{5, 10},
		{15, 20},
	}
	fmt.Println(canAttendMeetings(intervals))
}
