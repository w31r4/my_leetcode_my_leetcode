/*
 * @lc app=leetcode.cn id=112 lang=golang
 * @lcpr version=30204
 *
 * [112] 路径总和
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
//初版简洁写法，可读性略差
// func hasPathSum(root *TreeNode, targetSum int) bool {
// 	var dfsSum func(*TreeNode, int) bool
// 	//声明函数签名
// 	dfsSum = func(node *TreeNode, sum int) bool {
// 		//实现函数签名
// 		if node == nil {
// 			return false
// 		}

// 		sum += node.Val
// 		if node.Left == nil && node.Right == nil {
// 			if sum == targetSum {
// 				return true
// 			} else {
// 				return false
// 			}
// 		}
// 		return dfsSum(node.Left, sum) || dfsSum(node.Right, sum)
// 	}
// 	return dfsSum(root, 0)
// }

// TreeNode 定义假设已存在
// type TreeNode struct {
//     Val   int
//     Left  *TreeNode
//     Right *TreeNode
// }

// hasPathSum 函数用于检查是否存在一条从根节点到叶子节点的路径，
// 使得路径上所有节点值的总和等于给定的 targetSum。
func hasPathSum(root *TreeNode, targetSum int) bool {
	// 声明一个名为 dfsSum 的辅助递归函数（闭包）的类型签名。
	// 这个函数将用于执行深度优先搜索。
	var dfsSum func(*TreeNode, int) bool
	// 用户注释：声明函数签名

	// 定义 dfsSum 闭包的具体实现。
	// node: 当前正在访问的节点。
	// currentSumToParent: 从根节点到当前 `node` 父节点的路径上所有节点值的总和。
	//                   对于初始调用 (根节点)，这个值是 0。
	dfsSum = func(node *TreeNode, sumToParent int) bool {
		// 用户注释：实现函数签名

		// 基本情况 1：如果当前节点为空 (nil)。
		// 这意味着当前路径已经结束（例如，尝试访问叶子节点的子节点），
		// 或者树本身就是空的。这条路径不满足条件。
		if node == nil {
			return false
		}

		// 计算包含当前节点 `node` 在内的路径总和。
		// `sumToParent` 是到父节点的和，现在加上当前节点的值。
		currentPathSum := sumToParent + node.Val // 这与你原来的 sum += node.Val 效果相同，
		// 只是创建了一个新变量，更清晰地表达了“到当前节点的总和”
		// 如果你使用 sum += node.Val，那么后续的 sum 就是指到当前节点的总和。

		// 基本情况 2：检查当前节点是否是叶子节点。
		// 叶子节点是指既没有左孩子也没有右孩子的节点。
		if node.Left == nil && node.Right == nil {
			// 如果是叶子节点，则判断从根到此叶子节点的路径和 (currentPathSum) 是否等于目标和 targetSum。
			if currentPathSum == targetSum {
				return true // 找到了符合条件的路径！
			} else {
				return false // 这条路径的和不等于目标和。
			}
		}

		// 递归步骤：如果当前节点不是叶子节点，则继续在其左子树和右子树中寻找。
		// 将包含当前节点值的路径和 (currentPathSum) 传递给子节点的递归调用。
		// 只要左子树或右子树中任何一条路径能够最终满足条件，就认为找到了。
		// (注意：你原来的代码 `sum += node.Val` 后直接传递 `sum` 也是正确的，
		//  因为 `sum` 在每次递归调用时是作为参数传递的副本，
		//  所以 `sum += node.Val` 的修改只在当前函数帧内有效，
		//  传递给下一层的 `sum` 就是已经包含了当前节点值的路径和。)
		foundInLeft := dfsSum(node.Left, currentPathSum)
		if foundInLeft { // 如果在左子树中找到了，就没必要再搜索右子树了
			return true
		}

		foundInRight := dfsSum(node.Right, currentPathSum)
		return foundInRight

		// 你原来的写法 `return dfsSum(node.Left, sum) || dfsSum(node.Right, sum)` 也是完全正确的，
		// 并且更简洁。我这里分开写是为了注释更清晰。
		// 如果用你原来的写法，sum 在 `sum += node.Val` 之后，就代表了 currentPathSumIncludingNode。
		// 所以，以下是你代码的直接注释版本：
		/*
		   sum += node.Val // `sum` 现在是从根到当前节点 `node` (包括 node) 的路径和

		   if node.Left == nil && node.Right == nil { // 判断是否为叶子节点
		       if sum == targetSum { // 如果路径和等于目标值
		           return true
		       } else {
		           return false // 否则，这条路不通
		       }
		   }
		   // 如果不是叶子节点，则递归地在左子树和右子树中继续寻找。
		   // 将当前的路径和 `sum` (已包含当前节点值) 传递下去。
		   // 只要任意一边能找到符合条件的路径，就返回 true。
		   return dfsSum(node.Left, sum) || dfsSum(node.Right, sum)
		*/
	}

	// 初始调用 dfsSum：
	// 从根节点 root 开始搜索，初始的“到父节点的路径和”为 0。
	return dfsSum(root, 0)
}

//还有一种神奇写法：
//之和可以当成逐步递减直到为 0
// func hasPathSum(root *TreeNode, targetSum int) bool {
//     if root == nil {
//         return false
//     }
//     targetSum -= root.Val
//     if root.Left == root.Right { // root 是叶子
//         return targetSum == 0
//     }
//     return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
// }

// @lc code=end

/*
// @lcpr case=start
// [5,4,8,11,null,13,4,7,2,null,null,null,1]\n22\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// []\n0\n
// @lcpr case=end

*/
