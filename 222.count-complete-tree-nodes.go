/*
 * @lc app=leetcode.cn id=222 lang=golang
 * @lcpr version=30204
 *
 * [222] 完全二叉树的节点个数
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

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	hOfRightTree, hOfLeftTree := 0, 0
	rNode := root
	for rNode != nil {
		hOfRightTree++
		rNode = rNode.Right
	}
	lNode := root
	for lNode != nil {
		hOfLeftTree++
		lNode = lNode.Left
	}
	ans := 1<<hOfRightTree - 1
	if hOfLeftTree != hOfRightTree {
		// depth := 0
		//需要引入深度来判断是否纳入计数
		isTouchRightEnd := false
		var findDfs func(node *TreeNode, depth int)
		findDfs = func(node *TreeNode, depth int) {
			if isTouchRightEnd {
				return
			}
			depth++

			if node.Left == nil && node.Right == nil {
				if depth == hOfLeftTree {
					ans++
				} else {
					isTouchRightEnd = true
				}
				return
			}
			if node.Right == nil {
				isTouchRightEnd = true
				ans++
				//补偿被跳过的那个单个左子节点
				//有个致命问题是无法确认边缘右节点是叶子结点的情况
				//我们的算法目标是当有一个节点的右子树为空时才返回，且优先级在左右子树都为空之下
				//所以我引入了 depth 来计算深度，可以确保当遇到非左深度的叶子节点时可以及时识别并直接返回
				//目前该函数以及正常运行了
				return
			}
			findDfs(node.Left, depth)
			findDfs(node.Right, depth)
			return
		}
		findDfs(root, 0)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
