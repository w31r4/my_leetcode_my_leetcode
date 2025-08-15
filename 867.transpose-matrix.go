/*
 * @lc app=leetcode.cn id=867 lang=golang
 * @lcpr version=30204
 *
 * [867] 转置矩阵
 *
 * https://leetcode.cn/problems/transpose-matrix/description/
 *
 * algorithms
 * Easy (68.57%)
 * Likes:    281
 * Dislikes: 0
 * Total Accepted:    126.5K
 * Total Submissions: 184.5K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵。
 *
 * 矩阵的 转置 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[[1,4,7],[2,5,8],[3,6,9]]
 *
 *
 * 示例 2：
 *
 * 输入：matrix = [[1,2,3],[4,5,6]]
 * 输出：[[1,4],[2,5],[3,6]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 1000
 * 1 <= m * n <= 10^5
 * -10^9 <= matrix[i][j] <= 10^9
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}
	lins := len(matrix)
	cols := len(matrix[0])
	ans := make([][]int, cols)
	for i := range ans {
		ans[i] = make([]int, lins)
	}
	for i := 0; i < cols; i++ {
		for n := 0; n < lins; n++ {
			ans[i][n] = matrix[n][i]
		}
	}
	return ans
}

// 抽象思维，不用去思考先算行还是先换列，只要写出行列互换的代码，并可以让其正确迭代就行

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6]]\n
// @lcpr case=end

*/
