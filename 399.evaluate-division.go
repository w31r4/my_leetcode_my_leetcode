/*
 * @lc app=leetcode.cn id=399 lang=golang
 * @lcpr version=30204
 *
 * [399] 除法求值
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 两难，链表或数组

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// --- 第一步：构建图 ---
	// 我们应该如何构建图？
	// 我们需要一个足够灵活的数据结构供我们查询当前节点的邻接节点，并保存该邻接节点，同时我们需要保存每个邻接节点双向的权重
	// 通过哈希表吗，我们需要查询当前查询量是否存在于当前图中
	//而且存在不连通的节点，但是没关系，我们保存的是每一个节点的邻接节点而不是
	graph := map[string]map[string]float64{}
	for index, equation := range equations {
		e0 := equation[0]
		e1 := equation[1]

		if graph[e0] == nil {
			graph[e0] = map[string]float64{}
		}
		if graph[e1] == nil {
			graph[e1] = map[string]float64{}
			// 不能解引用空指针
			// 我们仅仅创建了外部表
			// 有多少内部表 go 语言也不知道
			// 所以默认为空，空就是空指针
		}

		graph[e0][e1] = values[index]
		graph[e1][e0] = 1 / values[index]
	}
	// --- 第二步：处理每个查询 ---
	results := make([]float64, len(queries))
	for i, query := range queries {
		startNode, endNode := query[0], query[1]
		//我们通过哈希表建立了图，通过查询
		// 预检查
		if _, existsStart := graph[startNode]; !existsStart {

			results[i] = -1.0
			continue
		}
		if _, existsEnd := graph[endNode]; !existsEnd {
			results[i] = -1.0
			continue
		}
		//如果其中一个查询值不存在
		if startNode == endNode {
			results[i] = 1.0
			continue
		}

		// BFS 初始化
		// queue := []NodeInfo{{Name: startNode, Product: 1.0}}
		queue := []string{startNode}
		queueVal := []float64{1}
		//我们要找到 endNode，并返回一个乘积
		//我们需要同时记录两个值，一个是当前 queue 中的节点，另一个是当前乘积总和
		visited := make(map[string]bool)
		found := false

		for len(queue) > 0 {
			//开始对当前查询进行 bfs 遍历
			//层序遍历整个图
			current := queue
			currentProduct := queueVal
			queueVal = nil
			queue = nil // 出队
			visited[startNode] = true
			for index, nowNode := range current {
				if nowNode == endNode {
					found = true
					results[i] = currentProduct[index]
					queue = nil
					break
				}
				neighbors := graph[nowNode]
				for neighborName, edgeWeight := range neighbors {
					if visited[neighborName] == false {
						visited[neighborName] = true
						queue = append(queue, neighborName)
						queueVal = append(queueVal, edgeWeight*currentProduct[index])
					}
				}

			}
		}

		if !found { // 如果 BFS 结束仍未找到路径
			results[i] = -1.0
		}
	}

	return results
}

// @lc code=end

/*
// @lcpr case=start
// [["a","b"],["b","c"]]\n[2.0,3.0]\n[["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]\n
// @lcpr case=end

// @lcpr case=start
// [["a","b"],["b","c"],["bc","cd"]]\n[1.5,2.5,5.0]\n[["a","c"],["c","b"],["bc","cd"],["cd","bc"]]\n
// @lcpr case=end

// @lcpr case=start
// [["a","b"]]\n[0.5]\n[["a","b"],["b","a"],["a","c"],["x","y"]]\n
// @lcpr case=end

*/
