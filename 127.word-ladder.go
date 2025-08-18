/*
 * @lc app=leetcode.cn id=127 lang=golang
 * @lcpr version=30204
 *
 * [127] 单词接龙
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 和 433 高度类似
func isDiffOne(word1, word2 string) bool {
	length := len(word1)
	diff := 0
	for i := 0; i < length; i++ {
		if word1[i] != word2[i] {
			diff++
		}
	}
	if diff == 1 {
		return true
	}
	return false
}

// 用 BFS 来做
func ladderLength(beginWord string, endWord string, wordList []string) int {
	length := len(wordList)
	visited := make([]bool, length)
	wordQueue := []string{}
	wordQueue = append(wordQueue, beginWord)
	step := 1
	for wordQueue != nil {
		temp := wordQueue
		wordQueue = nil
		for i := 0; i < len(temp); i++ {
			if temp[i] == endWord {
				return step
			}
			for j := 0; j < length; j++ {
				if !visited[j] && isDiffOne(temp[i], wordList[j]) {
					wordQueue = append(wordQueue, wordList[j])
					visited[j] = true
				}
			}
		}
		step++
	}
	return 0
}

// @lc code=end

/*
// @lcpr case=start
// "hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]\n
// @lcpr case=end

// @lcpr case=start
// "hit"\n"cog"\n["hot","dot","dog","lot","log"]\n
// @lcpr case=end

*/
