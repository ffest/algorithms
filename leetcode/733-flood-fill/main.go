package main

import (
	"fmt"
)

func floodFill(image [][]int, sr, sc int, newColor int) [][]int {
	n, m := len(image), len(image[0])
	visited := make(map[[2]int]struct{})

	queue := make([][2]int, 0)
	queue = append(queue, [2]int{sr, sc})

	for len(queue) > 0 {
		pixel := queue[0]
		queue = queue[1:]
		if _, ok := visited[pixel]; ok {
			continue
		}
		if pixel[0] < n-1 && image[pixel[0]][pixel[1]] == image[pixel[0]+1][pixel[1]] {
			queue = append(queue, [2]int{pixel[0] + 1, pixel[1]})
		}
		if pixel[0] > 0 && image[pixel[0]][pixel[1]] == image[pixel[0]-1][pixel[1]] {
			queue = append(queue, [2]int{pixel[0] - 1, pixel[1]})
		}
		if pixel[1] > 0 && image[pixel[0]][pixel[1]] == image[pixel[0]][pixel[1]-1] {
			queue = append(queue, [2]int{pixel[0], pixel[1] - 1})
		}
		if pixel[1] < m-1 && image[pixel[0]][pixel[1]] == image[pixel[0]][pixel[1]+1] {
			queue = append(queue, [2]int{pixel[0], pixel[1] + 1})
		}
		image[pixel[0]][pixel[1]] = newColor
		visited[pixel] = struct{}{}
	}
	return image
}

func main() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	sr := 1
	sc := 1
	newColor := 2
	fmt.Println(floodFill(image, sr, sc, newColor))
}
