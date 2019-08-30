package main

import "fmt"

type MyHashSet struct {
	capacity int
	data     []bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	capacity := 1000001
	return MyHashSet{
		capacity: capacity,
		data:     make([]bool, capacity),
	}
}

func (hs *MyHashSet) Hash(key int) int {
	return key % hs.capacity
}

func (hs *MyHashSet) Add(key int) {
	hs.data[hs.Hash(key)] = true
}

func (hs *MyHashSet) Remove(key int) {
	hs.data[hs.Hash(key)] = false
}

/** Returns true if this set contains the specified element */
func (hs *MyHashSet) Contains(key int) bool {
	return hs.data[hs.Hash(key)]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
func main() {
	hs := Constructor()
	hs.Add(1)
	hs.Add(2)
	fmt.Println(hs.Contains(1)) //true
	fmt.Println(hs.Contains(3)) //false
	hs.Add(2)
	fmt.Println(hs.Contains(2)) //true
	hs.Remove(2)
	fmt.Println(hs.Contains(2)) //false
}
