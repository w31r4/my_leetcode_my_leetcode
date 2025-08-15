/*
 * @lc app=leetcode.cn id=100 lang=golang
 * @lcpr version=30204
 *
 * [100] 相同的树
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

// 可以把查询两颗树是否相同看成是查询两颗树的根和左右子节点是否相同
// 先查左节点在查右节点最后查 root 节点
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil { // 如果两者都为 nil，则相同
		return true
	}
	if p == nil || q == nil { // 如果执行到这里，说明两者不都为 nil，但至少有一个是 nil，则不同
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	// 这里的逻辑是：“两棵树相同”当且仅当：
	// 它们的根节点值相同 (已经在上一步检查过了)。
	// 并且，它们的左子树相同。
	// 并且，它们的右子树相同。
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n[1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[1,null,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,1]\n[1,1,2]\n
// @lcpr case=end

*/
