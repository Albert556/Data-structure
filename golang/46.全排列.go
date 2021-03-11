/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

package main

import "fmt"

// visited为什么不用初始化

// @lc code=start
func permute(nums []int) [][]int {
	result := [][]int{}
	visited := map[int]bool{}
	// for _, v := range nums {
	// 	visited[v] = false
	// }
	var backtracking func(list []int)
	backtracking = func(list []int) {
		if len(list) == len(nums) {
			l := make([]int, len(list))
			copy(l, list)
			// fmt.Println(l)
			result = append(result, l)
			// fmt.Println(result)
			return
		}
		for _, num := range nums {
			if !visited[num] {
				list = append(list, num)
				visited[num] = true
				backtracking(list)
				list = list[:len(list)-1]
				visited[num] = false
			}
		}
	}

	backtracking([]int{})
	return result

}

// @lc code=end

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
