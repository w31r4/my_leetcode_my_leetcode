/*
 * @lc app=leetcode.cn id=210 lang=golang
 * @lcpr version=30204
 *
 * [210] 课程表 II
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 这道题默认是所有课程都有关联都在图上，不存在不连通图
func findOrder(numCourses int, prerequisites [][]int) []int {
	if len(prerequisites) == 0 {
		ans := make([]int, numCourses)
		for i := 0; i < numCourses; i++ {
			ans[i] = i
		}
		return ans
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
		//存在链接的节点就加入图中
		graph[u] = append(graph[u], v)
		inDegreeTable[v]++
		//需要创建 inode u
	}

	//二：按照拓扑排序遍历图，直到节点入度最小化
	notIndependentCourseNum := numCourses
	courseRoute := []int{}
	//BFS 遍历
	//需要找到入度表为 0 所有节点进行遍历
	zeroSet := []int{}
	for courseNum, inode := range inDegreeTable {
		if inode == 0 {
			zeroSet = append(zeroSet, courseNum)
			notIndependentCourseNum--
			courseRoute = append(courseRoute, courseNum)
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
					courseRoute = append(courseRoute, node)

					//入度为 0，那么
					//非独立课程减一
				}
			}
		}
	}
	if notIndependentCourseNum == 0 {
		return courseRoute
	}
	return nil
}

// @lc code=end

/*
// @lcpr case=start
// 2\n[[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[[1,0],[2,0],[3,1],[3,2]]\n
// @lcpr case=end

// @lcpr case=start
// 1\n[]\n
// @lcpr case=end

*/
