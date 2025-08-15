/*
 * @lc app=leetcode.cn id=30 lang=golang
 * @lcpr version=30204
 *
 * [30] 串联所有单词的子串
 *
 * https://leetcode.cn/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (38.18%)
 * Likes:    1218
 * Dislikes: 0
 * Total Accepted:    252.4K
 * Total Submissions: 662K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * 给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。
 *
 * s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。
 *
 *
 * 例如，如果 words = ["ab","cd","ef"]，那么 "abcdef"， "abefcd"，"cdabef"，
 * "cdefab"，"efabcd"，和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。
 *
 *
 * 返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "barfoothefoobarman", words = ["foo","bar"]
 * 输出：[0,9]
 * 解释：因为 words.length == 2 同时 words[i].length == 3，连接的子字符串的长度必须为 6。
 * 子串 "barfoo" 开始位置是 0。它是 words 中以 ["bar","foo"] 顺序排列的连接。
 * 子串 "foobar" 开始位置是 9。它是 words 中以 ["foo","bar"] 顺序排列的连接。
 * 输出顺序无关紧要。返回 [9,0] 也是可以的。
 *
 *
 * 示例 2：
 *
 * 输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
 * 输出：[]
 * 解释：因为 words.length == 4 并且 words[i].length == 4，所以串联子串的长度必须为 16。
 * s 中没有子串长度为 16 并且等于 words 的任何顺序排列的连接。
 * 所以我们返回一个空数组。
 *
 *
 * 示例 3：
 *
 * 输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
 * 输出：[6,9,12]
 * 解释：因为 words.length == 3 并且 words[i].length == 3，所以串联子串的长度必须为 9。
 * 子串 "foobarthe" 开始位置是 6。它是 words 中以 ["foo","bar","the"] 顺序排列的连接。
 * 子串 "barthefoo" 开始位置是 9。它是 words 中以 ["bar","the","foo"] 顺序排列的连接。
 * 子串 "thefoobar" 开始位置是 12。它是 words 中以 ["the","foo","bar"] 顺序排列的连接。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * 1 <= words.length <= 5000
 * 1 <= words[i].length <= 30
 * words[i] 和 s 由小写英文字母组成
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start

// 从头开始，每个字符都切一个长度为 len(word[0]) 的片，也就是 len(s) ,遍历 words，建立 words 词频哈希表，当切片存在于 words 哈希表中时，对应--,并记录当前位置，开始进入 if 一一求长度 len(word[0]) 切片，直到 words 哈希表全部为 0，返回记录位置，并从记录位置+len(word[0]) 开始重新匹配
// 这样的解法的时间复杂度是 O(len(s)*len(words))
// map 是按引用赋值
// 双哈希双指针比较法，直到满足原始哈希表为止
func findSubstring(s string, words []string) []int {
	if len(s) < len(words[0])*len(words) {
		return []int{}
	}
	totalLength := len(words) * len(words[0])
	//匹配字符总长度
	ans := make([]int, 0)
	wordLength := len(words[0])
	//单词长度
	originalHash := make(map[string]int)
	//原始哈希表
	for _, val := range words {
		originalHash[val]++
	} //求出 words 哈希表
	for i := 0; i+totalLength <= len(s); i++ {

		substr := s[i : i+wordLength]
		//word 长度大小的子字符串
		if originalHash[substr] != 0 {
			//查询该子字符串是否在哈希表中
			//如果存在则构建一个新的表
			wordsHash := make(map[string]int)
			//创建空哈希表
			locate := i
			//确定当前位置
			wordsHash[substr]++
			wordFound := 1
			for n := i + wordLength; n+wordLength <= i+totalLength; n += wordLength {
				substr = s[n : n+wordLength]
				//找到下一个长度为 word 的子串
				originalCount, ok := originalHash[substr]
				//如果原始哈希表中存在该子串向将当前哈希表上该子串出现次数++
				if ok {
					wordsHash[substr]++
				} else {
					break
					//如不存在，终止循环
				}
				if wordsHash[substr] > originalCount {
					break
					//如果大于原始表，终止循环
					//双哈希互相比较，因为哈希表不能按值传递
				}
				//当哈希表中不存在当前子串或者当前哈希表中的某子串出现次数大于原始哈希表中的值，终止循环
				wordFound++
				//如果以上都没有发生，则找到了一个单词
				//即原始哈希表中存在该子串，且当前哈希表中该子串出现次数没有超过原始哈希表
			}
			if wordFound == len(words) {
				//找满单词就返回
				ans = append(ans, locate)
			}
		}
	}
	return ans
	//这样一一遍历会产生大量冗余
	//更优解法是按单词偏移量 [0:len(words[0])] 来分开，以单词为单位枚举全部单词
	//时刻保持当前窗口长度为 totalLength，当窗口中的单词满足 originalHash 时，将当前左指针加入 ans 数组，将左右指针右移一个单词长度的位置，直到右指针抵达尾端
	//即维护一个长度时刻为 totalLength 的窗口，并记录其中的单词状态，当状态达到目标状态时，将其左指针记录进入 ans
	//这样的时间复杂度是 O(len(s))
	//减少冗余查询，只需要吧 word 的偏移量查完既可
	// LeetCode 题目：串联所有单词的子串

	// 函数签名，仅包含核心逻辑

	// --- 初始化和边界条件检查 ---
	// if len(s) == 0 || len(words) == 0 {
	// 	return []int{} // 如果 s 或 words 为空，直接返回空切片
	// }

	// wordLength := len(words[0])          // 单个单词的长度
	// numWords := len(words)               // words 数组中单词的总数
	// totalLength := wordLength * numWords // 需要匹配的子串总长度
	// ans := make([]int, 0)                // 存储结果的切片

	// // 如果 s 的长度小于所需子串的总长度，不可能找到匹配
	// if totalLength > len(s) {
	// 	return []int{}
	// }

	// // --- 预处理：统计 words 中每个单词的期望出现次数 ---
	// originalHash := make(map[string]int)
	// for _, word := range words {
	// 	originalHash[word]++
	// }
	// // 需要匹配的不同单词种类数 (用于优化匹配判断)
	// uniqueWordsCount := len(originalHash)

	// // --- 滑动窗口核心逻辑 ---
	// // 只需要遍历 wordLength 个起始偏移量
	// for i := 0; i < wordLength; i++ {
	// 	left := i                                 // 当前滑动窗口的左边界索引
	// 	currentWindowHash := make(map[string]int) // 当前窗口内各单词的计数
	// 	matches := 0                              // 当前窗口内满足 originalHash 计数的单词种类数

	// 	// right 是窗口的右边界索引（取到的单词是 s[right:right+wordLength]）
	// 	// 每次滑动 wordLength 步长
	// 	for right := i; right+wordLength <= len(s); right += wordLength {

	// 		// 1. 处理进入窗口的单词 (最右侧的单词)
	// 		word := s[right : right+wordLength]
	// 		originalCount, needed := originalHash[word] // 检查这个单词是否是我们需要的

	// 		if needed {
	// 			// 如果是需要的单词：
	// 			currentWindowHash[word]++ // 在当前窗口计数中增加该单词

	// 			// 更新 matches 计数：
	// 			if currentWindowHash[word] == originalCount {
	// 				matches++ // 该单词的数量正好达到要求，匹配种类数 +1
	// 			} else if currentWindowHash[word] == originalCount+1 {
	// 				// 注意：仅当数量从 "正好" 变为 "超出" 时，匹配种类数才需要 -1
	// 				matches-- // 该单词的数量刚刚超过要求，匹配种类数 -1
	// 			}

	// 			// 2. 判断窗口是否过长，需要从左侧移除单词
	// 			// 当前窗口包含的单词数量 = (当前右边界 - 左边界) / 单词长度
	// 			windowSizeWords := (right + wordLength - left) / wordLength
	// 			if windowSizeWords > numWords {
	// 				leftWord := s[left : left+wordLength] // 获取将要离开窗口的单词
	// 				left += wordLength                    // 将左边界右移 wordLength

	// 				// 如果离开的单词是需要的单词，需要更新 currentWindowHash 和 matches
	// 				// 注意：必须检查，因为可能之前有不需要的词导致窗口扩大
	// 				if _, ok := originalHash[leftWord]; ok {
	// 					leftOriginalCount := originalHash[leftWord]
	// 					// 在更新计数 *之前*，先根据当前状态调整 matches
	// 					if currentWindowHash[leftWord] == leftOriginalCount {
	// 						matches-- // 这个单词之前是匹配的，现在要离开了，匹配种类数 -1
	// 					} else if currentWindowHash[leftWord] == leftOriginalCount+1 {
	// 						matches++ // 这个单词之前是超出的，现在离开一个正好匹配，匹配种类数 +1
	// 					}
	// 					currentWindowHash[leftWord]-- // 更新当前窗口计数

	// 					// 可选：如果单词计数变为 0，可以从 map 中删除以节省空间
	// 					if currentWindowHash[leftWord] == 0 {
	// 						delete(currentWindowHash, leftWord)
	// 					}
	// 				}
	// 			}

	// 			// 3. 检查是否找到了一个解
	// 			// 当窗口内满足数量要求的单词种类数 等于 所需的不同单词总数时
	// 			if matches == uniqueWordsCount {
	// 				// 此时的 left 就是一个有效子串的起始索引
	// 				ans = append(ans, left)
	// 			}

	// 		} else {
	// 			// 如果遇到的单词 (word) 不是 words 中的任何一个：
	// 			// 那么从 left 到 right 的这个窗口肯定不满足条件了，需要重置窗口
	// 			currentWindowHash = make(map[string]int) // 清空当前窗口计数
	// 			matches = 0                              // 重置匹配计数
	// 			left = right + wordLength                // 将左边界移动到这个无效单词之后，开始新的窗口
	// 		}
	// 	}
	// }

	// return ans // 返回所有找到的起始索引
}

// @lc code=end

/*
// @lcpr case=start
// "barfoothefoobarman"\n["foo","bar"]\n
// @lcpr case=end

// @lcpr case=start
// "wordgoodgoodgoodbestword"\n["word","good","best","word"]\n
// @lcpr case=end

// @lcpr case=start
// "barfoofoobarthefoobarman"\n["bar","foo","the"]\n
// @lcpr case=end

*/
