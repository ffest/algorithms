package main

import (
	"fmt"
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		firstSpace, secondSpace := strings.Index(logs[i], " "), strings.Index(logs[j], " ")
		firstChar, secondChar := logs[i][firstSpace+1], logs[j][secondSpace+1]
		if firstChar <= '9' {
			if secondChar <= '9' {
				return i < j
			}
			return false
		}
		if secondChar <= '9' {
			return true
		}
		comparison := strings.Compare(logs[i][firstSpace+1:], logs[j][secondSpace+1:])
		if comparison == 0 {
			comparison = strings.Compare(logs[i][:firstSpace], logs[j][:secondSpace])
		}
		return comparison < 0
	})
	return logs
}

func main() {
	logs := []string{"8 fj dzz k", "5r 446 6 3", "63 gu psub", "5 ba iedrz", "6s 87979 5", "3r 587 01", "jc 3480612", "bb wsrd kp", "b aq cojj", "r5 6316 71"}
	fmt.Println(reorderLogFiles(logs))
}
