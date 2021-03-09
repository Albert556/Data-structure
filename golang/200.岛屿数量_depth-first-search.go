/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package main

import "fmt"

// @lc code=start
// dfs
// func numIslands(grid [][]byte) int {
// 	m := len(grid)
// 	n := len(grid[0])
// 	if m == 0 {
// 		return 0
// 	}
// 	var count int

// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if grid[i][j] == '1' {
// 				count++
// 				dfs(grid, i, j, m, n)
// 			}
// 		}
// 	}

// 	return count
// }

// func dfs(g [][]byte, x, y, r, c int) {
// 	if x < 0 || y < 0 || x >= r || y >= c || g[x][y] == '0' {
// 		return
// 	}
// 	g[x][y] = '0'
// 	dfs(g, x+1, y, r, c)
// 	dfs(g, x-1, y, r, c)
// 	dfs(g, x, y+1, r, c)
// 	dfs(g, x, y-1, r, c)
// }

//bfs
// func numIslands(grid [][]byte) int {
// 	rows := len(grid)
// 	cols := len(grid[0])
// 	if rows == 0 {
// 		return 0
// 	}
// 	count := 0
// 	quene := [][]int{}

// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < cols; j++ {
// 			if grid[i][j] == byte(1) {
// 				count++
// 				loc := []int{i, j}
// 				quene = append(quene, loc)
// 				grid[i][j] = byte(0)
// 			}
// 			for len(quene) > 0 {
// 				for size := len(quene); size > 0; size-- {
// 					q := quene[0]
// 					quene = quene[1:]
// 					x := q[0]
// 					y := q[1]
// 					if x-1 >= 0 && grid[x-1][y] == byte(1) {
// 						grid[x-1][y] = byte(0)
// 						loc := []int{x - 1, y}
// 						quene = append(quene, loc)
// 					}
// 					if x+1 < rows && grid[x+1][y] == byte(1) {
// 						grid[x+1][y] = byte(0)
// 						loc := []int{x + 1, y}
// 						quene = append(quene, loc)
// 					}
// 					if y-1 >= 0 && grid[x][y-1] == byte(1) {
// 						grid[x][y-1] = byte(0)
// 						loc := []int{x, y - 1}
// 						quene = append(quene, loc)
// 					}
// 					if y+1 < cols && grid[x][y+1] == byte(1) {
// 						grid[x][y+1] = byte(0)
// 						loc := []int{x, y + 1}
// 						quene = append(quene, loc)
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return count
// }

// Union Find
type unionSet struct {
	root  []int
	count int
}

func NewUnionFind(grid [][]byte) *unionSet {
	rows := len(grid)
	cols := len(grid[0])
	count := rows * cols
	root := make([]int, count)
	for i := 0; i < count; i++ {
		root[i] = i
	}

	return &unionSet{root: root, count: count}
}

func (u *unionSet) find(x int) int {
	if u.root[x] == x {
		return x
	} else {
		root := u.find(u.root[x])
		u.root[x] = root
		return u.root[x]
	}
}

func (u *unionSet) union(x, y int) {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX != rootY {
		u.root[rootX] = rootY
		u.count--
	}
}

func (u *unionSet) getCount() int {
	return u.count
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])
	water := 0
	dic := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	u := NewUnionFind(grid)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '0' {
				water++
			} else {
				for _, d := range dic {
					x := i + d[0]
					y := j + d[1]
					if x >= 0 && x < rows && y >= 0 && y < cols && grid[x][y] == '1' {
						u.union(x*cols+y, i*cols+j)
					}
				}
			}
		}
	}
	fmt.Println(u.root)
	return u.getCount() - water
}

// @lc code=end

func main() {
	a := [][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	fmt.Println(numIslands(a))
}
