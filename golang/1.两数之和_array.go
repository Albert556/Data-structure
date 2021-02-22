/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
// func twoSum(nums []int, target int) []int {
// 	for i, num1 := range nums {
// 		for j := i + 1; j < len(nums); j++ {
// 			if target == (num1 + nums[j]) {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, x := range nums {
		if j, ok := hashmap[target-x]; ok {
			return []int{j, i}
		}
		hashmap[x] = i
	}
	return nil
}

// @lc code=end

