package main

import (
	"fmt"
)

func partitionLabels(s string) []int {
	ends := make(map[byte]int)
	for i := range s {
		ends[s[i]] = i
	}
	partitions := make([]int, 0)
	var prev, current int
	for i := range s {
		end := ends[s[i]]
		if end > current {
			current = end
		}
		if i == current {
			partitions = append(partitions, current-prev+1)
			prev = current + 1
			continue
		}
	}
	return partitions
}

func main() {
	s := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(s))
}
