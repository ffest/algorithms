package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func numJewelsInStones(J string, S string) int {
	var count int
	jewels := make(map[rune]struct{})
	for _, r := range J {
		jewels[r] = struct{}{}
	}

	for _, r := range S {
		if _, ok := jewels[r]; ok {
			count++
		}
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	fmt.Println(numJewelsInStones(nr[0], nr[1]))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
