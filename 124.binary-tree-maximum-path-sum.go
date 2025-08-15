/*
 * @lc app=leetcode.cn id=124 lang=golang
 * @lcpr version=30204
 *
 * [124] 二叉树中的最大路径和
 */

// @lcpr-template-start
package leetcode

import (
	"math"
)

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

// 我们对每个结点都采取这样的策略：创建双路贡献，当下一级的单路贡献大于 0 时我们采纳下一级，并将其添置到贡献和中，如果不大于零，我们通过将贡献置零的方法将其贡献抹除的同时将其贡献和代表的节点移出当前节点的贡献路径中
// 节点最后都变成了从下至上的一段段的贡献数值，当某路径的贡献数值小于 0 时，将其在上游返回的贡献数值设置为零可以等价为将其移除出了上游节点的最大路径中
// 目标是从下到上收集所有可以造成正贡献的节点，如果不能造成正贡献，那么将其返回为 0，也即抛弃了这一段节点
// 如果全为负数值，那么最大的那个值肯定是单个负数
// 返回值为单边最大贡献，因为上一级只能连接下一级的单边作为自己的连接路径，路径不存在分叉
// 这里用的是后序遍历
func maxPathSum(root *TreeNode) int {
	// maxNum 用于记录和更新全局的最大路径和。
	// 初始化为 math.MinInt32 是为了确保任何有效的路径和（即使是负数）都能比它大。
	maxNum := math.MinInt32

	// 定义一个递归辅助函数 maxGet (或者叫 gainFromSubtree, oneSideMax 等)
	// 这个函数返回的是：从当前 node 出发向其子树延伸的“单边最大路径和”
	// （即 node.Val + max(左单边贡献，右单边贡献)，且子贡献<0 则不取）
	var maxGet func(node *TreeNode) int
	maxGet = func(node *TreeNode) int {
		if node == nil { // 基线条件：空节点对路径和的贡献为 0
			return 0
		}

		// nowVal:= node.Val // 当前节点的值

		// 递归计算左子树能提供的最大单边贡献。
		// 如果左子树的贡献是负数，则我们不选择它，所以用 max(0, ...)
		leftGet := max(0, maxGet(node.Left))
		//注意：这里为什么要和 0 比大小？
		//当下一级的返回的贡献实际小于 0 时，我们可以直接通过将贡献清零来将下一级的节点移出我们当前所求的节点路径
		// 递归计算右子树能提供的最大单边贡献。
		// 同样，如果贡献是负数，则不选择，取 max(0, ...)

		rightGet := max(0, maxGet(node.Right))

		// 更新全局最大路径和：
		// 考虑以当前 node 为“拱顶”的路径，其和为 node.Val + 左贡献 + 右贡献。
		// 将这个值与已记录的 maxNum 比较，取较大者。
		maxNum = max(node.Val+leftGet+rightGet, maxNum)
		//当前节点是可以左右子路贡献齐聚的
		// 返回给父节点的“单边最大贡献”：
		// 当前节点 node 加上其左右子树中贡献较大的那一边。
		// 这个返回值用于其父节点构建更长的单边路径。
		//但是返回给父节点的子路贡献只能是单边的而非双边的，我们的路径不存在分叉
		//如果当前节点的左右子路贡献都不大于零，那么就仅仅返回当前节点作为下级贡献和，下级节点作为贡献不大于零的节点被移除路径和中
		//如果贡献大于零，那么拖家带口转到上一级去，成为上一级的单路贡献之一
		return node.Val + max(leftGet, rightGet)
	}

	// 从根节点开始递归计算
	maxGet(root)

	// 最终，maxNum 中存储的就是整个树中的最大路径和
	return maxNum
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [-10,9,20,null,null,15,7]\n
// @lcpr case=end

*/
