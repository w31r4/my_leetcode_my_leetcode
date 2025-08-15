/*
 * @lc app=leetcode.cn id=263 lang=golang
 * @lcpr version=30204
 *
 * [263] 丑数
 *
 * https://leetcode.cn/problems/ugly-number/description/
 *
 * algorithms
 * Easy (50.19%)
 * Likes:    477
 * Dislikes: 0
 * Total Accepted:    209.8K
 * Total Submissions: 418K
 * Testcase Example:  '6'
 *
 * 丑数 就是只包含质因数 2、3 和 5 的 正 整数。
 *
 * 给你一个整数 n，请你判断 n 是否为 丑数。如果是，返回 true；否则，返回 false。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 6
 * 输出：true
 * 解释：6 = 2 × 3
 *
 * 示例 2：
 *
 * 输入：n = 1
 * 输出：true
 * 解释：1 没有质因数。
 *
 * 示例 3：
 *
 * 输入：n = 14
 * 输出：false
 * 解释：14 不是丑数，因为它包含了另外一个质因数 7。
 *
 *
 *
 *
 * 提示：
 *
 *
 * -2^31 <= n <= 2^31 - 1
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func isUgly(n int) bool {
	// if n == 1 {
	// 	return true
	// }
	// if n%2 != 0 {
	// 	return false
	// }
	// for n%2 == 0 {

	// 	if n == 2 {
	// 		return true
	// 	}
	// 	n /= 2
	// }
	// if n%3 != 0 {
	// 	return false
	// }
	// for n%3 == 0 {
	// 	if n == 3 {
	// 		return true
	// 	}
	// 	n /= 3
	// }
	// if n%5 != 0 {
	// 	return false
	// }
	// for n%5 == 0 {
	// 	if n == 5 {
	// 		return true
	// 	}
	// 	n /= 5
	// }
	// return false
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n&(n-1) == 0

}

// 笔记 : 抽丝剥茧，把 3 和 5 都作为因子去除后，剩下的如果还是 2 的次方那么这个数字一定是由 2,3,5 构成的，这很合理
// 而且不用考虑精度丢失问题，因为如果不是 3 or 5 的倍数我们也不会对其进行除法
// @lc code=end

/*
// @lcpr case=start
// 600\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 14\n
// @lcpr case=end

*/
