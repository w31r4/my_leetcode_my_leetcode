/*
 * @lc app=leetcode.cn id=383 lang=golang
 * @lcpr version=30204
 *
 * [383] 赎金信
 *
 * https://leetcode.cn/problems/ransom-note/description/
 *
 * algorithms
 * Easy (66.77%)
 * Likes:    968
 * Dislikes: 0
 * Total Accepted:    622.4K
 * Total Submissions: 928.4K
 * Testcase Example:  '"a"\n"b"'
 *
 * 给你两个字符串：ransomNote 和 magazine，判断 ransomNote 能不能由 magazine 里面的字符构成。
 *
 * 如果可以，返回 true；否则返回 false。
 *
 * magazine 中的每个字符只能在 ransomNote 中使用一次。
 *
 *
 *
 * 示例 1：
 *
 * 输入：ransomNote = "a", magazine = "b"
 * 输出：false
 *
 *
 * 示例 2：
 *
 * 输入：ransomNote = "aa", magazine = "ab"
 * 输出：false
 *
 *
 * 示例 3：
 *
 * 输入：ransomNote = "aa", magazine = "aab"
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= ransomNote.length, magazine.length <= 10^5
 * ransomNote 和 magazine 由小写英文字母组成
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 简单题简单做，建表查表即可
func canConstruct(ransomNote string, magazine string) bool {
	hashMagazine := [128]int{}
	for _, val := range magazine {
		hashMagazine[val]++
	}
	for _, val := range ransomNote {
		if hashMagazine[val] == 0 {
			return false
		}
		hashMagazine[val]--
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "a"\n"b"\n
// @lcpr case=end

// @lcpr case=start
// "aa"\n"ab"\n
// @lcpr case=end

// @lcpr case=start
// "aa"\n"aab"\n
// @lcpr case=end

*/
