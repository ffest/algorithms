package main

import (
	"fmt"
	"strconv"
)

func primePalindrome(N int) int {
	if N >= 8 && N <= 11 {
		return 11
	}

	for i := 1; i < 10000; i++ {
		numStr := strconv.Itoa(i)
		numStr += reverseStr(numStr)[1:]
		num, _ := strconv.Atoi(numStr)
		if num >= N && isPrime(num) {
			return num
		}
	}
	return -1
}

func reverseStr(s string) string {
	str := []byte(s)
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]
	}
	return string(str)
}

func isPrime(n int) bool {
	if n < 2 || n%2 == 0 {
		return n == 2
	}
	for i := 3; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	N := 9989900
	fmt.Println(primePalindrome(N))
}
