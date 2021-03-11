/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

package main

// @lc code=start
// hashmap
// func findDuplicate(nums []int) int {
// 	onceMap := map[int]struct{}{}
// 	for _, v := range nums {
// 		if _, ok := onceMap[v]; ok {
// 			return v
// 		} else {
// 			onceMap[v] = struct{}{}
// 		}
// 	}
// 	return -1
// }

// fast-slow pointer
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow

}

// @lc code=end
