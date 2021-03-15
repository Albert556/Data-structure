/*
 * @lc app=leetcode.cn id=706 lang=golang
 *
 * [706] 设计哈希映射
 */
package main

import "container/list"

// @lc code=start
type element struct {
	key   int
	value int
}

type MyHashMap struct {
	base int
	// 双向链表
	hashMap []list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	myHashMap := MyHashMap{769, make([]list.List, 769)}
	return myHashMap
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	k := key % this.base
	for e := this.hashMap[k].Front(); e != nil; e = e.Next() {
		if e.Value.(element).key == key {
			e.Value = element{key, value}
			return
		}
	}
	this.hashMap[k].PushBack(element{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	k := key % this.base
	for e := this.hashMap[k].Front(); e != nil; e = e.Next() {
		if e.Value.(element).key == key {
			return e.Value.(element).value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	k := key % this.base
	for e := this.hashMap[k].Front(); e != nil; e = e.Next() {
		if e.Value.(element).key == key {
			this.hashMap[k].Remove(e)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end
