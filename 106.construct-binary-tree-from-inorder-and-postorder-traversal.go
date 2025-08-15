/*
 * @lc app=leetcode.cn id=106 lang=golang
 * @lcpr version=30204
 *
 * [106] 从中序与后序遍历序列构造二叉树
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

// 和上题思路一样，不过从先序表中获取信息变成了从后序表中获取根节点信息
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	endOfPostOrder := len(postorder) - 1
	root := &TreeNode{Val: postorder[endOfPostOrder]}

	spiltPoint := 0
	for i, val := range inorder {
		if val == root.Val {
			spiltPoint = i
		}
	}
	leftInorder := inorder[:spiltPoint]
	leftPostorder := postorder[:spiltPoint]
	root.Left = buildTree(leftInorder, leftPostorder)
	rightInorder := inorder[spiltPoint+1:]
	rightPostorder := postorder[spiltPoint:endOfPostOrder]
	root.Right = buildTree(rightInorder, rightPostorder)
	return root
}

// @lc code=end

/*
// @lcpr case=start
// [9,3,15,20,7]\n[9,15,7,20,3]\n
// @lcpr case=end

// @lcpr case=start
// [-1]\n[-1]\n
// @lcpr case=end

*/
