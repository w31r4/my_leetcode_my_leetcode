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

func construct(grid [][]int) *Node {
	length := len(grid)
	var buildTree func(int, int, int) *Node
	buildTree = func(size, row, col int) *Node {
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
			return &Node{Val: valBool, IsLeaf: true}
		}
		//当区域内非纯色时开始递归四等分小方格
		halfSize := size / 2
		node := &Node{Val: true, IsLeaf: false}
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
