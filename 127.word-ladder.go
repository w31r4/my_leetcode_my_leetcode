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
// 试试双向 BFS
//双向 BFS 的思路如下：
//做两个标记数组，一个从前开始，一个从后开始标记
//前后交替进行 BFS 遍历，当前标记中出现后标记的内容，终止循环，返回和长度

func ladderLength(beginWord string, endWord string, wordList []string) int {
	length := len(wordList)
	frontVisited := make([]bool, length)
	backVisited := make([]bool, length)
	frontQueue, backQueue := []string{}, []string{}
	//建立双向队列和标记组
	frontQueue = append(frontQueue, beginWord)
	//初始化队列
	frontStep := 1
	backStep := 1
	isFrontNow := true
	//得初始化一下 back
	for index, word := range wordList {
		if word == endWord {
			backQueue = append(backQueue, word)
			backVisited[index] = true

			break
		}
	}
	//如果 backQueue 为空，直接终止循环，返回 0，因为 wordList 中没有需要的单词
	for backQueue != nil || frontQueue != nil {
		var temp []string
		var queue *[]string
		var leftVisited, visited *[]bool
		var step *int
		if isFrontNow {
			temp = frontQueue
			frontQueue = nil
			queue = &frontQueue
			visited = &frontVisited
			leftVisited = &backVisited
			step = &frontStep
		} else {
			temp = backQueue
			backQueue = nil
			queue = &backQueue
			visited = &backVisited
			leftVisited = &frontVisited
			step = &backStep
		}
		for i := 0; i < len(temp); i++ {
			//当去除作比较的单词存在于另一个 visited 中时，出现返回值
			for j, word := range wordList {
				if !(*visited)[j] && isDiffOne(temp[i], word) {
					//当该单词未在当前记录中出现且符合变化时
					if (*leftVisited)[j] {
						//当该单词被另一个记录表记录
						//找到了共同点，返回结果
						return backStep + frontStep
					}
					//更新当前记录表
					(*visited)[j] = true
					//扩充当前队列
					*queue = append(*queue, word)
				}
			}
		}
		isFrontNow = !isFrontNow
		*step += 1
	}
	// wordQueue := []string{}
	// wordQueue = append(wordQueue, beginWord)
	// step := 1
	// for wordQueue != nil {
	// 	temp := wordQueue
	// 	wordQueue = nil
	// 	for i := 0; i < len(temp); i++ {
	// 		if temp[i] == endWord {
	// 			return step
	// 		}
	// 		for j := 0; j < length; j++ {
	// 			if !visited[j] && isDiffOne(temp[i], wordList[j]) {
	// 				wordQueue = append(wordQueue, wordList[j])
	// 				visited[j] = true
	// 			}
	// 		}
	// 	}
	// 	step++
	// }
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
