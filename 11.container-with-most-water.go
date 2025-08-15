/*
 * @lc app=leetcode.cn id=11 lang=golang
 * @lcpr version=30204
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode.cn/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (61.08%)
 * Likes:    5392
 * Dislikes: 0
 * Total Accepted:    1.6M
 * Total Submissions: 2.6M
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
 *
 * 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 *
 * 返回容器可以储存的最大水量。
 *
 * 说明：你不能倾斜容器。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：[1,8,6,2,5,4,8,3,7]
 * 输出：49
 * 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 * 示例 2：
 *
 * 输入：height = [1,1]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == height.length
 * 2 <= n <= 10^5
 * 0 <= height[i] <= 10^4
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 思路：双指针，一个指向头部一个指向尾部，计算容量:(-head_index+tail_index)*min(head_val,tail_val)
// 比较 head_val 和 tail_val 大小，小的往里推进 (++or--)
// 直到 head_index==tail_index
func maxArea(height []int) int {
	head_index := 0
	tail_index := len(height) - 1
	maxContain := 0
	for head_index < tail_index {
		min_height := min(height[head_index], height[tail_index])
		nowContain := (tail_index - head_index) * min_height
		maxContain = max(nowContain, maxContain)
		if height[head_index] >= height[tail_index] {
			tail_index--
		} else {
			head_index++
		}
	}
	return maxContain
}

// @lc code=end

/*
// @lcpr case=start
// [1,8,6,2,5,4,8,3,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,1]\n
// @lcpr case=end

*/
