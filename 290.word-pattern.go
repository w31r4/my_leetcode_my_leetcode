/*
 * @lc app=leetcode.cn id=290 lang=golang
 * @lcpr version=30204
 *
 * [290] 单词规律
 *
 * https://leetcode.cn/problems/word-pattern/description/
 *
 * algorithms
 * Easy (45.04%)
 * Likes:    702
 * Dislikes: 0
 * Total Accepted:    233.8K
 * Total Submissions: 517.9K
 * Testcase Example:  '"abba"\n"dog cat cat dog"'
 *
 * 给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。
 *
 * 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。
 *
 *
 *
 * 示例 1:
 *
 * 输入：pattern = "abba", s = "dog cat cat dog"
 * 输出：true
 *
 * 示例 2:
 *
 * 输入:pattern = "abba", s = "dog cat cat fish"
 * 输出：false
 *
 * 示例 3:
 *
 * 输入：pattern = "aaaa", s = "dog cat cat dog"
 * 输出：false
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= pattern.length <= 300
 * pattern 只包含小写英文字母
 * 1 <= s.length <= 3000
 * s 只包含小写英文字母和 ' '
 * s 不包含 任何前导或尾随对空格
 * s 中每个单词都被 单个空格 分隔
 *
 *
 */

// @lcpr-template-start
package leetcode

import "strings"

// @lcpr-template-end
// @lc code=start
// 思路：和上一题一样
// 1.当 pchar 已经匹配过但是当前单词和原来匹配到的单词不一样，返回 false
// 2.当 pchar 没有匹配到单词，但是单词已经配匹配过了返回 false
func wordPattern(pattern string, s string) bool {
	parts := strings.Fields(s)
	if len(pattern) != len(parts) {
		return false
	}
	pToSMap := [27]string{}
	sWordMapped := make(map[string]bool)
	for i := 0; i < len(parts); i++ {
		sWord := parts[i]
		pChar := pattern[i] - 'a' + 1
		pCharHasMapped := pToSMap[pChar] != ""

		if pCharHasMapped && sWord != pToSMap[pChar] {
			return false
		}
		sWordHasMapped := sWordMapped[sWord]
		if !pCharHasMapped && sWordHasMapped {
			return false
		}
		if !pCharHasMapped && !sWordHasMapped {
			sWordMapped[sWord] = true
			pToSMap[pChar] = sWord
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "abba"\n"dog cat cat dog"\n
// @lcpr case=end

// @lcpr case=start
// "abba"\n"dog cat cat fish"\n
// @lcpr case=end

// @lcpr case=start
// "aaaa"\n"dog cat cat dog"\n
// @lcpr case=end

*/
