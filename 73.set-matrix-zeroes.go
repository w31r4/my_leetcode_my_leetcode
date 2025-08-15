/*
 * @lc app=leetcode.cn id=73 lang=golang
 * @lcpr version=30204
 *
 * [73] 矩阵置零
 *
 * https://leetcode.cn/problems/set-matrix-zeroes/description/
 *
 * algorithms
 * Medium (69.91%)
 * Likes:    1205
 * Dislikes: 0
 * Total Accepted:    546.2K
 * Total Submissions: 776.5K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * 给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用 原地 算法。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
 * 输出：[[1,0,1],[0,0,0],[1,0,1]]
 *
 *
 * 示例 2：
 *
 * 输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
 * 输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[0].length
 * 1 <= m, n <= 200
 * -2^31 <= matrix[i][j] <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
 * 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
 * 你能想出一个仅使用常量空间的解决方案吗？
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 简单题，空间 O(m*n) 解决
// 建立是否置零哈希表，然后通过查询哈希表判断是否置零
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return // 处理空矩阵或无效矩阵
	}
	rowLen := len(matrix)
	// zeroRow := make([]int, rowLen)
	// colLen := len(matrix[0])
	// zeroCol := make([]int, colLen)

	// for i := 0; i < rowLen; i++ {
	// 	for j := 0; j < colLen; j++ {
	// 		if matrix[i][j] == 0 {
	// 			zeroRow[i] = 1
	// 			zeroCol[j] = 1
	// 		}
	// 	}
	// }
	// for i := 0; i < rowLen; i++ {
	// 	for j := 0; j < colLen; j++ {
	// 		if zeroRow[i] == 1 || zeroCol[j] == 1 {
	// 			matrix[i][j] = 0
	// 		}
	// 	}
	// }
	//O(1) 解法将用于判断的哈希表放于第一行和第一列中，当某行需要置零时，就将对应的第一列的对应行的数值置零，当某列需要置零时，将第一行对应列的数值置零
	//当现在有一个问题，那就是第一行第一列的原始信息被修改了，我们无法得知第一行第一列是否需要置零，很简单，将第一行第一列最先遍历判断是否需要将第一行第一列进行置零，但我们不能用 [0,0] 号元素来，因为 [0][0] 一旦置零第一行第一列都要置零，需要建立两个单独的 bool 变量保存第一行第一列的置零状态
	//而且如果某行某列出现了 0，那么第一行的某列和第一列的某行都是要置零的，所以可行，把结果提前反映在第一行第一列中
	//记得从第一行第一列之外开始遍历是否需要置零以及置零，因为过早遍历第一行会导致剩下的行列是否需要置零的信息丢失，如果 [0][0] 号位置是 0，那么第一行在置零遍历中会全部置零，导致二维数组最终全部置零
	colLen := len(matrix[0])
	isFirstRowEmpty, isFirstColEmpty := false, false
	for i := 0; i < rowLen; i++ {
		if matrix[i][0] == 0 {
			isFirstColEmpty = true
			break
		}
	}
	for j := 0; j < colLen; j++ {
		if matrix[0][j] == 0 {
			isFirstRowEmpty = true
			break
		}
	}
	for i := 1; i < rowLen; i++ {
		for j := 1; j < colLen; j++ {

			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < rowLen; i++ {
		for j := 1; j < colLen; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if isFirstColEmpty {
		for i := 0; i < rowLen; i++ {
			matrix[i][0] = 0
		}
	}
	if isFirstRowEmpty {
		for j := 0; j < colLen; j++ {
			matrix[0][j] = 0
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [[1,1,1],[1,0,1],[1,1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1,2,0],[3,4,5,2],[1,3,1,5]]\n
// @lcpr case=end

*/
