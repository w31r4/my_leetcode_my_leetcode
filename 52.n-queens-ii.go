/*
 * @lc app=leetcode.cn id=52 lang=golang
 * @lcpr version=30204
 *
 * [52] N 皇后 II
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func operateQueen(i, j, n int, board [][]int, isPut bool) {
	num := 0
	if isPut {
		num = i + 1
	}
	for index := 0; index < n; index++ {
		if board[i][index] == 0 || board[i][index] == i+1 {
			board[i][index] = num
		}
	}
	for index := 0; index < n; index++ {
		if board[index][j] == 0 || board[index][j] == i+1 {
			board[index][j] = num
		}
	}
	index1 := i
	index2 := j
	for index1 < n && index2 < n {
		if board[index1][index2] == 0 || board[index1][index2] == i+1 {
			board[index1][index2] = num
		}
		index1++
		index2++
	}
	index1 = i
	index2 = j
	for index1 >= 0 && index2 >= 0 {
		if board[index1][index2] == 0 || board[index1][index2] == i+1 {
			board[index1][index2] = num
		}
		index1--
		index2--
	}
	index1 = i
	index2 = j
	for index1 >= 0 && index2 < n {
		if board[index1][index2] == 0 || board[index1][index2] == i+1 {
			board[index1][index2] = num
		}
		index1--
		index2++
	}
	index1 = i
	index2 = j
	for index1 < n && index2 >= 0 {
		if board[index1][index2] == 0 || board[index1][index2] == i+1 {
			board[index1][index2] = num
		}
		index1++
		index2--
	}

}

func totalNQueens(n int) (answer int) {
	//N 皇后的回溯
	//建立一个棋盘，当放上一个皇后后，皇后及其可以攻击到的区域都赋值为 1

	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}

	var dfs func(level int)
	dfs = func(level int) {
		if level == n {
			answer++
			return
		}
		// temp := make([][]int, n)
		// for i := 0; i < n; i++ {
		// 	temp[i] = make([]int, n)
		// }
		//存在浅拷贝的情况
		for i := 0; i < n; i++ {
			//第 level 行，第 i 列
			if board[level][i] == 0 {
				// copy(temp, board)
				operateQueen(level, i, n, board, true)
				dfs(level + 1)
				// copy(board, temp)
				operateQueen(level, i, n, board, false)
			}
		}
		return
	}
	dfs(0)
	return
}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
