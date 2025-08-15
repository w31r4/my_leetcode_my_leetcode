/*
 * @lc app=leetcode.cn id=48 lang=golang
 * @lcpr version=30204
 *
 * [48] 旋转图像
 *
 * https://leetcode.cn/problems/rotate-image/description/
 *
 * algorithms
 * Medium (78.09%)
 * Likes:    2049
 * Dislikes: 0
 * Total Accepted:    753.6K
 * Total Submissions: 961.5K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
 *
 * 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
 *
 *
 *
 * 示例 1：
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[[7,4,1],[8,5,2],[9,6,3]]
 *
 *
 * 示例 2：
 *
 * 输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
 * 输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == matrix.length == matrix[i].length
 * 1 <= n <= 20
 * -1000 <= matrix[i][j] <= 1000
 *
 *
 *
 *
 */

// @lcpr-template-start
package leetcode

import "slices"

// @lcpr-template-end
// @lc code=start
// 思路：各数移动规律如下:+a +b /+b -a/-a -b/-b +a; a+b==n-1,a,b>=0
// 建立一个 next 变量存储下一个要更换位置的值
// 每一层 n-2，直到 n-1==0,break
func rotate(matrix [][]int) {
	// n := len(matrix)
	// //代表当前内循环的单行长度
	// roundNum := 0
	// //循环轮次
	// //同时代表当前内循环指向元素的行数
	// for ; n-1 > 0; n -= 2 {
	// 	y := n - 1
	// 	//初始最大偏移量
	// 	x := 0
	// 	//初始偏移量为 0

	// 	for i := roundNum; i < n-1+roundNum; i++ {
	// 		//i 代表当前指向元素的列数
	// 		//这里的 n-1 是因为第 n-1 位已经在第一位旋转过了
	// 		//最开始设计的时候是把 i 当作循环次数来的，从 0 开始递增，当循环次数等于 n 时终止循环，但后来发现还需要一个当前元素列数，就把 i 初始化为 roundNum 了，但是后面的终止条件没有给 n-1 加上 roundNum，前后设计不一致，导致了很多逻辑错位
	// 		//收获的警示是：不要轻易修改最初设计时赋予的抽象概念，新的抽象概念交给新的变量，防止逻辑错位

	// 		next := 0
	// 		x_loc := roundNum + x
	// 		y_loc := i + y
	// 		matrix[x_loc][y_loc], next = matrix[roundNum][i], matrix[x_loc][y_loc]
	// 		x_loc = x_loc + y
	// 		y_loc = y_loc - x
	// 		matrix[x_loc][y_loc], next = next, matrix[x_loc][y_loc]
	// 		x_loc -= x
	// 		y_loc -= y
	// 		matrix[x_loc][y_loc], next = next, matrix[x_loc][y_loc]
	// 		x_loc -= y
	// 		y_loc += x
	// 		matrix[x_loc][y_loc] = next
	// 		x++
	// 		y--
	// 	}
	// 	roundNum++
	// }
	//先观察结果和当前数组的数学关系
	//两次翻转等于一次旋转
	//一次转置 + 一次行翻转==旋转 90 度
	//所以位于 i 行 j 列的元素，去到 j 行 n−1−i 列，即 (i,j)→(j,n−1−i)。
	//(i,j)→(j,n−1−i) 可以通过两次翻转操作得到：

	// (i,j)
	// 转置：
	//  (j,i)
	// 行翻转：
	//  (j,n−1−i)
	// n := len(matrix)
	// 第一步：转置
	// for i := range n {
	// 	for j := range i { // 遍历对角线下方元素
	// 		matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
	// 	}
	// }

	// // 第二步：行翻转
	// for _, row := range matrix {
	// 	slices.Reverse(row)
	// }
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]\n
// @lcpr case=end

*/
