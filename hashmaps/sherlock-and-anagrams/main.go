package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	var count int32

	for l := 1; l < len(s); l++ {
		for i := 0; i < len(s)-l; i++ {
			substr1 := s[i : i+l]
			for j := i + 1; j < len(s)-l+1; j++ {
				substr2 := s[j : j+l]
				if checkSubstrs(substr1, substr2) {
					count++
				}
			}
		}
	}

	return count
}

func checkSubstrs(a string, b string) bool {
	first := make(map[uint8]int)
	second := make(map[uint8]int)
	for i := 0; i < len(a); i++ {
		first[a[i]-97]++
		second[b[i]-97]++
	}

	var c uint8 = 0
	for c <= 26 {
		if first[c] != second[c] {
			return false
		}
		c++
	}

	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
