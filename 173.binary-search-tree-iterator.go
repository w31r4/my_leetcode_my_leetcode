/*
 * @lc app=leetcode.cn id=173 lang=golang
 * @lcpr version=30204
 *
 * [173] 二叉搜索树迭代器
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }
type BSTIterator struct {
	// BstQueue       []int
	// Length         int
	// nowPointLength int
	BSTStack []*TreeNode
}

// 正常人的思路：将沟槽的二叉搜索树中序遍历后存储在切片中
// 一个一个取出存储好的元素
// 节约空间思路：用栈模拟中序遍历：先找到左子树底 (一定是整棵树最小的)，并把路径上的所有节点入栈
// 当 next 时，出栈栈顶元素，它一定是最小的，判断栈顶元素是否存在 right(right 一定比 node 大),如果存在将其入栈并遍历至 right 的 left 底部，该底部为除当前节点最小值
func Constructor5(root *TreeNode) BSTIterator {
	// var bstQueue []int
	// var midDfs func(node *TreeNode)
	// midDfs = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	midDfs(node.Left)
	// 	bstQueue = append(bstQueue, node.Val)
	// 	midDfs(node.Right)
	// 	return
	// }
	// midDfs(root)
	// return BSTIterator{BstQueue: bstQueue, Length: len(bstQueue), nowPointLength: 0}
	stack := []*TreeNode{}
	node := root
	for node != nil {
		stack = append(stack, node)
		node = node.Left
	}
	return BSTIterator{stack}
}

func (this *BSTIterator) Next() int {
	// nowVal := this.BstQueue[this.nowPointLength]
	// this.nowPointLength++
	// return nowVal
	nowNode := this.BSTStack[len(this.BSTStack)-1]
	nowVal := nowNode.Val
	this.BSTStack = this.BSTStack[:len(this.BSTStack)-1]
	// curr := this.BSTStack[len(this.BSTStack)-1]
	curr := nowNode
	//从当前节点右节点开始找，而不是当前节点父节点
	if curr.Right != nil {
		curr = curr.Right
		for curr != nil {
			this.BSTStack = append(this.BSTStack, curr)
			curr = curr.Left
		}
	}
	return nowVal
}

func (this *BSTIterator) HasNext() bool {
	// return !(this.nowPointLength == this.Length)
	return len(this.BSTStack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end
