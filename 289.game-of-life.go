/*
 * @lc app=leetcode.cn id=289 lang=golang
 * @lcpr version=30204
 *
 * [289] 生命游戏
 *
 * https://leetcode.cn/problems/game-of-life/description/
 *
 * algorithms
 * Medium (77.15%)
 * Likes:    672
 * Dislikes: 0
 * Total Accepted:    146.2K
 * Total Submissions: 189.4K
 * Testcase Example:  '[[0,1,0],[0,0,1],[1,1,1],[0,0,0]]'
 *
 * 根据 百度百科 ， 生命游戏 ，简称为 生命，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。
 *
 * 给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态：1 即为 活细胞（live），或 0 即为
 * 死细胞（dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：
 *
 *
 * 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
 * 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
 * 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
 * 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
 *
 *
 * 下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是 同时 发生的。给你 m x n 网格面板 board
 * 的当前状态，返回下一个状态。
 *
 * 给定当前 board 的状态，更新 board 到下一个状态。
 *
 * 注意 你不需要返回任何东西。
 *
 *
 *
 * 示例 1：
 *
 * 输入：board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
 * 输出：[[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
 *
 *
 * 示例 2：
 *
 * 输入：board = [[1,1],[1,0]]
 * 输出：[[1,1],[1,1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == board.length
 * n == board[i].length
 * 1 <= m, n <= 25
 * board[i][j] 为 0 或 1
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你可以使用原地算法解决本题吗？请注意，面板上所有格子需要同时被更新：你不能先更新某些格子，然后使用它们的更新后的值再更新其他格子。
 * 本题中，我们使用二维数组来表示面板。原则上，面板是无限的，但当活细胞侵占了面板边界时会造成问题。你将如何解决这些问题？
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 思路：简单题简单做，维护一个 m*n 的数组差不多得了，将经过生存规则判定后的结果保存在新的数组里
// func isOutOfRange(rowLen int, colLen int, nowRow int, nowCol int) bool {
// 	// 检查行是否超出范围
// 	if nowRow < 0 || nowRow >= rowLen {
// 		return true
// 	}
// 	// 检查列是否超出范围
// 	if nowCol < 0 || nowCol >= colLen {
// 		return true
// 	}
// 	// 如果行和列都在范围内，则不超出范围
// 	return false
// }
// func gameOfLife(board [][]int) {
// 	cols := len(board[0])
// 	//列数
// 	rows := len(board)
// 	//行数
// 	// nextRows := make([]int, cols)
// 	nextWorld := make([][]int, rows)
// 	for i := range nextWorld {
// 		nextWorld[i] = make([]int, cols)
// 	}
// 	for i := 0; i < rows; i++ {
// 		//第几行
// 		for j := 0; j < cols; j++ {
// 			//第几列
// 			nextWorld[i][j] = board[i][j]
// 			survivalCell := 0
// 			lastRow := i - 1
// 			lastCol := j - 1
// 			nextRow := i + 1
// 			nextCol := j + 1
// 			if !isOutOfRange(rows, cols, lastRow, lastCol) {
// 				if board[lastRow][lastCol] == 1 {
// 					survivalCell++
// 				}
// 			}
// 			if !isOutOfRange(rows, cols, lastRow, j) {
// 				if board[lastRow][j] == 1 {
// 					survivalCell++
// 				}
// 			}
// 			if !isOutOfRange(rows, cols, lastRow, nextCol) {
// 				if board[lastRow][nextCol] == 1 {
// 					survivalCell++
// 				}
// 			}
// 			if !isOutOfRange(rows, cols, i, lastCol) {
// 				if board[i][lastCol] == 1 {
// 					survivalCell++
// 				}
// 			}
// 			if !isOutOfRange(rows, cols, i, nextCol) {
// 				if board[i][nextCol] == 1 {
// 					survivalCell++
// 				}
// 			}
// 			if !isOutOfRange(rows, cols, nextRow, lastCol) {
// 				if board[nextRow][lastCol] == 1 {
// 					survivalCell++
// 				}
// 			}
// 			if !isOutOfRange(rows, cols, nextRow, j) {
// 				if board[nextRow][j] == 1 {
// 					survivalCell++
// 				}
// 			}
// 			if !isOutOfRange(rows, cols, nextRow, nextCol) {
// 				if board[nextRow][nextCol] == 1 {
// 					survivalCell++
// 				}
// 			}
// 			if nextWorld[i][j] == 1 {
// 				if survivalCell < 2 || survivalCell > 3 {
// 					nextWorld[i][j] = 0
// 				}
// 			} else {
// 				if survivalCell == 3 {
// 					nextWorld[i][j] = 1
// 				}
// 			}
// 		}
// 	}
// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < cols; j++ {
// 			board[i][j] = nextWorld[i][j]
// 		}
// 	}

// }
// 使用 2 和 3 表示从活到死和从死到活
func isOutOfRange(rowLen int, colLen int, nowRow int, nowCol int) bool {
	// ... (不变)
	if nowRow < 0 || nowRow >= rowLen {
		return true
	}
	if nowCol < 0 || nowCol >= colLen {
		return true
	}
	return false
}

func gameOfLife(board [][]int) { // 原地修改，无返回值
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	rows := len(board)
	cols := len(board[0])

	// 定义邻居的相对偏移量
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	// 第一遍：计算并标记状态变化
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			liveNeighbors := 0
			// 检查所有 8 个邻居
			for k := 0; k < 8; k++ {
				ni, nj := i+dx[k], j+dy[k]
				// 检查邻居的 *原始* 状态 (1 或 2 都表示原始是活的)
				if !isOutOfRange(rows, cols, ni, nj) && (board[ni][nj] == 1 || board[ni][nj] == 2) {
					liveNeighbors++
				}
			}

			// 应用规则，使用 2 和 3 标记变化
			currentState := board[i][j]
			if currentState == 1 { // 当前是活的
				if liveNeighbors < 2 || liveNeighbors > 3 {
					board[i][j] = 2 // 标记为 活 -> 死
				}
				// 如果是 liveNeighbors == 2 || liveNeighbors == 3，保持 1 (活 -> 活)
			} else { // 当前是死的 (currentState == 0)
				if liveNeighbors == 3 {
					board[i][j] = 3 // 标记为 死 -> 活
				}
				// 否则保持 0 (死 -> 死)
			}
		}
	}

	// 第二遍：应用标记，恢复到 0 和 1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 2 { // 活 -> 死
				board[i][j] = 0
			} else if board[i][j] == 3 { // 死 -> 活
				board[i][j] = 1
			}
			// 0 和 1 保持不变
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1],[1,0]]\n
// @lcpr case=end

*/
