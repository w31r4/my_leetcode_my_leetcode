/*
 * @lc app=leetcode.cn id=133 lang=golang
 * @lcpr version=30204
 *
 * [133] 克隆图
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
// type Node struct {
// 	Val       int
// 	Neighbors []*Node
// }

// 深拷贝，给每一个原节点都创建一个新的 neighbor
// 不能修改原图，之前拷贝链表是可以将其还原的，但是图的还原很复杂，而且这样的思路行不通，实现太过复杂
// 我们可以使用我们的哈希表保存新创建的克隆节点
// 要标记自己经过的结点
// 使用 map 标记，不要破坏原有信息
// 不要不舍得用 map
// type Node struct {
// 	Val       int
// 	Neighbors []*Node
// }

// 需要
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	cloneMap := map[int]*Node{}
	var getCloneNodeDfs func(originalNode *Node) *Node // 声明 dfs 辅助函数
	getCloneNodeDfs = func(originalNode *Node) *Node {
		if clonedNode, ok := cloneMap[originalNode.Val]; ok {
			return clonedNode
		}
		cloneNode := &Node{Val: originalNode.Val, Neighbors: []*Node{}}
		cloneMap[originalNode.Val] = cloneNode
		for _, neighbor := range originalNode.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, getCloneNodeDfs(neighbor))
		}
		return cloneNode
	}
	return getCloneNodeDfs(node)

}

// func cloneGraph(node *Node) *Node {
// 	if node == nil {
// 		return nil
// 	}

// 	// cloneMap 用于存储原始节点到其对应克隆节点的映射
// 	// 键是原始节点的指针，值是克隆节点的指针
// 	cloneMap := make(map[*Node]*Node)

// 	var dfs func(originalNode *Node) *Node // 声明 dfs 辅助函数

// 	dfs = func(originalNode *Node) *Node {
// 		// 如果 originalNode 已经在 map 中，说明它的克隆已被创建，直接返回克隆节点
// 		if clonedNode, ok := cloneMap[originalNode]; ok {
// 			return clonedNode
// 		}

// 		// 1. 创建当前 originalNode 的克隆节点
// 		//    注意：Neighbors 列表先初始化为空，后续会填充
// 		clonedNode := &Node{Val: originalNode.Val, Neighbors: []*Node{}}

// 		// 2. *立即*将原始节点和克隆节点的映射关系存入 map
// 		//    这非常重要，用于处理图中的环路，以及确保每个节点只被克隆一次
// 		cloneMap[originalNode] = clonedNode

// 		// 3. 递归克隆当前 originalNode 的所有邻居
// 		for _, originalNeighbor := range originalNode.Neighbors {
// 			// 递归调用 dfs 获取邻居的克隆节点
// 			clonedNeighbor := dfs(originalNeighbor)
// 			// 将克隆的邻居添加到当前克隆节点的 Neighbors 列表中
// 			clonedNode.Neighbors = append(clonedNode.Neighbors, clonedNeighbor)
// 		}
// 		// 4. 返回当前创建的克隆节点
// 		return clonedNode
// 	}

// 	// 从起始节点开始 DFS 克隆过程
// 	return dfs(node)
// }

// @lc code=end

/*
// @lcpr case=start
// [[2,4],[1,3],[2,4],[1,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
