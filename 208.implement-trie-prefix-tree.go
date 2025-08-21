/*
 * @lc app=leetcode.cn id=208 lang=golang
 * @lcpr version=30204
 *
 * [208] 实现 Trie (前缀树)
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
type TrieNode struct {
	child map[rune]*TrieNode
	isEnd bool
}

// 前缀树是一种数，用于高效检索保存字符串数据
type Trie struct {
	root *TrieNode
}

func constructor() Trie {
	return Trie{
		root: &TrieNode{
			child: make(map[rune]*TrieNode),
			isEnd: false,
		},
	}
}

// 地道地实现了向前缀树（Trie）中插入一个单词的功能。从根节点开始，逐个字符检查路径，如果路径不存在就创建，最后在单词的末尾打上标记。
func (this *Trie) Insert(word string) {
	node := this.root
	for _, char := range word {
		if node.child[char] != nil {
			node = node.child[char]
		} else {
			node.child[char] = &TrieNode{map[rune]*TrieNode{}, false}
			node = node.child[char]
		}
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, char := range word {
		if node.child[char] != nil {
			node = node.child[char]
		} else {
			return false
		}
	}
	if node.isEnd {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, char := range prefix {
		if node.child[char] != nil {
			node = node.child[char]
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
