/*
 * @lc app=leetcode.cn id=275 lang=golang
 * @lcpr version=30204
 *
 * [275] H 指数 II
 *
 * https://leetcode.cn/problems/h-index-ii/description/
 *
 * algorithms
 * Medium (44.77%)
 * Likes:    369
 * Dislikes: 0
 * Total Accepted:    121.5K
 * Total Submissions: 271.5K
 * Testcase Example:  '[0,1,3,5,6]'
 *
 * 给你一个整数数组 citations，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数，citations 已经按照 升序排列
 * 。计算并返回该研究者的 h 指数。
 *
 * h 指数的定义：h 代表“高引用次数”（high citations），一名科研人员的 h 指数是指他（她）的（n 篇论文中）至少 有 h
 * 篇论文分别被引用了至少 h 次。
 *
 * 请你设计并实现对数时间复杂度的算法解决此问题。
 *
 *
 *
 * 示例 1：
 *
 * 输入：citations = [0,1,3,5,6]
 * 输出：3
 * 解释：给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 0, 1, 3, 5, 6 次。
 * 由于研究者有 3 篇论文每篇 至少 被引用了 3 次，其余两篇论文每篇被引用 不多于 3 次，所以她的 h 指数是 3。
 *
 * 示例 2：
 *
 * 输入：citations = [1,2,100]
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == citations.length
 * 1 <= n <= 10^5
 * 0 <= citations[i] <= 1000
 * citations 按 升序排列
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func hIndex1(citations []int) int {
	// h := 0
	// // h 表示比 h 大的数的数量

	// for i := len(citations) - 1; i >= 0; i-- {
	// 	if citations[i] > h {
	// 		h++
	// 	}
	// }
	// return h
	//二分查找法：闭区间，创建 left 和 right，先求出 mid，再看总索引-mid index 对应的值是否大于等于 mid，如果是继续向右找更大的数
	n := len(citations)
	left := 0
	//可能的最小指数
	right := n
	//可能的最大指数
	for left < right {
		mid := left + (right-left)/2
		//mid 是所 h 指数所包含的论文的篇数，当边界论文引用小于所对应的论文篇数时，我们就缩小包含范围到中间，并再次确认边界论文引用和包含论文篇数的大小关系，然后决定向左向右缩小包含范围
		//取中位数
		if citations[n-mid-1] > mid {
			//看看从最后一个往后数第 mid 个的引用量是否大于 mid
			//如果大于的话就将 mid 向右看，将探测是否大于 mid 的探针向左看，直到 left 和 right 重叠
			left = mid + 1
			//边界仅 +1
			// 			单调性：

			// 如果 h 满足条件，那么所有更小的 h 也一定满足。
			// 我们需要找最大的满足条件的 h。
			// 高效排除：
			// 如果 citations[n-h-1] >= h，说明 h 可以更大。
			//可以扩容查找
			// 否则，h 需要更小。
			//缩容查找，要看区间里面没有什么
			//不断地利用原有信息缩小搜索区间
		} else {
			right = mid
		}
	}
	return right
	//二分查找模板
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,3,5,6]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,100]\n
// @lcpr case=end

*/
