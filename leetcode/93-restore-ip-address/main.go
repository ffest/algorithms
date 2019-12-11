package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	ips := make([]string, 0)

	parts := make([]string, 4)
	for i := 1; i < min(4, len(s)); i++ {
		parts[0] = s[:i]
		if !isValid(parts[0]) {
			continue
		}
		for j := 1; j < min(4, len(s)-i); j++ {
			parts[1] = s[i : i+j]
			if !isValid(parts[1]) {
				continue
			}
			for k := 1; k < min(4, len(s)-i-j); k++ {
				parts[2] = s[i+j : i+j+k]
				if !isValid(parts[2]) {
					continue
				}
				parts[3] = s[i+j+k:]
				if !isValid(parts[3]) {
					continue
				}
				ips = append(ips, strings.Join(parts, "."))
			}
		}
	}

	return ips
}

func isValid(part string) bool {
	if len(part) > 1 && part[0] == '0' {
		return false
	}
	raw, _ := strconv.Atoi(part)
	if raw > 255 {
		return false
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "010010"
	fmt.Println(restoreIpAddresses(s))
}
