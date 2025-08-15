/*
 * @lc app=leetcode.cn id=326 lang=golang
 * @lcpr version=30204
 *
 * [326] 3 的幂
 *
 * https://leetcode.cn/problems/power-of-three/description/
 *
 * algorithms
 * Easy (52.35%)
 * Likes:    353
 * Dislikes: 0
 * Total Accepted:    270.7K
 * Total Submissions: 517.1K
 * Testcase Example:  '27'
 *
 * 给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true；否则，返回 false。
 *
 * 整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3^x
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 27
 * 输出：true
 *
 *
 * 示例 2：
 *
 * 输入：n = 0
 * 输出：false
 *
 *
 * 示例 3：
 *
 * 输入：n = 9
 * 输出：true
 *
 *
 * 示例 4：
 *
 * 输入：n = 45
 * 输出：false
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
 *
 *
 * 进阶：你能不使用循环或者递归来完成本题吗？
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func isPowerOfThree(n int) bool {
	return n == 3
	//使用 3 在范围内的最大次方取模，看是否余为 0
	//质数的 n 次幂一定是质数的较小次幂*质数的较小次幂
}

// @lc code=end

/*
// @lcpr case=start
// 27\n
// @lcpr case=end

// @lcpr case=start
// 0\n
// @lcpr case=end

// @lcpr case=start
// 9\n
// @lcpr case=end

// @lcpr case=start
// 45\n
// @lcpr case=end

*/
