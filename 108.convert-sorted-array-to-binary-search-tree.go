/*
 * @lc app=leetcode.cn id=108 lang=golang
 * @lcpr version=30204
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	// rootNum := length / 2
	var buildTree func(left, right int) *TreeNode
	buildTree = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)/2
		//边界条件
		root := &TreeNode{Val: nums[mid]}
		root.Left = buildTree(left, mid-1)
		root.Right = buildTree(mid+1, right)
		// if left {
		// 	root.Left = buildTree(index-1, left, false)
		// }
		// if right {
		// 	root.Right = buildTree(index+1, false, right)
		// }
		//这样做不算平衡树，平衡树的每一个节点都要求平衡
		return root
	}
	return buildTree(0, length-1)
}

// @lc code=end

/*
// @lcpr case=start
// [-10,-3,0,5,9]\n
// @lcpr case=end

// @lcpr case=start
// [1,3]\n
// @lcpr case=end

*/
