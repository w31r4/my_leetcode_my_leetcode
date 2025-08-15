/*
 * @lc app=leetcode.cn id=129 lang=golang
 * @lcpr version=30204
 *
 * [129] 求根节点到叶节点数字之和
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

// 难点在于拼接
// 需要导入 math 包

// TreeNode 定义（假设题目已提供或类似如下）
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func sumNumbers(root *TreeNode) int {
	return dfsSumNum(root, 0)
}

// 原有数字不断的进位可以理解为一步一乘十，再加上当前值
func dfsSumNum(node *TreeNode, nowSum int) int {
	if node == nil {
		return 0
	}
	nowSum = nowSum*10 + node.Val
	//将原有数字乘 10 后加上现有数字
	if node.Left == nil && node.Right == nil {
		return nowSum
	}
	return dfsSumNum(node.Left, nowSum) + dfsSumNum(node.Right, nowSum)
}

//下面方法诸多不足，决心重构
// func dfsSumNum(node *TreeNode, slice []int) int {
// 	if node == nil { // 基线条件：如果节点为空，此路径不贡献任何值
// 		return 0
// 	}

// 	// 将当前节点的值加入路径
// 	// 问题 1：这里的 slice append 会影响到后续对兄弟节点及其子树的递归调用，
// 	// 因为 slice 是引用类型，如果 append 没有触发底层数组扩容，
// 	// 那么所有共享这个底层数组的 slice 都会看到这个修改。
// 	// 即便扩容，slice 变量也只在当前函数栈帧中更新，传递给递归调用的仍可能是旧 slice 的副本。
// 	// 正确的做法是确保每次递归有独立的路径状态，或在返回后进行回溯。
// 	slice = append(slice, node.Val)

// 	// 判断是否是叶子节点
// 	if node.Left == nil && node.Right == nil {
// 		length := len(slice) - 1 // 用于计算 10 的幂次，表示数字的位数
// 		sum := 0
// 		// 问题 2：每次到达叶子节点都重新计算整个数字，效率不高。
// 		// 例如路径 [1, 2, 3]，length = 2
// 		// i=0, num=1: tmp = 10^(2-0) = 100, sum = 100 * 1 = 100
// 		// i=1, num=2: tmp = 10^(2-1) = 10,  sum = 100 + 10 * 2 = 120
// 		// i=2, num=3: tmp = 10^(2-2) = 1,   sum = 120 + 1 * 3 = 123
// 		for i, num := range slice {
// 			tmp := math.Pow10(length - i) // 计算 10 的 (总长度 -1-当前索引) 次方
// 			sum += int(tmp) * num         // 累加 (位数权重 * 数字)
// 		}
// 		// 注意：因为 slice 在递归中被不当共享，这里的 sum 可能是基于一个错误的路径计算出来的。
// 		// 并且，因为没有回溯，当处理完左子树的一个叶子节点返回后，
// 		// slice 中仍然包含着左子树叶子节点的路径，这会污染右子树的计算。
// 		return sum
// 	}

// 递归计算左右子树的和
// 问题 1 的影响：传递给 dfsSumNum(node.Left, slice) 和 dfsSumNum(node.Right, slice) 的 slice
// 可能不是它们期望的、只包含从根到其父节点的路径。
// 例如，处理完左子树后，slice 包含了左子树的路径，这个被污染的 slice 被传给右子树。
// 	return dfsSumNum(node.Left, slice) + dfsSumNum(node.Right, slice)
// }

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [4,9,0,5,1]\n
// @lcpr case=end

*/
