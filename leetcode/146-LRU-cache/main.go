package main

import (
	"fmt"
)

type DLinkedList struct {
	key  int
	val  int
	prev *DLinkedList
	next *DLinkedList
}

type LRUCache struct {
	capacity int
	head     *DLinkedList
	tail     *DLinkedList
	cache    map[int]*DLinkedList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedList),
	}
}

func (c *LRUCache) Get(key int) int {
	l, ok := c.cache[key]
	if !ok {
		return -1
	}

	c.removeFromChain(l)
	c.addToChain(l)
	return l.val
}

func (c *LRUCache) Put(key int, value int) {
	l, ok := c.cache[key]
	if !ok {
		l = &DLinkedList{key: key, val: value}
		c.cache[key] = l
	} else {
		l.val = value
		c.removeFromChain(l)
	}
	c.addToChain(l)
	if len(c.cache) > c.capacity {
		l := c.tail
		c.removeFromChain(l)
		delete(c.cache, l.key)
	}
}

func (c *LRUCache) addToChain(l *DLinkedList) {
	l.next = nil
	if c.head != nil {
		c.head.next = l
		l.prev = c.head
	}
	c.head = l
	if c.tail == nil {
		c.tail = l
	}
}

func (c *LRUCache) removeFromChain(l *DLinkedList) {
	if l == c.head {
		c.head = l.prev
	}
	if l == c.tail {
		c.tail = l.next
	}
	if l.next != nil {
		l.next.prev = l.prev
	}
	if l.prev != nil {
		l.prev.next = l.next
	}
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
