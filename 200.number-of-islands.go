/*
 * @lc app=leetcode.cn id=200 lang=golang
 * @lcpr version=30204
 *
 * [200] 岛屿数量
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 使用 BFS 或者 DFS 把遇到的每一个 1 遍历，并修改为 0
// 外部给遇到的岛屿数量加一
func numIslands(grid [][]byte) (ans int) {
	lenOfCol := len(grid[0]) - 1
	lenOfRow := len(grid) - 1
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r > lenOfRow || r < 0 {
			return
		}
		if c > lenOfCol || c < 0 {
			return
		}
		if grid[r][c] == '0' {
			return
		}
		grid[r][c] = '0'
		//记录上下左右状态并全部进行标记
		dfs(r+1, c)
		//有可能 U 型岛屿，需要向上看
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
		//上岛上下左右看
		return
	}
	for r, row := range grid {
		for c, val := range row {
			if val == '1' {
				ans++
				dfs(r, c)
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]\n
// @lcpr case=end

*/
