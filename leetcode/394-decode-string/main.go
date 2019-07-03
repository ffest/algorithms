package main

import (
	"fmt"
	"strings"
)

func decodeString(s string) string {
	stackNums := make([]int, 0)
	stackStr := make([]string, 0)
	var res string
	var num int
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			num = 10*num + int(s[i]) - '0'
		case s[i] == '[':
			stackNums = append(stackNums, num)
			num = 0
			stackStr = append(stackStr, res)
			res = ""
		case s[i] == ']':
			tmp := stackStr[len(stackStr)-1]
			stackStr = stackStr[:len(stackStr)-1]
			count := stackNums[len(stackNums)-1]
			stackNums = stackNums[:len(stackNums)-1]
			res = tmp + strings.Repeat(res, count)
		default:
			res += string(s[i])
		}
	}
	return res
}

func main() {
	s := "3[a2[c]]"
	fmt.Println(decodeString(s))
}
