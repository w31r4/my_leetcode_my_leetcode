/*
 * @lc app=leetcode.cn id=74 lang=golang
 * @lcpr version=30204
 *
 * [74] 搜索二维矩阵
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0])
	lastNumIndex := col - 1
	targetRow := -1
	for i := 0; i < row; i++ {
		// if matrix[i][0] >= target && matrix[i][lastNumIndex] <= target

		if target >= matrix[i][0] && matrix[i][lastNumIndex] >= target {
			//傻逼了，今天老是写反大小比较
			targetRow = i
			break
		}
	}
	if targetRow == -1 {
		return false
	}
	left := 0
	right := col
	for left < right {
		mid := left + (right-left)/2
		if matrix[targetRow][mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if matrix[targetRow][right] == target {
		return true
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3\n
// @lcpr case=end

// @lcpr case=start
// [[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n13\n
// @lcpr case=end

*/
