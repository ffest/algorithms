package main

import "fmt"

func fib(N int) int {
	if N < 2 {
		return N
	}
	first, second := 0, 1
	for i := 2; i < N; i++ {
		tmp := first + second
		first = second
		second = tmp
	}
	return first + second
}

func main() {
	N := 4
	fmt.Println(fib(N))
}
