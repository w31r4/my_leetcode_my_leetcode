/*
 * @lc app=leetcode.cn id=452 lang=golang
 * @lcpr version=30204
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lcpr-template-start
package leetcode

import (
	"slices"
)

// @lcpr-template-end
// @lc code=start
// 这里建议先排序
// 不排序算法复杂度初步估算为 O n^2
// 将多个气球融合为重叠区间，一个重叠区间射一箭
// 可以按左端点或者右端点排序，我们求取的是在重叠区间中一箭最多可以射到多远，而下一个气球是否在这个范围内，如果不在那么需要多射一箭
func findMinArrowShots(points [][]int) (ans int) {
	slices.SortFunc(points, func(q, p []int) int { return q[0] - p[0] })
	maxRange := points[0][1]
	ans = 1
	for i := 1; i < len(points); i++ {
		if maxRange >= points[i][0] {
			//判断当前气球是否在箭的射程内
			maxRange = min(maxRange, points[i][1])
			//更新箭的射程，确保这一箭可以射到所有气球
		} else {
			ans++
			maxRange = points[i][1]
			//超出射程，新射一箭，更新箭发射到的距离
		}
	}
	//下面是优化方案，可以使用 if else 语句来控制流程，因为一共就两种可能性
	// for i := 1; i < len(points); i++ {
	// 	if points[i][0] > currentArrowReachEnd { // 当前气球的起点在上一箭的覆盖范围之外
	// 		ans++
	// 		currentArrowReachEnd = points[i][1] // 新的箭，新的覆盖结束点
	// 	} else {
	// 		// 当前气球可以被同一支箭覆盖，更新箭的覆盖结束点（取更小者，保证能射中当前这个）
	// 		currentArrowReachEnd = min(currentArrowReachEnd, points[i][1])
	// 	}
	// }
	return
	// slices.SortFunc(points, func(a, b []int) int { return a[1] - b[1] }) // 按照右端点从小到大排序
	// pre := math.MinInt
	// for _, p := range points {
	// 	if p[0] > pre { // 上一个放的点在区间左边
	// 		ans++
	// 		pre = p[1] // 在区间的最右边放一个点
	// 	}
	// }
	// return
}

// @lc code=end

/*
// @lcpr case=start
// [[10,16],[2,8],[1,6],[7,12]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[3,4],[5,6],[7,8]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[2,3],[3,4],[4,5]]\n
// @lcpr case=end

*/
