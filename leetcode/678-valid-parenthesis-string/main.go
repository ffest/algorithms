package main

import "fmt"

// Stack based
/*func checkValidString(s string) bool {
	starsStack, opensStack := make([]int, 0), make([]int, 0)
	for i := range s {
		if s[i] == '*' {
			starsStack = append(starsStack, i)
		} else if s[i] == '(' {
			opensStack = append(opensStack, i)
		} else {
			if len(opensStack) == 0 && len(starsStack) == 0 {
				return false
			} else if len(opensStack) > 0 {
				opensStack = opensStack[:len(opensStack)-1]
			} else {
				starsStack = starsStack[:len(starsStack)-1]
			}
		}
	}
	for len(opensStack) > 0 {
		if len(starsStack) == 0 || starsStack[len(starsStack)-1] < opensStack[len(opensStack)-1] {
			return false
		}
		starsStack = starsStack[:len(starsStack)-1]
		opensStack = opensStack[:len(opensStack)-1]
	}
	return true
}*/

// Count the Open Parenthesis
func checkValidString(s string) bool {
	var cmin, cmax int
	for i := range s {
		if s[i] == '(' {
			cmax++
			cmin++
		} else if s[i] == ')' {
			cmax--
			cmin = max(cmin-1, 0)
		} else {
			cmax++
			cmin = max(cmin-1, 0)
		}
		if cmax < 0 {
			return false
		}
	}
	return cmin == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "(*)**()((()(*)"
	fmt.Println(checkValidString(s))
}
