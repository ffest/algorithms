package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	stack := make([]string, 0)
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if part == "" || part == "." {
			continue
		} else if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, part)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	path := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(path))
}
