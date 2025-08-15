/*
 * @lc app=leetcode.cn id=392 lang=golang
 * @lcpr version=30204
 *
 * [392] 判断子序列
 *
 * https://leetcode.cn/problems/is-subsequence/description/
 *
 * algorithms
 * Easy (52.86%)
 * Likes:    1160
 * Dislikes: 0
 * Total Accepted:    557.9K
 * Total Submissions: 1.1M
 * Testcase Example:  '"abc"\n"ahbgdc"'
 *
 * 给定字符串 s 和 t，判断 s 是否为 t 的子序列。
 *
 *
 * 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
 *
 * 进阶：
 *
 * 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10 亿，你需要依次检查它们是否为 T
 * 的子序列。在这种情况下，你会怎样改变代码？
 *
 * 致谢：
 *
 * 特别感谢 @pbrother 添加此问题并且创建所有测试用例。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "abc", t = "ahbgdc"
 * 输出：true
 *
 *
 * 示例 2：
 *
 * 输入：s = "axc", t = "ahbgdc"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 100
 * 0 <= t.length <= 10^4
 * 两个字符串都只由小写字符组成。
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func isSubsequence(s string, t string) bool {
	// //KMP 算法实现，狗狗的屎
	// prefix := make([]int, len(s))
	// //先建立前缀表
	// j := 0
	// for i := 1; i < len(s); i++ {
	// 	for s[i] != s[j] && j > 0 {
	// 		j = prefix[j-1]
	// 		//如果当前 j 指向的前缀末位值不等于 i 指向的后缀末为值，将 j 指向最近的和当前前缀的后缀相同的前相同的前缀的前缀的末位值
	// 		//直到 j==0 为止
	// 	}
	// 	//动态规划双指针
	// 	if s[i] == s[j] {
	// 		j++
	// 		prefix[i] = j
	// 		//为前缀表赋值
	// 	}
	// }
	// //构建完前缀表后，通过前缀表进行查询
	// j = 0
	// for i := 0; i < len(t); i++ {
	// 	if s[j] == t[i] {
	// 		j++
	// 	}
	// 	if s[j] != t[i] && j > 0 {
	// 		j = prefix[j-1]
	// 	}
	// 	if j == len(s)-1 {
	// 		return true
	// 	}
	// }
	// return false
	//狗屎，不是 tmd 的找完整子串，是找相对位置子串
	//ace 是 abcde 的子串
	j := 0
	if s == "" {
		return true
	}
	for _, val := range t {
		if s[j] == byte(val) {
			j++
		}
		if j == len(s) {
			return true
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// "abc"\n"ahbgdc"\n
// @lcpr case=end

// @lcpr case=start
// "axc"\n"ahbgdc"\n
// @lcpr case=end

*/
