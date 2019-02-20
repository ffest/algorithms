package main

import "fmt"

func judgeCircle(moves string) bool {
	var ud int8
	var lr int8
	for _, move := range moves {
		switch move {
		case 'U':
			ud++
		case 'D':
			ud--
		case 'L':
			lr--
		case 'R':
			lr++
		}
	}

	return ud == 0 && lr == 0
}

func main() {
	input := "UDLR"
	fmt.Println(judgeCircle(input))
}
