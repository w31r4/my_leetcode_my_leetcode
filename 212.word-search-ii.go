/*
 * @lc app=leetcode.cn id=212 lang=golang
 * @lcpr version=30204
 *
 * [212] 单词搜索 II
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
type letterNode struct {
	child   [26]*letterNode
	wordNum int
	//wordNum 用于标记这是 wordList 中第几号单词的结尾位置
}
type letterTrie struct {
	root *letterNode
}

func constructTrie() letterTrie {
	return letterTrie{
		&letterNode{
			[26]*letterNode{},
			// false,
			-1,
		},
	}
}
func (this *letterTrie) insert(num int, word string) {
	node := this.root
	for _, char := range word {
		char -= 'a'
		if node.child[char] == nil {
			node.child[char] = &letterNode{[26]*letterNode{}, -1}
		} //如果为空，那么将其初始化使其有内容
		node = node.child[char]
	}
	// node.isEnd = true
	node.wordNum = num
}

func findWords(board [][]byte, words []string) (answer []string) {
	// 这玩意和前缀树有什么关联吗？
	//如果，我们将 words 构成前缀树，然后将 board 内字母进行遍历
	//如果当前字母在第一层，那么就遍历周围的字母是否有符合第二层的。出现一次 isEnd 向 answer 里加一个
	//直到找不到新字母停止循环
	colLen := len(board)
	rowLen := len(board[0])
	letterTrie := constructTrie()
	for i := 0; i < len(words); i++ {
		//将单词表插入前缀树中
		letterTrie.insert(i, words[i])
	}
	visited := make([][]bool, colLen)
	for i, _ := range visited {
		visited[i] = make([]bool, rowLen)
	}
	//首先我们得记录当前状态
	var dfs func(i, j int, letter *letterNode)
	dfs = func(i, j int, letter *letterNode) {
		//i j 是当前遍历到的节点的位置
		if i >= colLen || j >= rowLen || i < 0 || j < 0 || visited[i][j] {
			//边界条件
			return
		}
		ch := board[i][j] - 'a'
		if letter.child[ch] != nil {
			visited[i][j] = true

			if letter.child[ch].wordNum != -1 {
				answer = append(answer, words[letter.child[ch].wordNum])
				letter.child[ch].wordNum = -1
				//防止重复
			}
			dfs(i+1, j, letter.child[ch])
			dfs(i-1, j, letter.child[ch])
			dfs(i, j+1, letter.child[ch])
			dfs(i, j-1, letter.child[ch])
			visited[i][j] = false

		}

		return
	}
	//接下来怎么进行查找呢
	for i := 0; i < colLen; i++ {
		//第几行
		for j := 0; j < rowLen; j++ {
			//第几列
			dfs(i, j, letterTrie.root)
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]\n["oath","pea","eat","rain"]\n
// @lcpr case=end

// @lcpr case=start
// [["a","b"],["c","d"]]\n["abcb"]\n
// @lcpr case=end

*/
