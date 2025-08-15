/*
 * @lc app=leetcode.cn id=637 lang=golang
 * @lcpr version=30204
 *
 * [637] 二叉树的层平均值
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

// BFS 层序遍历并将每一个都加起来
func averageOfLevels(root *TreeNode) (ans []float64) {
	if root == nil {
		return
	}
	bfsQueue := []*TreeNode{root}
	for len(bfsQueue) != 0 {
		curr := bfsQueue
		bfsQueue = nil
		var sum float64
		for _, node := range curr {
			sum += float64(node.Val)
			if node.Left != nil {
				bfsQueue = append(bfsQueue, node.Left)
			}
			if node.Right != nil {
				bfsQueue = append(bfsQueue, node.Right)
			}
		}

		ans = append(ans, sum/float64(len(curr)))

	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [3,9,20,15,7]\n
// @lcpr case=end

*/
