package main

import (
	"fmt"
)

func isLongPressedName(name string, typed string) bool {
	if len(typed) < len(name) {
		return false
	}
	nameIter := 0
	typedIter := 0
	for nameIter < len(name) && typedIter < len(typed) {
		if typed[typedIter] == name[nameIter] {
			nameIter++
		}
		typedIter++
	}

	return nameIter == len(name)
}

func main() {
	name := "alex"
	typed := "aaleexd"
	fmt.Println(isLongPressedName(name, typed))
}
