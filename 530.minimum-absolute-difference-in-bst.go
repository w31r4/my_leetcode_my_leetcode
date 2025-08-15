/*
 * @lc app=leetcode.cn id=530 lang=golang
 * @lcpr version=30204
 *
 * [530] 二叉搜索树的最小绝对差
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

// 先中序再遍历
func getMinimumDifference(root *TreeNode) int {
	minDiff := math.MaxInt32
	pre := (math.MinInt32) / 2
	// saveNodes := []int{}
	var midDfsWithDiff func(node *TreeNode)
	midDfsWithDiff = func(node *TreeNode) {
		if node == nil {
			return
		}
		midDfsWithDiff(node.Left)
		//二叉搜索树的特性是，在中序遍历中，当前节点值总是一级一级变大的
		minDiff = min(minDiff, node.Val-pre)
		//所以我们可以事先确定要相减的两数大小
		//并将其相减
		pre = node.Val
		midDfsWithDiff(node.Right)
		return
	}
	midDfsWithDiff(root)
	return minDiff
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,6,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,48,null,null,12,49]\n
// @lcpr case=end

*/
