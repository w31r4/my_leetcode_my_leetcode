/*
 * @lc app=leetcode.cn id=101 lang=golang
 * @lcpr version=30204
 *
 * [101] 对称二叉树
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

// TreeNode 是二叉树节点的定义。
// (假设 TreeNode 结构已在别处定义或作为上下文给出)
// type TreeNode struct {
// Val int
// Left *TreeNode
// Right *TreeNode
// }

// isSymmetric 是主函数，用于判断一棵二叉树是否轴对称。
// 它接收树的根节点作为参数。
func isSymmetric(root *TreeNode) bool {
	// 基本情况：如果树为空 (root == nil)，则认为它是轴对称的。
	if root == nil {
		return true
	}
	// 如果树不为空，则对称性取决于其左子树和右子树是否互为镜像。
	// 我们调用辅助函数 isMirrored 来进行这个判断。
	return isMirrored(root.Left, root.Right)
}

// isMirrored 是一个辅助递归函数，用于判断两棵（子）树 node1 和 node2 是否互为镜像。
// node1 通常代表原始树中某个节点的左子节点（或其后代）。
// node2 通常代表原始树中对应节点的右子节点（或其后代）。
func isMirrored(node1, node2 *TreeNode) bool {
	// 基本情况 1：处理一个或两个节点为空的情况。
	// 如果 node1 或 node2 中至少有一个是 nil...
	if node1 == nil || node2 == nil {
		// ...那么只有当它们两个都同时为 nil 时，才算是互为镜像的。
		// 例如，一个节点的左孩子是 nil，对应镜像位置的右孩子也是 nil，这是对称的。
		// 但如果一个是 nil，另一个不是 nil，则不对称。
		return node1 == nil && node2 == nil
	}

	// 执行到这里，说明 node1 和 node2 都不为 nil。
	// 基本情况 2：比较当前两个节点的值。
	// 如果它们的值不相等，那么它们不可能是镜像的一部分。
	if node1.Val != node2.Val {
		return false
	}

	// 递归步骤：
	// 如果当前两个节点的值相等，我们还需要确保它们的子树也满足镜像关系。
	// 具体来说：
	// 1. node1 的左子树 (node1.Left) 必须与 node2 的右子树 (node2.Right) 互为镜像。
	// 2. 并且，node1 的右子树 (node1.Right) 必须与 node2 的左子树 (node2.Left) 互为镜像。
	// 只有这两个条件同时满足，node1 和 node2 所代表的结构才是互为镜像的。
	return isMirrored(node1.Left, node2.Right) && isMirrored(node1.Right, node2.Left)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,3,4,4,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,2,null,3,null,3]\n
// @lcpr case=end

*/
