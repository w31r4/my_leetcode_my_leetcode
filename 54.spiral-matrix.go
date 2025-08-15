/*
 * @lc app=leetcode.cn id=54 lang=golang
 * @lcpr version=30204
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode.cn/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (52.97%)
 * Likes:    1920
 * Dislikes: 0
 * Total Accepted:    731.4K
 * Total Submissions: 1.4M
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个 m 行 n 列的矩阵 matrix，请按照 顺时针螺旋顺序，返回矩阵中的所有元素。
 *
 *
 *
 * 示例 1：
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2：
 *
 * 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 10
 * -100 <= matrix[i][j] <= 100
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 思路，开局建立两个哈希表，哈希表的键分别为行列数，值为 true，代表行或列当前存在，按当前第一行，当前最后一列，当前最后一行，当前第一列的顺序遍历二维数组，将对应的行或列的哈希表的值调整为 false，从哈希值为 true 的地方开始继续按之前的顺序遍历，直到全为 false 为止
// 我的实际实现是既有标记又有剩余行数
func spiralOrder(matrix [][]int) (ans []int) {

	if len(matrix) == 1 {
		return matrix[0]
	}
	hashRow := make([]int, len(matrix))
	hashCol := make([]int, len(matrix[0]))
	//哈希标记
	ans = make([]int, len(matrix)*len(matrix[0]))
	//创建 ans 数组
	rowLen := len(hashRow)
	colLen := len(hashCol)
	//剩余行列数
	n := 0
	//ans 数组索引
	for i := 0; ; i++ {
		//处理当前首行
		for index, val := range matrix[0+i] {
			if hashCol[index] == 1 {
				continue
			}
			//跳过单行中已经处理过的列
			ans[n] = val
			n++
		}
		hashRow[i] = 1
		//当前行标记为处理过
		rowLen--
		//行数减一
		if rowLen == 0 {
			break
		}
		//处理当前最后一列
		for index, val := range hashRow {
			if val == 1 {
				continue
			}
			//跳过当前列中已经处理过的行
			ans[n] = matrix[index][len(hashCol)-1-i]
			n++
		}
		hashCol[len(hashCol)-1-i] = 1
		//当前列标记为处理过
		colLen--
		//列数减一
		if colLen == 0 {
			break
		}
		//逆序处理当前最后一行
		for index := 0; index < colLen; index++ {
			//从当前最后一行的最后未被处理过的一列开始
			//一位一位向前遍历
			//遍历当前剩余的列数
			ans[n] = matrix[len(hashRow)-1-i][len(hashCol)-2-i-index]
			n++
		}
		hashRow[len(hashRow)-1-i] = 1
		//当前行标记为处理过
		rowLen--
		//行数减一
		if rowLen == 0 {
			break
		}
		//逆序处理当前第一列
		for index := 0; index < rowLen; index++ {
			//从当前第一列的最后未被处理过的一行开始
			//一位一位向上遍历
			//遍历当前剩余的行数
			ans[n] = matrix[len(hashRow)-2-i-index][i]
			n++
		}

		hashCol[i] = 1
		//当前列标记为处理过
		colLen--
		//列数减一
		if colLen == 0 {
			break
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3,4],[5,6,7,8],[9,10,11,12]]\n
// @lcpr case=end

*/
