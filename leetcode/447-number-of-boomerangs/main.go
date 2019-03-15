package main

import (
	"fmt"
)

func numberOfBoomerangs(points [][]int) int {
	var boomerangs int

	for i := 0; i < len(points); i++ {
		distances := make(map[int]int, len(points)-1)
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			dist := (points[i][0]-points[j][0])*(points[i][0]-points[j][0]) + (points[i][1]-points[j][1])*(points[i][1]-points[j][1])
			distances[dist] += 1
		}

		for _, cnt := range distances {
			boomerangs += cnt * (cnt - 1)
		}

	}

	return boomerangs
}

func main() {
	var points [][]int

	// example
	points = append(points, []int{0, 0})
	points = append(points, []int{1, 0})
	points = append(points, []int{-1, 0})
	points = append(points, []int{0, 1})
	points = append(points, []int{0, -1})

	fmt.Println(numberOfBoomerangs(points))
}
