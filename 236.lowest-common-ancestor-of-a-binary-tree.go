/*
 * @lc app=leetcode.cn id=236 lang=golang
 * @lcpr version=30204
 *
 * [236] 二叉树的最近公共祖先
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

// 愚蠢的思路是：DFS 建立数组记录
// 返回路径之后找到公共祖先
// 还有更好的递归解法
// var isGetTarget bool

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	var ans *TreeNode
// 	isGetTarget = false
// 	pPath := []*TreeNode{}
// 	pPath = trackDfs(root, p.Val, pPath)
// 	isGetTarget = false
// 	qPath := []*TreeNode{}
// 	qPath = trackDfs(root, q.Val, qPath)
// 	minLength := min(len(qPath), len(pPath)) - 1
// 	for i := minLength; i >= 0; i-- {
// 		if qPath[i] == pPath[i] {
// 			ans = qPath[i]
// 			break
// 		}
// 	}
// 	return ans
// }

//	func trackDfs(node *TreeNode, target int, path []*TreeNode) []*TreeNode {
//		if isGetTarget {
//			return nil
//		}
//		if node == nil {
//			return nil
//		}
//		var currPath []*TreeNode
//		if node.Val == target {
//			currPath = append(path, node)
//			isGetTarget = true
//			return currPath
//		}
//		currPath = append(path, node)
//		leftPath := trackDfs(node.Left, target, currPath)
//		rightPath := trackDfs(node.Right, target, currPath)
//		if leftPath != nil {
//			return leftPath
//		}
//		if rightPath != nil {
//			return rightPath
//		}
//		return nil
//	}
//
// 接下来介绍递归解法
// 思考方法：当一个节点为最近祖先节点时有三种情况：
// 1.两个子节点存在左右子树中
// 2.当前节点为祖先节点，另一个节点在左子树中
// 3.当前节点为祖先节点，另一个节点在右子树中
// 从路过的每一个节点查找左右子树中是否存在目标节点，或者当前节点是否为目标节点，如果是就返回当前节点，同时用 nil 来判断左右子树中是否存在目标节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
		//找到目标节点，更新公共祖先为当前目标节点，如果公共祖先始终没有得到更新，那么当前目标节点就是公共祖先
	}
	//边界条件，当找到或者为空时返回当前节点
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//当前节点不是目标节点且不为空，那么以当前节点开始左右子树查找，如果，看看左右子树中是否有目标节点
	//当当前节点的左右子树中都有返回值，那么当前节点是最近的公共祖先
	//如果当前节点仅一边有返回值或没有返回值，那么这个当前节点绝对不是公共祖先
	//如果公共祖先是目标节点中的一员，那么会因为父节点查询其中一个子树为空而直接将公共祖先返回
	if left == nil || right == nil {
		if left == nil {
			return right
		} else {
			return left
		}
	}
	return root
	//当且仅当左右子树都为空时更新最近祖先节点并返回
	//后续的父节点会因为只有一个子树有返回值而无法跟新公共祖先
}

// @lc code=end

/*
// @lcpr case=start
// [3,5,1,6,2,0,8,null,null,7,4]\n5\n1\n
// @lcpr case=end

// @lcpr case=start
// [3,5,1,6,2,0,8,null,null,7,4]\n5\n4\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n2\n
// @lcpr case=end

*/
