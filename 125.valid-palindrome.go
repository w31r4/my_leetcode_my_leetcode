/*
 * @lc app=leetcode.cn id=125 lang=golang
 * @lcpr version=30204
 *
 * [125] 验证回文串
 *
 * https://leetcode.cn/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (48.05%)
 * Likes:    809
 * Dislikes: 0
 * Total Accepted:    711.3K
 * Total Submissions: 1.5M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串。
 *
 * 字母和数字都属于字母数字字符。
 *
 * 给你一个字符串 s，如果它是 回文串，返回 true；否则，返回 false。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "A man, a plan, a canal: Panama"
 * 输出：true
 * 解释："amanaplanacanalpanama" 是回文串。
 *
 *
 * 示例 2：
 *
 * 输入：s = "race a car"
 * 输出：false
 * 解释："raceacar" 不是回文串。
 *
 *
 * 示例 3：
 *
 * 输入：s = " "
 * 输出：true
 * 解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
 * 由于空字符串正着反着读都一样，所以是回文串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 2 * 10^5
 * s 仅由可打印的 ASCII 字符组成
 *
 *
 */

// @lcpr-template-start
package leetcode

import (
	"strings"
	"unicode"
)

// @lcpr-template-end
// @lc code=start
// 两步走：
// 第一步：将数据转化为全小写英文字符和数字
// 第二步：验证该全小写英文字符和数字字符组成的字符串是否为回文串
// 第一步思路：调库 X or 双指针处理
func isPalindrome(s string) bool {
	var builder strings.Builder
	for _, val := range s {
		if unicode.IsLower(val) {
			builder.WriteRune(val)
		} else if unicode.IsUpper(val) {
			builder.WriteRune(unicode.ToLower(val))
		} else if unicode.IsDigit(val) {
			builder.WriteRune(val)
		}
	}
	//先构建一个只包含小写字母和数字的字符串
	s = builder.String()
	if s == "" {
		return true
	}
	//空字符串自动返回 true

	length := len(s) - 1

	for i := 0; i <= length/2; i++ {
		if s[i] == s[length-i] {
			continue
		} else {
			return false
		}
		//相向双指针
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "123123123p123123123123"\n
// @lcpr case=end

// @lcpr case=start
// "race a car"\n
// @lcpr case=end

// @lcpr case=start
// " "\n
// @lcpr case=end

*/
