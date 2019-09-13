package main

import (
	"fmt"
)

// TODO: we can do it with slice and use binary search
type LinkedList struct {
	value     string
	timestamp int
	next      *LinkedList
}

type TimeMap struct {
	cache map[string]*LinkedList
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		cache: make(map[string]*LinkedList),
	}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	cached := tm.cache[key]
	tm.cache[key] = &LinkedList{
		value:     value,
		timestamp: timestamp,
		next:      cached,
	}
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	cached := tm.cache[key]
	if cached == nil {
		return ""
	}
	for cached.next != nil && cached.timestamp > timestamp {
		cached = cached.next
	}
	if cached.timestamp <= timestamp {
		return cached.value
	}
	return ""
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
