/*
 * @lc app=leetcode.cn id=199 lang=golang
 * @lcpr version=30204
 *
 * [199] 二叉树的右视图
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

// 思路:BFS 小实践。把 BFS 每层的队列的队列尾取出来
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	bfsQueue := []*TreeNode{root}
	for len(bfsQueue) != 0 {
		curr := bfsQueue
		ans = append(ans, curr[len(curr)-1].Val)
		bfsQueue = nil
		for _, node := range curr {
			if node.Left != nil {
				bfsQueue = append(bfsQueue, node.Left)
			}
			if node.Right != nil {
				bfsQueue = append(bfsQueue, node.Right)
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,null,5,null,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,null,null,null,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
