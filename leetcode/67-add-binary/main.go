package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	var result string
	var carry bool
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		sum := 0
		if i >= 0 && j >= 0 {
			if a[i] == '1' {
				sum++
			}
			if b[j] == '1' {
				sum++
			}
			i--
			j--
		} else if i >= 0 {
			if a[i] == '1' {
				sum++
			}
			i--
		} else {
			if b[j] == '1' {
				sum++
			}
			j--
		}
		if carry {
			sum++
		}
		result = string(strconv.Itoa(sum%2)) + result
		carry = sum > 1
	}
	if carry {
		result = "1" + result
	}
	return result
}

func main() {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a, b))
}
