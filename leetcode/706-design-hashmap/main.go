package main

import "fmt"

type MyHashMap struct {
	capacity int
	data     []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	capacity := 1000001
	data := make([]int, capacity)
	for i := range data {
		data[i] = -1
	}
	return MyHashMap{
		capacity: capacity,
		data:     data,
	}
}

func (hm *MyHashMap) Hash(key int) int {
	return key % hm.capacity
}

/** value will always be non-negative. */
func (hm *MyHashMap) Put(key int, value int) {
	hm.data[hm.Hash(key)] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (hm *MyHashMap) Get(key int) int {
	return hm.data[hm.Hash(key)]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (hm *MyHashMap) Remove(key int) {
	hm.data[hm.Hash(key)] = -1
}

func main() {
	hm := Constructor()
	hm.Put(1, 1)
	hm.Put(2, 2)
	fmt.Println(hm.Get(1)) //1
	fmt.Println(hm.Get(3)) //-1
	hm.Put(2, 1)
	fmt.Println(hm.Get(2)) //1
	hm.Remove(2)
	fmt.Println(hm.Get(2)) //-1
}
