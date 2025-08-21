/*
 * @lc app=leetcode.cn id=211 lang=golang
 * @lcpr version=30204
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 优化解法，通过 26 位数组来优化哈希表
type wordNode struct {
	children [26]*wordNode
	isEnd    bool
}

func (t *wordNode) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &wordNode{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

type WordDictionary struct {
	trieRoot *wordNode
}

func Constructor() WordDictionary {
	return WordDictionary{&wordNode{}}
}

func (d *WordDictionary) AddWord(word string) {
	d.trieRoot.Insert(word)
}

func (d *WordDictionary) Search(word string) bool {
	var dfs func(int, *wordNode) bool
	dfs = func(index int, node *wordNode) bool {
		if index == len(word) {
			return node.isEnd
		}
		ch := word[index]
		if ch != '.' {
			child := node.children[ch-'a']
			if child != nil && dfs(index+1, child) {
				return true
			}
		} else {
			for i := range node.children {
				child := node.children[i]
				if child != nil && dfs(index+1, child) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, d.trieRoot)
}

// type WordNode struct {
// 	child map[rune]*WordNode
// 	isEnd bool
// }
// type WordDictionary struct {
// 	root *WordNode
// }

// func Constructor() WordDictionary {
// 	return WordDictionary{
// 		&WordNode{
// 			child: make(map[rune]*WordNode),
// 			isEnd: false},
// 	}
// }

// func (this *WordDictionary) AddWord(word string) {
// 	node := this.root
// 	for _, char := range word {
// 		if node.child[char] != nil {
// 			node = node.child[char]
// 		} else {
// 			node.child[char] = &WordNode{make(map[rune]*WordNode), false}
// 			node = node.child[char]
// 		}
// 	}
// 	node.isEnd = true
// }

// func (this *WordDictionary) Search(word string) bool {
// 	return this.searchHelper(this.root, word)
// }
// func (this *WordDictionary) searchHelper(node *WordNode, word string) bool {
// 	//这个 dfs 该如何进行呢？
// 	//当前要处理的字符有两种情况，一种是通配符 . ,一种是正常字母
// 	//先来处理一下边界条件
// 	if word == "" {
// 		//当处理到 word 为空时
// 		if node.isEnd {
// 			return true
// 		}
// 		return false
// 	}
// 	firstChar := rune(word[0])
// 	if firstChar == '.' {
// 		for _, char := range node.child {
// 			if this.searchHelper(char, word[1:]) {
// 				//只要出现一个 true，直接返回 true
// 				return true
// 			}
// 		}
// 		return false
// 	} else {
// 		if node.child[firstChar] != nil {
// 			return this.searchHelper(node.child[firstChar], word[1:])
// 		}
// 		return false

// 	}
// }

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

/*
// @lcpr case=start
// ["WordDictionary","addWord","addWord","addWord","search","search","search","search"][[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]\n
// @lcpr case=end

*/
