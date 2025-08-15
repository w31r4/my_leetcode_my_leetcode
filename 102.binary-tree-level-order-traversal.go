/*
 * @lc app=leetcode.cn id=102 lang=golang
 * @lcpr version=30204
 *
 * [102] 二叉树的层序遍历
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

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	bfsQueue := []*TreeNode{root}
	for len(bfsQueue) != 0 {
		curr := bfsQueue
		bfsQueue = nil
		currVal := []int{}
		for _, node := range curr {
			currVal = append(currVal, node.Val)
			if node.Left != nil {
				bfsQueue = append(bfsQueue, node.Left)
			}
			if node.Right != nil {
				bfsQueue = append(bfsQueue, node.Right)
			}
		}
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
