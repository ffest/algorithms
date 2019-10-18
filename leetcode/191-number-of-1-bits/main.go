package main

import "fmt"

func hammingWeight(num uint32) int {
	var count uint32
	for num > 0 {
		count += num & 1
		num = num >> 1
	}
	return int(count)
}

func main() {
	num := uint32(8)
	fmt.Println(hammingWeight(num))
}
