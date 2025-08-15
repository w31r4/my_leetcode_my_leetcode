/*
 * @lc app=leetcode.cn id=98 lang=golang
 * @lcpr version=30204
 *
 * [98] 验证二叉搜索树
 */

// @lcpr-template-start
package leetcode

import "math"

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

func isValidBST(root *TreeNode) (ans bool) {
	// if root == nil {
	// 	return true
	// }
	// isBst := true
	// if root.Left != nil {
	// 	isBst = root.Val > root.Left.Val
	// }
	// if root.Right != nil {
	// 	isBst = root.Val < root.Right.Val
	// }
	// return isBst && isValidBST(root.Left) && isValidBST(root.Right)
	// BST 是从头开始比较的，所以不能单单用左小右大来判断
	//思路错误
	// 换一个思路：中序遍历获得展开后的二叉搜索树，如果出现前者大于后者那么就返回 false
	//引入外部变量控制剪枝不如函数内部可以返回 bool 值控制剪枝
	// pre := math.MinInt64
	// var midDfs func(node *TreeNode)
	// isGetTrouble := false
	// midDfs = func(node *TreeNode) {
	// 	if isGetTrouble {
	// 		return
	// 	}
	// 	if node == nil {
	// 		return
	// 	}
	// 	midDfs(node.Left)
	// 	if pre < node.Val {
	// 		pre = node.Val
	// 	} else {
	// 		isGetTrouble = true
	// 		return
	// 	}
	// 	//中序遍历更新 pre
	// 	midDfs(node.Right)
	// 	return

	// }
	// midDfs(root)
	// return !isGetTrouble
	pre := math.MinInt
	//函数返回 bool 值，更新 pre 的同时返回判断情况，如果判断情况为正确继续下一个 bst 节点的判断
	var midDfs func(node *TreeNode) bool
	midDfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !midDfs(node.Left) {
			//错误就返回错误，对就返回对
			return false
		}
		//中序遍历更新 pre
		if node.Val <= pre {
			return false
		}
		pre = node.Val
		return midDfs(node.Right)

	}
	return midDfs(root)
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [5,1,4,null,null,3,6]\n
// @lcpr case=end

*/
