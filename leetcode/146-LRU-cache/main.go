package main

import (
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]int
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]int),
	}
}

func (c *LRUCache) Get(key int) int {
	val, ok := c.cache[key]
	if ok {
		return val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	c.cache[key] = value
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4
}
