/*
 * @lc app=leetcode.cn id=130 lang=golang
 * @lcpr version=30204
 *
 * [130] 被围绕的区域
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
type Node1 struct {
	Val       int
	Neighbors []*Node2
}

// 后序遍历，当遍历完所有的 o 还没有遇到边界时，将所有的 o 转化为 x
// 但是也有一个问题：重复遍历，当一个 O 被还原后还会被重复的找到然后遍历
// 所以我们要换个思路，不要被之前的想法局限
// 题目要求被包围的区域，那么边界中所有存在的 O 都遍历找到所有相连区域，并标记上 S(safe),然后把所有非 safe 的 O 转变为 X
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 { // <--- 加上这个检查！
		return
	}
	rLength := len(board) - 1
	cLength := len(board[0]) - 1
	var dfs func(r, c int)

	dfs = func(r, c int) {
		//代码错误在没有处理重复访问问题
		//会导致一块地点多次被访问导致内存溢出
		//代码为所有和边界上 O 相连的 O 打上 s 的标记

		if r > rLength || r < 0 {
			// isTouchEdge = true
			return
		}
		if c > cLength || c < 0 {
			// isTouchEdge = true
			return
		}
		if board[r][c] == 'X' || board[r][c] == 's' {
			return
		}
		board[r][c] = 's'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)

		return
	}
	//找到四条边并遍历
	for i, val := range board[0] {
		if val == 'O' {
			dfs(0, i)
		}
	}
	for i, val := range board[rLength] {
		if val == 'O' {
			dfs(rLength, i)
		}
	}
	for i := 0; i <= rLength; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][cLength] == 'O' {
			dfs(i, cLength)
		}
	}
	for i, row := range board {
		for j, val := range row {
			if val == 'O' {
				board[i][j] = 'X'
			}
			if val == 's' {
				board[i][j] = 'O'
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]\n
// @lcpr case=end

// @lcpr case=start
// [["X"]]\n
// @lcpr case=end

*/
