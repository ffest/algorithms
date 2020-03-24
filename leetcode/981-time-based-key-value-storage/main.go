package main

import (
	"fmt"
)

type Data struct {
	value     string
	timestamp int
}

type TimeMap struct {
	cache map[string][]Data
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		cache: make(map[string][]Data),
	}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.cache[key] = append(tm.cache[key], Data{value, timestamp})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	cached := tm.cache[key]
	if cached == nil {
		return ""
	}
	left, right := 0, len(cached)
	for left < right {
		mid := (left + right) / 2
		if cached[mid].timestamp <= timestamp {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if right == 0 {
		return ""
	}
	return cached[right-1].value
}

func main() {
	tm := Constructor()
	tm.Set("love", "high", 10)
	tm.Set("love", "low", 20)
	fmt.Println(tm.Get("love", 5))
	fmt.Println(tm.Get("love", 10))
	fmt.Println(tm.Get("love", 15))
	fmt.Println(tm.Get("love", 20))
	fmt.Println(tm.Get("love", 25))
}
