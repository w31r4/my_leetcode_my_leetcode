/*
 * @lc app=leetcode.cn id=79 lang=golang
 * @lcpr version=30204
 *
 * [79] 单词搜索
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 获取棋盘中所有字母出现的次数
// 由于字母包含大小写，所以有 52 个
func getBoardLetterTimes(board [][]byte) [52]int {
	var letterMap [52]int // Our array is now size 52

	for _, line := range board {
		for _, ch := range line {
			if ch >= 'a' && ch <= 'z' {
				// It's a lowercase letter, map to 0-25
				letterMap[ch-'a']++
			} else if ch >= 'A' && ch <= 'Z' {
				// It's an uppercase letter, map to 26-51
				letterMap[ch-'A'+26]++
			}
		}
	}
	return letterMap
}

// 获取单词中所有字母出现的次数
// 由于字母包含大小写，所以有 52 个
func getWordLetterTimes(word string) (wordLetterCounts [52]int) {
	for _, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			// It's a lowercase letter, map to 0-25
			wordLetterCounts[ch-'a']++
		} else if ch >= 'A' && ch <= 'Z' {
			// It's an uppercase letter, map to 26-51
			wordLetterCounts[ch-'A'+26]++
		}
	}
	return wordLetterCounts
}

// 通过 rune 切片将 string 反转
func reverseStringRunes(s string) string {
	// 1. 将 string 转换为 []rune
	runes := []rune(s)

	// 2. 使用双指针法交换元素
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// 3. 将 []rune 转换回 string
	return string(runes)
}
func exist(board [][]byte, word string) (answer bool) {
	// --- 预处理与优化阶段 ---

	// 1. 统计棋盘（供应）和单词（需求）的字符频率。
	boardCounts := getBoardLetterTimes(board)
	wordCounts := getWordLetterTimes(word)

	// 2.【优化点 1: 快速拒绝】
	// 思想：如果棋盘上的某个字符数量，连单词所需要的都满足不了，
	// 那么绝对不可能拼出这个单词，可以直接返回 false，避免无效的搜索。
	for i := 0; i < 52; i++ {
		if wordCounts[i] > boardCounts[i] {
			return false // 使用 return, 因为函数定义了 answer 默认为 false
		}
	}

	// 3.【优化点 2: 启发式搜索方向】
	// 思想：深度优先搜索的开销很大。我们应该从棋盘上更稀有的字母开始搜索，
	// 这样可以大幅减少启动 DFS 的次数。
	length := len(word)
	var firstLetterCount, lastLetterCount int

	// 获取首字母在棋盘上的数量
	firstChar := word[0]
	if firstChar >= 'a' && firstChar <= 'z' {
		firstLetterCount = boardCounts[firstChar-'a']
	} else {
		firstLetterCount = boardCounts[firstChar-'A'+26]
	}

	// 获取尾字母在棋盘上的数量
	lastChar := word[length-1]
	if lastChar >= 'a' && lastChar <= 'z' {
		lastLetterCount = boardCounts[lastChar-'a']
	} else {
		lastLetterCount = boardCounts[lastChar-'A'+26]
	}

	// 核心决策：如果首字母比尾字母更常见，就反转单词，从更稀有的尾字母开始搜。
	if firstLetterCount > lastLetterCount {
		word = reverseStringRunes(word)
	}
	//以上是关于提高效率的优化
	//下面是核心搜索代码
	//一段带回溯性 visited 的 dfs 遍历
	visited := make([][]bool, len(board))
	m := len(board)
	n := len(board[0])
	for i, _ := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	var dfs func(i, j, level int)
	dfs = func(i, j, level int) {

		if answer {
			return
		}
		if level == length {
			answer = true
			return
		}
		if i < 0 || j < 0 || i == m || j == n {
			return
		}
		if visited[i][j] {
			return
		}
		if board[i][j] != word[level] {
			return
		}

		visited[i][j] = true
		dfs(i+1, j, level+1)
		dfs(i, j+1, level+1)
		dfs(i-1, j, level+1)
		dfs(i, j-1, level+1)
		visited[i][j] = false
		return
	}
	//使用了外部变量 answer 进行剪枝，会导致代码高度耦合
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !answer {
				dfs(i, j, 0)
			} else {
				break
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']]\n"ABCCED"\n
// @lcpr case=end

// @lcpr case=start
// [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']]\n"SEE"\n
// @lcpr case=end

// @lcpr case=start
// [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']]\n"ABCB"\n
// @lcpr case=end

*/
