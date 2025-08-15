/*
 * @lc app=leetcode.cn id=114 lang=golang
 * @lcpr version=30204
 *
 * [114] 二叉树展开为链表
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

// 先序遍历为链表
// 递归先序遍历，传入 dummy 结点
// 闭包递归
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	dummy := &TreeNode{}
	currTail := dummy
	var frontRecursive func(node *TreeNode)
	frontRecursive = func(node *TreeNode) {
		if node == nil {
			return
		}
		//边界条件是 node==nil
		leftNode := node.Left
		rightNode := node.Right
		currTail.Right = node
		//尾插法创建链表
		currTail.Left = nil
		currTail = node
		frontRecursive(leftNode)
		//先序遍历递归
		frontRecursive(rightNode)
		return
	}
	frontRecursive(root)
	root = dummy.Right
}

// 该怎么保存更新后的 listHead 呢？
// func frontRecursive(node *TreeNode, listHead *TreeNode) {
// 	if node == nil {
// 		return
// 	}
// 	leftNode := node.Left
// 	rightNode := node.Right
// 	listHead.Right = node
// 	frontRecursive(leftNode, listHead.Right)
// 	node.Left = nil
// 	frontRecursive(rightNode, listHead.Right)
// 	return
// }

// @lc code=end

/*
// @lcpr case=start
// [1,2,5,3,4,null,6]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/
