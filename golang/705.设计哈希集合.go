/*
 * @lc app=leetcode.cn id=705 lang=golang
 *
 * [705] 设计哈希集合
 */
package main

import (
	"container/list"
	"fmt"
)

// @lc code=start
type MyHashSet struct {
	base int
	// 双向链表
	hashMap []list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	MyHashMap := MyHashSet{769, make([]list.List, 769)}
	return MyHashMap
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		k := key % this.base
		this.hashMap[k].PushBack(key)
	}
}

func (this *MyHashSet) Remove(key int) {
	k := key % this.base
	for e := this.hashMap[k].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			this.hashMap[k].Remove(e)
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	k := key % this.base
	for e := this.hashMap[k].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end

func main() {
	ops := []string{"add", "add", "contains", "contains", "add", "contains", "remove", "contains"}
	nums := []int{1, 2, 1, 3, 2, 2, 2, 2}
	myhashmap := Constructor()
	for i, f := range ops {
		switch f {
		case "add":
			myhashmap.Add(nums[i])
		case "contains":
			fmt.Println(myhashmap.Contains(nums[i]))
		case "remove":
			myhashmap.Remove(nums[i])
		}
	}

}
