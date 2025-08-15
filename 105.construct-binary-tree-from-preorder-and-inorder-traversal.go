/*
 * @lc app=leetcode.cn id=105 lang=golang
 * @lcpr version=30204
 *
 * [105] 从前序与中序遍历序列构造二叉树
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

// 思路清晰：先通过先序表快速找到根节点，再通过中序表获取左子树一共有多少个结点的信息
// 遍历全部左子树范围，左子树的第一位为后续左右子树的根，放入二叉树中
// 当左子树为长度为 0 时，查找先序表后一位，该为为右子树开端
// 找到根节点，将根节点的左指针指向左子树构成的二叉树，右指针指向右子树构成的子树
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	// if len(preorder) == 1 {
	// 	return root
	// }
	rootVal := preorder[0]
	lengthOfLeftTree := 0
	for i, val := range inorder {
		if val == rootVal {
			lengthOfLeftTree = i
		}
	}
	leftPreorder := preorder[1 : lengthOfLeftTree+1]
	leftInorder := inorder[:lengthOfLeftTree]
	root.Left = buildTree(leftPreorder, leftInorder)
	rightPreorder := preorder[lengthOfLeftTree+1:]
	rightInorder := inorder[lengthOfLeftTree+1:]
	root.Right = buildTree(rightPreorder, rightInorder)

	return root
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,15,7]\n[9,3,15,20,7]\n
// @lcpr case=end

// @lcpr case=start
// [-1]\n[-1]\n
// @lcpr case=end

*/
