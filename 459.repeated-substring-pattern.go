/*
 * @lc app=leetcode.cn id=459 lang=golang
 * @lcpr version=30204
 *
 * [459] 重复的子字符串
 *
 * https://leetcode.cn/problems/repeated-substring-pattern/description/
 *
 * algorithms
 * Easy (52.16%)
 * Likes:    1280
 * Dislikes: 0
 * Total Accepted:    324.3K
 * Total Submissions: 621K
 * Testcase Example:  '"abab"'
 *
 * 给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。
 *
 *
 *
 * 示例 1:
 *
 * 输入：s = "abab"
 * 输出：true
 * 解释：可由子串 "ab" 重复两次构成。
 *
 *
 * 示例 2:
 *
 * 输入：s = "aba"
 * 输出：false
 *
 *
 * 示例 3:
 *
 * 输入：s = "abcabcabcabc"
 * 输出：true
 * 解释：可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 *
 * 1 <= s.length <= 10^4
 * s 由小写英文字母组成
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func repeatedSubstringPattern(s string) bool {
	j := 0
	n := len(s)
	if n == 0 {
		return false
	}
	prefix := make([]int, len(s))
	for i := 1; i < n; i++ {

		for j > 0 && s[i] != s[j] {
			j = prefix[j-1]
		}
		if s[j] == s[i] {
			j++
			prefix[i] = j
		}
		// if prefix[i] < prefix[i-1] {
		// 	return false
		// }
		//可能出现 aabbcaabbc 这样的，子串中仍有重复
	}
	if prefix[n-1] != 0 && n%(n-prefix[n-1]) == 0 {
		//对 n%(n-prefix[n-1])==0 的解释:n 可以整除字符串减去最大前后缀相同子串字符数剩余部分的字符数
		//该部分字符数一定是小于 n/2 的
		//为什么可以这么求？
		//因为当字符串为重复子串时，前缀表中的值在重复串第一子串过去后递增
		//最后前缀表最后一位的值为最大重复前后缀长度
		//如 123123123
		//   000123456
		return true
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// "abab"\n
// @lcpr case=end

// @lcpr case=start
// "aba"\n
// @lcpr case=end

// @lcpr case=start
// "abcabcabcabc"\n
// @lcpr case=end

*/
