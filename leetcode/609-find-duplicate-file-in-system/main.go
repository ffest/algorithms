package main

import (
	"fmt"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	cache := make(map[string][]string)
	for _, path := range paths {
		parts := strings.Split(path, " ")
		dir := parts[0]
		for i := 1; i < len(parts); i++ {
			bracketPosition := strings.IndexByte(parts[i], '(')
			content := parts[i][bracketPosition+1 : len(parts[i])-1]
			cache[content] = append(cache[content], dir+"/"+parts[i][:bracketPosition])
		}
	}
	result := make([][]string, 0, len(cache))
	for _, group := range cache {
		if len(group) < 2 {
			continue
		}
		result = append(result, group)
	}
	return result
}

func main() {
	paths := []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}
	fmt.Println(findDuplicate(paths))
}
