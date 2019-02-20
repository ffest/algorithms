package main

import (
	"fmt"
)

/*func toLowerCase(str string) string {
	for i, r := range str {
		if r >= 'A' && r <= 'Z' {
			r = r - 'A' + 'a'
			str = str[:i] + string(r) + str[i+1:]
		}
	}
	return str
}*/

func toLowerCase(str string) string {
	b := []byte(str)
	for i := 0; i < len(str); i++ {
		if str[i] >= 65 && str[i] <= 90 {
			b[i] = b[i] + 32
		}
	}
	return string(b)
}

func main() {
	input := "HeEllo"
	fmt.Println(toLowerCase(input))
}
