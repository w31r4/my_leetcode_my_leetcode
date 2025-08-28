/*
 * @lc app=leetcode.cn id=427 lang=golang
 * @lcpr version=30204
 *
 * [427] 建立四叉树
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 一直在递归执行当前大方格是否纯色？如果纯色标记 isLeaf 并赋值 val 后返回，不纯色就把当前大方格切成四等分，然后继续问当前方格是否纯色
// 整体思路：使用“分治” (Divide and Conquer) 的递归思想来解决问题。
//
// 我们定义一个递归辅助函数 `buildTree`，它的核心任务是：
// “给我一个指定的正方形区域（由起始点 row, col 和边长 size 定义），
// 我负责为这个区域构建一个四叉树节点并返回。”
//
// 在 `buildTree` 函数内部，主要逻辑分为两步：
// 1. 检查当前区域是否为“纯色”：
//   - 遍历这个区域内的所有格子，判断它们的值是否完全相同。
//
// 2. 根据检查结果做出决策：
//
//   - 如果是纯色的（递归的终止条件 Base Case）：
//
//   - 创建一个叶子节点 (IsLeaf: true)，Val 设为该区域的颜色，然后直接返回。
//
//   - 如果不是纯色的（递归的递推步骤 Recursive Step）：
//
//   - 创建一个内部节点 (IsLeaf: false)。
//
//   - 将当前区域“一分为四”，计算出四个子区域的起始点和大小。
//
//   - 递归调用 `buildTree` 四次，分别处理这四个子区域。
//
//   - 将返回的四个子节点连接到当前内部节点的四个孩子指针上，然后返回这个内部节点。
//
// 主函数 `construct` 的作用就是启动整个过程，它对整个 grid (从 (0,0) 开始，边长为 n) 调用一次 `buildTree`。
func construct(grid [][]int) *Node2 {
	length := len(grid)
	var buildTree func(int, int, int) *Node2
	buildTree = func(size, row, col int) *Node2 {
		val := grid[row][col]

		isPure := true
		//遍历区域内所有节点是否纯色

		for i := row; i < row+size; i++ {
			for j := col; j < col+size; j++ {
				if grid[i][j] != val {
					isPure = false
					break
				}
			}
		}
		//当区域内全部纯色时返回叶子节点
		if isPure {
			valBool := true
			if val == 0 {
				valBool = false
			}
			return &Node2{Val: valBool, IsLeaf: true}
		}
		//当区域内非纯色时开始递归四等分小方格
		halfSize := size / 2
		node := &Node2{Val: true, IsLeaf: false}
		node.TopLeft = buildTree(halfSize, row, col)
		node.TopRight = buildTree(halfSize, row, col+halfSize)
		node.BottomLeft = buildTree(halfSize, row+halfSize, col)
		node.BottomRight = buildTree(halfSize, row+halfSize, col+halfSize)
		return node
	}
	return buildTree(length, 0, 0)
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1],[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,1,1,1,1],[1,1,1,1,1,1,1,1],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0]]\n
// @lcpr case=end

*/
