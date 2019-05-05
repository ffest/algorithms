package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	var count int64
	arrMap3 := make(map[int64]int64)
	arrMap2 := make(map[int64]int64)

	for _, element := range arr {
		count += arrMap3[element]

		if arrMap2[element] > 0 {
			if arrMap3[element*r] > 0 {
				arrMap3[element*r] += arrMap2[element]
			} else {
				arrMap3[element*r] = arrMap2[element]
			}
		}

		if arrMap2[element*r] > 0 {
			arrMap2[element*r]++
		} else {
			arrMap2[element*r] = 1
		}
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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
