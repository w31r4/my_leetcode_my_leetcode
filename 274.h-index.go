/*
 * @lc app=leetcode.cn id=274 lang=golang
 * @lcpr version=30204
 *
 * [274] H 指数
 *
 * https://leetcode.cn/problems/h-index/description/
 *
 * algorithms
 * Medium (46.27%)
 * Likes:    575
 * Dislikes: 0
 * Total Accepted:    275.8K
 * Total Submissions: 596.2K
 * Testcase Example:  '[3,0,6,1,5]'
 *
 * 给你一个整数数组 citations，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。
 *
 * 根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，并且 至少 有 h
 * 篇论文被引用次数大于等于 h。如果 h 有多种可能的值，h 指数 是其中最大的那个。
 *
 *
 *
 * 示例 1：
 *
 * 输入：citations = [3,0,6,1,5]
 * 输出：3
 * 解释：给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
 * 由于研究者有 3 篇论文每篇 至少 被引用了 3 次，其余两篇论文每篇被引用 不多于 3 次，所以她的 h 指数是 3。
 *
 * 示例 2：
 *
 * 输入：citations = [1,3,1]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == citations.length
 * 1 <= n <= 5000
 * 0 <= citations[i] <= 1000
 *
 *
 */

// @lcpr-template-start
package leetcode

import "sort"

// @lcpr-template-end
// @lc code=start
func hIndex(citations []int) int {
	// 1.h<len(citation)
	// 2.有 h 篇论文的引用>=h
	//思路：遍历一遍，维护 h 指数，初始值为 0，遇到比 h 指数大的就加一，同时维护一个哈希表，在遍历结束后查表，看看有没有对应值，如果没有则减一，直到有为止
	//不对，还得哈希表，当前方法逻辑不对，当 h 为 2 时，可能只有一个 2, 可以试试先 hash 表存储后积分制，当有数大于当前 h 时，h++
	//或者直接先排序，倒序遍历一遍，到中位就停了
	//狗屎，一直以为 h 是在 citations 数组里面的其中一个值，没想到就是最大的那个
	//好好读题目比什么都重要
	h := 0
	// h 表示比 h 大的数的数量
	sort.Ints(citations)
	//快排
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] > h {
			h++
		}
	}
	return h

}

// @lc code=end

/*
// @lcpr case=start
// [1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,1]\n
// @lcpr case=end

*/
