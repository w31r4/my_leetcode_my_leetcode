/*
 * @lc app=leetcode.cn id=258 lang=golang
 * @lcpr version=30204
 *
 * [258] 各位相加
 *
 * https://leetcode.cn/problems/add-digits/description/
 *
 * algorithms
 * Easy (70.20%)
 * Likes:    721
 * Dislikes: 0
 * Total Accepted:    226.8K
 * Total Submissions: 323K
 * Testcase Example:  '38'
 *
 * 给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。返回这个结果。
 *
 *
 *
 * 示例 1:
 *
 * 输入：num = 38
 * 输出：2
 * 解释：各位相加的过程为：
 * 38 --> 3 + 8 --> 11
 * 11 --> 1 + 1 --> 2
 * 由于 2 是一位数，所以返回 2。
 *
 *
 * 示例 2:
 *
 * 输入：num = 0
 * 输出：0
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= num <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：你可以不使用循环或者递归，在 O(1) 时间复杂度内解决这个问题吗？
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
	//模九是因为 :200-2=99*2,200%9=2,114514 1+1+4+5+1 +4=16，这是个不断减去 9 的过程，每一步都在减去 9 的倍数，所以可以使用 9 去取模获得结果
	//其他数也是一样的，如果有个函数是在不断地减去 8 直到结果小于 8，那么直接模 8 就能获得结果
}

// @lc code=end

/*
// @lcpr case=start
// 38\n
// @lcpr case=end

// @lcpr case=start
// 0\n
// @lcpr case=end

*/
