/*
 * @lc app=leetcode.cn id=56 lang=golang
 * @lcpr version=30204
 *
 * [56] 合并区间
 */

// @lcpr-template-start
package leetcode

import "slices"

// @lcpr-template-end
// @lc code=start
// 合并区间的规则是什么？
// 两个区间之间：两区间有交集，将两区间头最小值和区间尾最大值合并为新区间

// 但是如果不排序那么几乎做不了，所以先排序
// 按区间头大小进行排序
// 如何判断有没有交集呢？
// 现在我们已经按区间头大小进行排序
// 当当前 ans 二维数组中的最后一位数组区间尾大于等于当前区间头，那么进行合并，让两区间区间尾的最大值变为当前 ans 区间尾
// 如果没有交集，则创建新的 ans 中的数组，因为区间头已经按大小排序，所以后续不会有比之前区间尾更小的区间头，因为当前区间头一定大于之前的区间尾
func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })
	ans := [][]int{}
	for i := 0; i < len(intervals); i++ {
		length := len(ans)
		if length > 0 && ans[length-1][1] >= intervals[i][0] {
			ans[length-1][1] = max(intervals[i][1], ans[length-1][1])
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3],[2,6],[8,10],[15,18]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,4],[4,5]]\n
// @lcpr case=end

*/
