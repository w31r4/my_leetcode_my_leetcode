/*
 * @lc app=leetcode.cn id=1281 lang=golang
 * @lcpr version=30204
 *
 * [1281] 整数的各位积和之差
 *
 * https://leetcode.cn/problems/subtract-the-product-and-sum-of-digits-of-an-integer/description/
 *
 * algorithms
 * Easy (82.37%)
 * Likes:    178
 * Dislikes: 0
 * Total Accepted:    132.3K
 * Total Submissions: 160.6K
 * Testcase Example:  '234'
 *
 * 给你一个整数 n，请你帮忙计算并返回该整数「各位数字之积」与「各位数字之和」的差。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 234
 * 输出：15
 * 解释：
 * 各位数之积 = 2 * 3 * 4 = 24
 * 各位数之和 = 2 + 3 + 4 = 9
 * 结果 = 24 - 9 = 15
 *
 *
 * 示例 2：
 *
 * 输入：n = 4421
 * 输出：21
 * 解释：
 * 各位数之积 = 4 * 4 * 2 * 1 = 32
 * 各位数之和 = 4 + 4 + 2 + 1 = 11
 * 结果 = 32 - 11 = 21
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^5
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func productEach(n int) int {
	if n < 10 {
		return n
	}
	return productEach(n/10) * (n % 10)
}

// 边界条件是 n<10，非边界条件下将最后一位数字相加
func sumEach(n int) int {
	if n < 10 {
		return n
	}
	return sumEach(n/10) + n%10
}
func subtractProductAndSum(n int) int {
	return productEach(n) - sumEach(n)
}

// @lc code=end

/*
// @lcpr case=start
// 234\n
// @lcpr case=end

// @lcpr case=start
// 4421\n
// @lcpr case=end

*/
