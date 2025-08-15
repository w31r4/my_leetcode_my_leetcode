/*
 * @lc app=leetcode.cn id=103 lang=golang
 * @lcpr version=30204
 *
 * [103] 二叉树的锯齿形层序遍历
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
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历，把当前层的值存入数组中，每遍历一次转换一次存储方向
func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	bfsQueue := []*TreeNode{root}
	isFromLeftToRight := true
	for len(bfsQueue) != 0 {
		curr := bfsQueue
		bfsQueue = nil
		currVal := []int{}
		for i, node := range curr {
			if isFromLeftToRight {
				currVal = append(currVal, node.Val)
			} else {
				currVal = append(currVal, curr[len(curr)-1-i].Val)
			}

			if node.Left != nil {
				bfsQueue = append(bfsQueue, node.Left)
			}
			if node.Right != nil {
				bfsQueue = append(bfsQueue, node.Right)
			}
		}
		isFromLeftToRight = !isFromLeftToRight
		ans = append(ans, currVal)
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
