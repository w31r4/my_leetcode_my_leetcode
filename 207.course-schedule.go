/*
 * @lc app=leetcode.cn id=207 lang=golang
 * @lcpr version=30204
 *
 * [207] 课程表
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 当出现环或出现超越当前课程数量的课程时无法完成全部课程
// 需要验证这是不是一个有向无环图
func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	if numCourses == 1 {
		return false
	}
	// 一，建立图和入度表
	//建立图和入度表，只有遍历图的所有节点以后才能知道有没有环
	graph := make([][]int, numCourses)
	//建立图
	inDegreeTable := make([]int, numCourses)
	//入度表
	for _, pair := range prerequisites {
		v := pair[0]
		//v 是条件课程

		u := pair[1]
		//u 是前置课程
		if u >= numCourses || v >= numCourses {
			return false
		}
		//存在链接的节点就加入图中
		graph[u] = append(graph[u], v)
		inDegreeTable[v]++
		//需要创建 inode u
	}

	//二：按照拓扑排序遍历图，直到节点入度最小化
	notIndependentCourseNum := numCourses
	//BFS 遍历
	//需要找到入度表为 0 所有节点进行遍历
	zeroSet := []int{}
	for courseNum, inode := range inDegreeTable {
		if inode == 0 {
			zeroSet = append(zeroSet, courseNum)
			notIndependentCourseNum--
			//非独立课程减一
		}
	}
	for len(zeroSet) != 0 {
		curr := zeroSet
		zeroSet = []int{}
		for _, zeroNode := range curr {
			//将所有入度为 0 的节点移出入度表
			for _, node := range graph[zeroNode] {
				//将入度为 0 的节点指向的节点入度减一
				inDegreeTable[node]--
				if inDegreeTable[node] == 0 {
					zeroSet = append(zeroSet, node)
					notIndependentCourseNum--
					//入度为 0，那么
					//非独立课程减一
				}
			}
		}
	}
	return notIndependentCourseNum == 0
}

// @lc code=end

/*
// @lcpr case=start
// 2\n[[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// 2\n[[1,0],[0,1]]\n
// @lcpr case=end

*/
