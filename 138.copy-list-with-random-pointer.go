/*
 * @lc app=leetcode.cn id=138 lang=golang
 * @lcpr version=30204
 *
 * [138] 随机链表的复制
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
// type Node struct {
// 	Val    int
// 	Next   *Node
// 	Random *Node
// }

//法一哈希表：建立新旧节点对应表，先建立后 random

// 法二原地构建：三次循环，第一次构建复制后的链表
// 第二次将复制后的链表元素 random 赋值
// 第三次分离复制前链表和复制后链表
// 不要一边赋值 random 一遍分离链表，链表的 random 赋值操作需要保持链表整体保持不变，如果改变了的话那么 random 指向的值将不会是预期的值
// 分离两个链表会修改节点.Next，导致 copyNode.Random = lastRandom.Next 语句产生歧义，无法实现目标功能
func copyRandomList(head *Node) *Node {
	recordHead := head
	for head != nil {
		newNode := &Node{head.Val, head.Next, nil}
		head.Next = newNode
		head = newNode.Next
	}
	//在原链表后复制新链表
	currentNode := recordHead
	for currentNode != nil {
		lastRandom := currentNode.Random
		//记录原节点 random 指向节点
		copyNode := currentNode.Next
		//记录被复制节点
		if lastRandom != nil {
			copyNode.Random = lastRandom.Next
			//当原节点 random 指向节点不为 nil 时，将当前被复制节点 random 指向原 random 节点下一个节点
		}
		// if isFirstTime {
		// 	ans = copyNode
		// 	ansHead = ans
		// 	isFirstTime = false
		// } else {
		// 	ans.Next = copyNode
		// 	ans = ans.Next
		// 	//这里对 ans 的修改会同步到当前长链表中，导致原来的链表发生改变，所以建议不要边修改元素边修改链表指向
		// }
		currentNode = copyNode.Next
		//配置 copy 后的 random
	}
	var ans, ansHead *Node
	isFirstTime := true
	originHead := recordHead
	for originHead != nil {
		if isFirstTime {
			ans = originHead.Next
			originHead.Next = ans.Next
			originHead = originHead.Next
			ansHead = ans
			isFirstTime = false
		} else {
			ans.Next = originHead.Next
			ans = ans.Next
			originHead.Next = ans.Next
			originHead = originHead.Next
		}
		//将一切分开
	}
	return ansHead
}

// @lc code=end

/*
// @lcpr case=start
// [[7,null],[13,0],[11,4],[10,2],[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1],[2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[3,null],[3,0],[3,null]]\n
// @lcpr case=end

*/
