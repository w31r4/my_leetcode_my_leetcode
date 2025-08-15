/*
 * @lc app=leetcode.cn id=57 lang=golang
 * @lcpr version=30204
 *
 * [57] 插入区间
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 先找到第一个重叠区间，如果没有就插入到两个区间之间，如果有重叠区间就找到所有和重叠区间相重叠的区间并合并
func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	isIntervals := 0
	//状态码
	// 	0：初始状态，表示 newInterval 还没有被处理（既没有合并，也没有直接放入结果 ans）。
	// 1：表示 newInterval 当前正与 intervals 中的一个或多个区间发生重叠，并且正在进行合并。合并后的结果暂时存储在更新后的 newInterval 变量中。
	// 2：表示合并过程已经结束，合并后的 newInterval 已经被添加到了 ans 中，现在开始处理 intervals 中后续不重叠的区间。
	// 3：表示 newInterval 在没有与任何区间重叠的情况下，被直接插入到了 ans 中（通常是它比当前遍历到的 nowVal 要小，且不重叠）。
	for i, nowVal := range intervals {
		if (newInterval[0] >= nowVal[0] && newInterval[0] <= nowVal[1]) || (newInterval[1] >= nowVal[0] && newInterval[1] <= nowVal[1]) || (newInterval[0] <= nowVal[0] && newInterval[1] >= nowVal[1]) {
			//如果出现重叠情况，融合所有重叠区间
			newInterval[0] = min(nowVal[0], newInterval[0])
			newInterval[1] = max(nowVal[1], newInterval[1])
			// newInterval = intervals[i]
			isIntervals = 1
			if i == len(intervals)-1 {
				ans = append(ans, newInterval)
			}
			continue
		}
		if isIntervals == 1 {
			//不再重叠
			ans = append(ans, newInterval)
			ans = append(ans, nowVal)
			isIntervals = 2
			//设置状态为重叠过了，且 new 填充过了
			continue
		}
		if isIntervals == 0 && nowVal[0] > newInterval[1] {
			//当没重叠过且区间已经被略过时

			ans = append(ans, newInterval)
			isIntervals = 3
			//已经被填充过了
		}
		ans = append(ans, nowVal)
	}
	if isIntervals == 0 {
		//new 根本没有当没有被填充过时
		ans = append(ans, newInterval)
	}
	return
	//法一：状态机法
	//接下来是法二：分阶段处理法
	//区间可以分为三个阶段，阶段一：区间尾在新区间头部之前，我们将这部分的区间直接插入答案数组
	//阶段二：区间和新区间有交集，我们将这部分区间融合后插入答案数组
	//阶段三：区间头大于融合后的新区间尾，我们将这段区间直接插入答案数组
	// i := 0

	// // isExpand := false
	// for i < len(intervals) && intervals[i][1] < newInterval[0] {
	// 	ans = append(ans, intervals[i])
	// 	i++
	// } //对阶段一的处理结束
	// for i < len(intervals) && ((newInterval[0] >= intervals[i][0] && newInterval[0] <= intervals[i][1]) || (newInterval[1] >= intervals[i][0] && newInterval[1] <= intervals[i][1]) || (newInterval[0] <= intervals[i][0] && newInterval[1] >= intervals[i][1])) {
	// 	newInterval[0] = min(intervals[i][0], newInterval[0])
	// 	newInterval[1] = max(intervals[i][1], newInterval[1])
	// 	// isExpand = true
	// 	i++
	// }
	// ans = append(ans, newInterval)

	// //对第二阶段的处理结束
	// for i < len(intervals) && intervals[i][0] > newInterval[1] {
	// 	ans = append(ans, intervals[i])
	// 	i++
	// }

	// return
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3],[6,9]]\n[2,5]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[3,5],[6,7],[8,10],[12,16]]\n[4,8]\n
// @lcpr case=end

*/
