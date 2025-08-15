/*
 * @lc app=leetcode.cn id=49 lang=golang
 * @lcpr version=30204
 *
 * [49] 字母异位词分组
 *
 * https://leetcode.cn/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (69.62%)
 * Likes:    2260
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 1.5M
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
 *
 * 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
 *
 *
 *
 * 示例 1:
 *
 * 输入：strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
 * 输出：[["bat"],["nat","tan"],["ate","eat","tea"]]
 *
 * 示例 2:
 *
 * 输入：strs = [""]
 * 输出：[[""]]
 *
 *
 * 示例 3:
 *
 * 输入：strs = ["a"]
 * 输出：[["a"]]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= strs.length <= 10^4
 * 0 <= strs[i].length <= 100
 * strs[i] 仅包含小写字母
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 建立 ans string 二维切片，给每个词建立哈希表，并建立一个二维哈希表作为表的存储池，当当前词哈希表在池中时，将其加入 ans 切片中的和存储池一样索引的位置
// 这样一个一个对比太过繁杂，会让时间复杂度趋向 On^2
// 我们可以建立一个键为词频，值为同词频单词切片的 map:anagramGroup，和之前思路一样建议一个 wordCount 遍历每个单词的词频，将有着相同词频的单词加入对应切片中，由于 map 是可变长度的，当没有对应键时会直接加入对应键，所以可以直接使用 append 函数
func groupAnagrams(strs []string) [][]string {
	// if len(strs) == 1 || len(strs) == 0 {
	// 	ans = append(ans, strs)
	// 	return
	// }
	// wordMapPoll := make([][26]int, 0)

	// for _, word := range strs {
	// 	wordMap := [26]int{}
	// 	for _, char := range word {
	// 		char = char - 'a'
	// 		wordMap[char]++
	// 	}
	// 	isWordInPoll := false
	// 	for i, val := range wordMapPoll {
	// 		if val == wordMap {
	// 			isWordInPoll = true
	// 			ans[i] = append(ans[i], word)
	// 			break
	// 		}
	// 	}
	// 	if !isWordInPoll {
	// 		wordMapPoll = append(wordMapPoll, wordMap)
	// 		wordString := []string{word}
	// 		ans = append(ans, wordString)
	// 	}
	// }
	// return
	//关键是用词频数组直接做哈希表的索引键
	if len(strs) == 0 {
		return [][]string{}
	}
	if len(strs) == 1 {
		return [][]string{strs} // 直接返回包含单个字符串的二维切片
	}

	// 使用 map 来存储，key 是字符频率数组 [26]int，value 是该频率对应的字符串列表 []string
	anagramGroups := make(map[[26]int][]string)

	for _, word := range strs {
		// 1. 计算当前单词的字符频率
		charCount := [26]int{} // 数组会自动初始化为零值
		for _, char := range word {
			charCount[char-'a']++
		}

		// 2. 使用字符频率数组作为 key，将单词添加到对应的分组
		// 如果 charCount 这个 key 不存在，会自动创建一个空 []string
		anagramGroups[charCount] = append(anagramGroups[charCount], word)
	}

	// 3. 将 map 中的分组收集到结果切片中
	ans := make([][]string, 0, len(anagramGroups)) // 预分配容量
	for _, group := range anagramGroups {
		ans = append(ans, group)
	}

	return ans
}

// @lc code=end

/*
// @lcpr case=start
// ["eat", "tea", "tan", "ate", "nat", "bat"]\n
// @lcpr case=end

// @lcpr case=start
// [""]\n
// @lcpr case=end

// @lcpr case=start
// ["a"]\n
// @lcpr case=end

*/
