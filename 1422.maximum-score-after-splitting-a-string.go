/*
 * @lc app=leetcode.cn id=1422 lang=golang
 * @lcpr version=30204
 *
 * [1422] 分割字符串的最大得分
 *
 * https://leetcode.cn/problems/maximum-score-after-splitting-a-string/description/
 *
 * algorithms
 * Easy (57.89%)
 * Likes:    153
 * Dislikes: 0
 * Total Accepted:    67.5K
 * Total Submissions: 116.5K
 * Testcase Example:  '"011101"'
 *
 * 给你一个由若干 0 和 1 组成的字符串 s，请你计算并返回将该字符串分割成两个 非空 子字符串（即 左 子字符串和 右
 * 子字符串）所能获得的最大得分。
 *
 * 「分割字符串的得分」为 左 子字符串中 0 的数量加上 右 子字符串中 1 的数量。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "011101"
 * 输出：5
 * 解释：
 * 将字符串 s 划分为两个非空子字符串的可行方案有：
 * 左子字符串 = "0" 且 右子字符串 = "11101"，得分 = 1 + 4 = 5
 * 左子字符串 = "01" 且 右子字符串 = "1101"，得分 = 1 + 3 = 4
 * 左子字符串 = "011" 且 右子字符串 = "101"，得分 = 1 + 2 = 3
 * 左子字符串 = "0111" 且 右子字符串 = "01"，得分 = 1 + 1 = 2
 * 左子字符串 = "01110" 且 右子字符串 = "1"，得分 = 2 + 1 = 3
 *
 *
 * 示例 2：
 *
 * 输入：s = "00111"
 * 输出：5
 * 解释：当 左子字符串 = "00" 且 右子字符串 = "111" 时，我们得到最大得分 = 2 + 3 = 5
 *
 *
 * 示例 3：
 *
 * 输入：s = "1111"
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= s.length <= 500
 * 字符串 s 仅由字符 '0' 和 '1' 组成。
 *
 *
 */

// @lcpr-template-start
package leetcode

import "strings"

// @lcpr-template-end
// @lc code=start
func maxScore(s string) int {
	//初解：没一点思路，看着题目发呆
	//思路如下：统计整个字符串中 1 的数量，同时我们也得到了 0 的数
	score := strings.Count(s, "1")
	//将 1 的数量记为 score
	maxScore := 0
	// 设置 maxScore，用于记录 score 最大值
	for _, n := range s[:len(s)-1] {
		// len(s)-1 是因为双方两组均不为空，即无法统计到最后一位，当为 0000 时，输出为 3 而非 4
		if n == '0' {
			score += 1
		} else {
			score -= 1
		}
		//当字符串从左往右开始分组，每当出现一次 0 就加一分，出现一次 1 就减一分，因为我们已经统计出来了 1 的个数，那么右边组队最大值我们已经获取到了，目前只需要确定左边组的 0 是否比 1 多
		if maxScore < score {
			maxScore = score
		}
	}
	return maxScore
}

// @lc code=end

/*
// @lcpr case=start
// "0000"\n
// @lcpr case=end

// @lcpr case=start
// "00111"\n
// @lcpr case=end

// @lcpr case=start
// "1111"\n
// @lcpr case=end

*/
