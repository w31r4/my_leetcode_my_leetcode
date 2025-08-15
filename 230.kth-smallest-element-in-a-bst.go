/*
 * @lc app=leetcode.cn id=230 lang=golang
 * @lcpr version=30204
 *
 * [230] 二叉搜索树中第 K 小的元素
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
//如何设计代码？
//全局中有个 k 记录目标第几层，同时维护一个 i 记录当前第几层
//当 i==k 时直接全部剪枝返回
func kthSmallest(root *TreeNode, k int) (ans int) {
	isGetTarget := false
	nowLevel := 0
	var midDfs func(node *TreeNode)
	midDfs = func(node *TreeNode) {
		if isGetTarget {
			return
		}
		if node == nil {
			return
		}
		midDfs(node.Left)
		nowLevel++
		if nowLevel == k {
			isGetTarget = true
			ans = node.Val
			return
		}
		midDfs(node.Right)
		return

	}
	midDfs(root)
	return
}

// @lc code=end

/*
// @lcpr case=start
// [3,1,4,null,2]\n1\n
// @lcpr case=end

// @lcpr case=start
// [5,3,6,2,4,null,null,1]\n3\n
// @lcpr case=end

*/
