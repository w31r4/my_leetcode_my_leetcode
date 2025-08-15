/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=30204
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (40.88%)
 * Likes:    10730
 * Dislikes: 0
 * Total Accepted:    3.4M
 * Total Submissions: 8.2M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串 s，请你找出其中不含有重复字符的 最长 子串 的长度。
 *
 *
 *
 * 示例 1:
 *
 * 输入：s = "abcabcbb"
 * 输出：3
 * 解释：因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入：s = "bbbbb"
 * 输出：1
 * 解释：因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 * 输入：s = "pwwkew"
 * 输出：3
 * 解释：因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 5 * 10^4
 * s 由英文字母、数字、符号和空格组成
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 思路：和上一题类似，用两个指针创建窗口，当右指针遇到和上一个指向的对象相同的字符时，记录下此时窗口长度，并将左指针指向当前右指针指向的对象，右指针继续右移，直到遇到和上一个指向的对象相同的字符，记录下最大长度
// 又读错题了，是不含有重复字符的最长子串，而不是不含有连续重复字符的最长子串
// 需要建立哈希表来处理是否重复，用左指针来框选窗口，当出现重复时，将左指针移动到下一位，并将遇到的数都移出哈希表，直到遇到和右指针指向元素重复的字符，将其移除后移动到下一位，停止移动等待下一个重复
func lengthOfLongestSubstring(s string) int {
	right, ans, left := 0, 1, 0
	// hashMap := make(map[byte]int)
	hashMap := [128]int{}
	//[128] 代表 ASCII 的全部字符对应的数字编码
	//可以用这个代替 map，因为 map 的查询效率更低？

	//哈希表时刻表示 left 和 right 之间字符出现次数
	if len(s) == 1 {
		return 1
	}
	if len(s) == 0 {
		return 0
	}
	// length := 0
	//窗口实时长度

	for right < len(s) {
		hashMap[s[right]] += 1
		//将哈希表对应的字符键创建并记值为 1，如果已存在就加 1
		//窗口实时长度
		for hashMap[s[right]] == 2 {
			//当出现重复字符时
			hashMap[s[left]]--
			//将 left 指向的字符键的值减一
			//这一步是将从 left 到 right 指向的重复字符之间的所有出现的字符移出哈希表，该哈希表时刻表示 left 和 right 之间字符出现次数
			left++
			//left 右移
			//窗口长度减一
			//直到查找到那个和当前 right 指向字符相同的字符，将其减一
			//并将 left 移动到其下一位
		}
		right++

		//right 右移
		ans = max(ans, right-left)
		//可以使用 right-left 维护实时窗口长度
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

*/
