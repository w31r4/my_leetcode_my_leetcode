/*
 * @lc app=leetcode.cn id=226 lang=golang
 * @lcpr version=30204
 *
 * [226] 翻转二叉树
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
//先找到所有的左右节点，然后反转他们
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

// 这个问题也很好地体现了“问题的递归结构是内嵌的，难点在于总结出来”。“反转一棵树”这个操作，其自然的递归定义就是：

// 如果树为空，反转结果还是空树。
// 如果树不为空，它的反转结果是这样一棵树：其根节点不变，其新的左子树是原右子树的反转，其新的右子树是原左子树的反转。你的代码正是这个定义的直接体现。

// @lc code=end

/*
// @lcpr case=start
// [4,2,7,1,3,6,9]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
