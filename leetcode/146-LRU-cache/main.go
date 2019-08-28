package main

import (
	"fmt"
)

type DLinkedList struct {
	val  int
	key  int
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
	n, ok := c.cache[key]
	if !ok {
		n = &DLinkedList{val: value, key: key}
		c.cache[key] = n
	} else {
		n.val = value
		c.removeFromChain(n)
	}

	c.addToChain(n)
	if len(c.cache) > c.capacity {
		n = c.tail
		if n != nil {
			c.removeFromChain(n)
			delete(c.cache, n.key)
		}
	}
}

func (c *LRUCache) addToChain(l *DLinkedList) {
	l.prev = nil
	l.next = c.head
	if l.next != nil {
		l.next.prev = l
	}
	c.head = l
	if c.tail == nil {
		c.tail = l
	}
}

func (c *LRUCache) removeFromChain(l *DLinkedList) {
	if l == c.head {
		c.head = l.next
	}

	if l == c.tail {
		c.tail = l.prev
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
